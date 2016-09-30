package sql

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/lostz/sql/charset"
)

type Lexer struct {
	reader    *bufio.Reader
	buf       []byte
	pos       int
	lastPos   int
	nextState uint
	tokStart  int
	charset   *charset.CharsetInfo
	mode      mode
	inComment uint
	ParseTree Statement
	LastError string
}

type mode struct {
	ModeAnsiQuotes         bool
	ModeHighNotPrecedence  bool
	ModePipesAsConcat      bool
	ModeNoBackslashEscapes bool
	ModeIgnoreSpace        bool
}

func NewLexer(sql string) *Lexer {
	s := &Lexer{
		reader:    bufio.NewReader(strings.NewReader(sql)),
		buf:       []byte(sql),
		nextState: MyLexStart,
		charset:   charset.CSUtf8GeneralCli,
	}
	return s
}

func (s *Lexer) Lex(lval *MySQLSymType) (retstate int) {
	var resultState int
	var c byte
	var length int
	state := s.nextState
	cs := s.charset
	stateMap := cs.StateMap
	identMap := cs.IdentMap
	s.nextState = MyLexOperatorOrIdent
	for {
		switch state {
		case MyLexOperatorOrIdent, MyLexStart:
			for c = s.next(); stateMap[c] == MyLexSkip; c = s.next() {

			}
			s.tokStart = s.pos - 1
			state = stateMap[c]
		case MyLexEscape:
			if s.next() == 'N' {
				retstate = NULL_SYM
				return
			}
			fallthrough
		case MyLexChar, MyLexSkip:
			if c == '-' && s.peek(0) == '-' && (cs.IsSpace(s.peek(1)) || cs.IsCntrl(s.peek(1))) {
				state = MyLexComment
				break
			}
			s.pos = s.tokStart

			c = s.next()
			if c != ')' {
				s.nextState = MyLexStart
			}

			if c == ',' {
				s.tokStart = s.pos
			} else if c == '?' && identMap[s.peek(0)] == 0 {
				retstate = PARAM_MARKER
				return
			}
			retstate = int(c)
			return
		case MyLexIdentOrNchar:
			if s.peek(0) != '\'' {
				state = MyLexIdent
				break
			}

			retstate, c = s.scanChar(lval)
			return
		case MyLexIdentOrHex:
			if s.peek(0) == '\'' {
				state = MyLexHexNumber
				break
			}
			fallthrough
		case MyLexIdentOrBin:
			if s.peek(0) == '\'' {
				state = MyLexBinNumber
				break
			}
			fallthrough
		case MyLexIdent:
			retstate, lval.bytes = s.scanIdentifier()
			return
		case MyLexIdentSep: // Found ident before
			// And Now '.'
			c = s.next()
			s.nextState = MyLexIdentStart
			if identMap[s.peek(0)] == 0 {
				s.nextState = MyLexStart
			}
			retstate = int(c)
			return
		case MyLexNumberIdent: // number or ident which num-start
			for cs.IsDigit(s.peek(0)) {
				c = s.next()
			}
			c = s.peek(0)

			if identMap[c] == 0 {
				// Can't be identifier
				state = MyLexIntOrReal
				break
			}

			if c == 'e' || c == 'E' {
				s.next()
				var ok bool
				if retstate, ok = s.scanFloat(lval, &c); ok {
					return
				}
			} else if c == 'x' && (s.pos-s.tokStart) == 1 && s.buf[s.tokStart] == '0' {
				s.next()
				// 0xdddd number
				for c = s.next(); cs.IsXdigit(c); c = s.next() {
				}

				if s.pos-s.tokStart >= 4 && identMap[c] == 0 {
					retstate = HEX_NUM
					return
				}

				s.back()
			} else if c == 'b' && s.pos-s.tokStart == 1 && s.buf[s.tokStart] == '0' {
				s.next()
				// binary number 0bxxxx
				for c = s.next(); cs.IsXdigit(c); c = s.next() {
				}

				if s.pos-s.tokStart >= 4 && identMap[c] == 0 {
					retstate = BIN_NUM
					return
				}
				s.back()
			}
			fallthrough
		case MyLexIdentStart:
			retstate, lval.bytes = s.getPureIdentifier()
			return
		case MyLexUserVariableDelimiter:
			quoteChar := c

			for c = s.next(); c != EOF; c = s.next() {
				if c == NamesSepChar {
					break
				}

				if c == quoteChar && s.peek(0) != quoteChar {
					break
				}
			}

			if c == EOF {
				retstate = ABORT_SYM
				return
			}

			s.nextState = MyLexStart
			lval.bytes = s.buf[s.tokStart:s.pos]
			retstate = IDENT_QUOTED
			return

		case MyLexIntOrReal:
			if c != '.' {
				retstate = s.scanInt(lval)
				return
			}
			s.next()
			fallthrough
		case MyLexReal:
			for cs.IsDigit(s.peek(0)) {
				c = s.next()
			}
			c = s.peek(0)

			if c == 'e' || c == 'E' {
				s.next() // skip for 'e'/'E'
				var ok bool
				if retstate, ok = s.scanFloat(lval, &c); ok {
					return
				}

				state = MyLexChar
				break
			}

			lval.bytes = s.buf[s.tokStart:s.pos]
			retstate = DECIMAL_NUM
			return
		case MyLexHexNumber:
			s.next() // skip '
			for c = s.next(); cs.IsXdigit(c); c = s.next() {
			}

			length = s.pos - s.tokStart

			if (length&1) == 0 || c != '\'' {
				retstate = ABORT_SYM
				return
			}

			lval.bytes = s.buf[s.tokStart:s.pos]
			retstate = HEX_NUM
			return

		case MyLexBinNumber:
			s.next()
			for c = s.next(); c == '0' || c == '1'; c = s.next() {
			}

			length = s.pos - s.tokStart
			if c != '\'' {
				retstate = ABORT_SYM
				return
			}

			s.next()

			retstate = BIN_NUM
			return
		case MyLexCmpOp:
			if stateMap[s.peek(0)] == MyLexCmpOp || stateMap[s.peek(0)] == MyLexLongCmpOp {
				s.next()
			}

			var ok bool
			if retstate, ok = findKeywords(s.buf[s.tokStart:s.pos], false); ok {
				s.nextState = MyLexStart
				return
			}
			state = MyLexChar

		case MyLexLongCmpOp:
			if stateMap[s.peek(0)] == MyLexCmpOp || stateMap[s.peek(0)] == MyLexLongCmpOp {
				s.next()
				if stateMap[s.peek(0)] == MyLexCmpOp {
					s.next()
				}
			}

			var ok bool
			if retstate, ok = findKeywords(s.buf[s.tokStart:s.pos], false); ok {
				s.nextState = MyLexStart
				return
			}
			state = MyLexChar

		case MyLexBool:
			if c != s.peek(0) {
				state = MyLexChar
			} else {
				s.next()
				retstate, _ = findKeywords(s.buf[s.tokStart:s.tokStart+2], false)
				s.nextState = MyLexStart
				return
			}

		case MyLexStringOrDelimiter:
			if s.mode.ModeAnsiQuotes {
				state = MyLexUserVariableDelimiter
				break
			}
			fallthrough
		case MyLexString:
			b, err := s.scanText()
			if err != nil {
				s.Error(err.Error())
				retstate = ABORT_SYM
			} else {
				lval.bytes = b
				retstate = TEXT_STRING
			}
			return
		case MyLexComment:
			c = s.next()
			n := s.peek(0)
			for c != '\n' && !(c == '\r' && n != '\n') {
				c = s.next()
				n = s.peek(0)
			}

			s.back() // Safety against eof
			state = MyLexStart
		case MyLexLongComment:
			if s.peek(0) != '*' {
				state = MyLexChar
				break
			}

			s.next() // skip '*'
			if s.peek(0) == '!' {
				version := MysqlVersionID
				s.next() // skip '!'
				state = MyLexStart
				if cs.IsDigit(s.peek(0)) {
					start := s.pos
					s.next() // skip first digit
					for c = s.peek(0); cs.IsDigit(c); c = s.peek(0) {
						s.next()
					}

					i, err := strconv.Atoi(string(s.buf[start:s.pos]))
					if err != nil {
						s.Error(err.Error())
						return ABORT_SYM
					}
					version = i
				}

				if version <= MysqlVersionID {
					s.inComment = 1
					break
				}
			}

			// scan util match `*/`
			for c = s.next(); c != EOF && !(c == '*' && s.peek(0) == '/'); c = s.next() {
			}

			if c == '*' {
				s.next() // skip for '*/'
			}

			state = MyLexStart

		case MyLexEndLongComment:
			if s.inComment != 0 && s.peek(0) == '/' {
				s.next()
				s.inComment = 0
				state = MyLexStart
			} else {
				state = MyLexChar
			}
		case MyLexSetVar:
			if s.peek(0) != '=' {
				state = MyLexChar
			} else {
				s.next()
				retstate = SET_VAR
				return
			}

		case MyLexSemicolon:
			if s.peek(0) != 0 {
				state = MyLexChar
				break
			}
			fallthrough
		case MyLexEol:
			if s.pos >= len(s.buf) {
				s.nextState = MyLexEnd
				retstate = END_OF_INPUT
				return
			}

			state = MyLexChar
		case MyLexEnd:
			s.nextState = MyLexEnd
			retstate = 0
			return
		case MyLexRealOrPoint:
			if cs.IsDigit(s.peek(0)) {
				state = MyLexReal
			} else {
				state = MyLexIdentSep
				s.back()
			}
		case MyLexUserEnd: // end '@' of user@hostname
			switch stateMap[s.peek(0)] {
			case MyLexString, MyLexUserVariableDelimiter, MyLexStringOrDelimiter:
			case MyLexUserEnd:
				s.nextState = MyLexSystemVar
			default:
				s.nextState = MyLexHostname
			}

			retstate = int('@')
			return

		case MyLexHostname:
			for c = s.next(); cs.IsAlnum(c) || c == '.' || c == '_' || c == '$'; c = s.next() {
			}

			retstate = LEX_HOSTNAME
			return
		case MyLexSystemVar:
			s.next()
			s.nextState = func() uint {
				if stateMap[s.peek(0)] == MyLexUserVariableDelimiter {
					return MyLexOperatorOrIdent
				}
				return MyLexIdentOrKeyword
			}()

			retstate = int('@')
			return
		case MyLexIdentOrKeyword:
			resultState = 0
			c = s.next()
			for identMap[c] != 0 {
				resultState |= int(c)
				c = s.next()
			}

			if resultState&0x80 != 0 {
				resultState = IDENT_QUOTED
			} else {
				resultState = IDENT
			}

			if c == '.' {
				s.nextState = MyLexIdentSep
			}

			length = s.pos - s.tokStart - 1
			if length == 0 {
				retstate = ABORT_SYM
				return
			}

			val := s.buf[s.tokStart : s.pos-1]
			var ok bool
			if retstate, ok = findKeywords(val, false); ok {
				s.back()
				return
			}

			lval.bytes = val
			retstate = resultState

		}

	}
	return
}

func (s *Lexer) next() (b byte) {
	if s.pos < len(s.buf) {
		b = s.buf[s.pos]
		s.pos++
	} else {
		b = EOF
	}
	return
}

func (s *Lexer) peek(pos int) (b byte) {
	if s.pos+pos < len(s.buf) {
		b = s.buf[s.pos+pos]
	} else {
		b = EOF
	}
	return

}

func (s *Lexer) back() {
	s.pos--
}

func (s *Lexer) head() (b byte) {
	return s.buf[s.pos-1]
}

func (s *Lexer) scanChar(lval *MySQLSymType) (int, byte) {
	s.next()

	var c byte
	for c = s.next(); c != 0 && c != '\''; c = s.next() {
	}

	if c != '\'' {
		return ABORT_SYM, c
	}

	lval.bytes = s.buf[s.tokStart:s.pos]

	return NCHAR_STRING, c

}

func (s *Lexer) scanIdentifier() (int, []byte) {
	identMap := s.charset.IdentMap
	c := s.peek(0)
	rs := int(c)
	for identMap[s.peek(0)] != 0 {
		rs |= int(c)
		c = s.next()
	}
	if rs&0x80 != 0 {
		rs = IDENT_QUOTED
	} else {
		rs = IDENT
	}
	idc := s.buf[s.tokStart:s.pos]
	start := s.pos

	c = s.peek(0)
	if start == s.pos && s.peek(0) == '.' && identMap[int(s.peek(0))] != 0 {
		s.nextState = MyLexIdentSep
	} else if ret, ok := findKeywords(idc, c == '('); ok {
		s.nextState = MyLexStart
		return ret, idc
	}

	if idc[0] == '_' && charset.IsValidCharsets(idc[1:]) {
		return UNDERSCORE_CHARSET, idc
	}

	return rs, idc

}

func (s *Lexer) scanInt(lval *MySQLSymType) int {
	length := s.pos - s.tokStart
	lval.bytes = s.buf[s.tokStart:s.pos]

	if length < LongLen {
		return NUM
	}

	neg := false
	start := s.tokStart
	if s.buf[start] == '+' {
		start++
		length--
	} else if s.buf[start] == '-' {
		start++
		length--
		neg = true
	}

	// ignore any '0' character
	for start < s.pos && s.buf[start] == '0' {
		start++
		length--
	}

	if length < LongLen {
		return NUM
	}

	var cmp []byte
	var smaller int
	var bigger int
	if neg {
		if length == LongLen {
			cmp = SignedLong[1:len(SignedLong)]
			smaller = NUM
			bigger = LONG_NUM
		} else if length < SignedLonglongLen {
			return LONG_NUM
		} else if length > SignedLonglongLen {
			return DECIMAL_NUM
		}
		cmp = SignedLonglong[1:len(SignedLonglong)]
		smaller = LONG_NUM
		bigger = DECIMAL_NUM
	} else {
		if length == LongLen {
			cmp = Long
			smaller = NUM
			bigger = LONG_NUM
		} else if length < LonglongLen {
			return LONG_NUM
		} else if length > LonglongLen {
			if length > UnsignedLonglongLen {
				return DECIMAL_NUM
			}
			cmp = UnsignedLonglong
			smaller = ULONGLONG_NUM
			bigger = DECIMAL_NUM
		} else {
			cmp = Longlong
			smaller = LONG_NUM
			bigger = ULONGLONG_NUM
		}
	}

	idx := 0
	for idx < len(cmp) && cmp[idx] == s.buf[start] {
		idx++
		start++
	}

	if idx == len(cmp) {
		return smaller
	}

	if s.buf[start] <= cmp[idx] {
		return smaller
	}
	return bigger
}

func (s *Lexer) scanFloat(lval *MySQLSymType, c *byte) (int, bool) {
	cs := s.charset

	// try match (+|-)? digit+
	if s.peek(0) == '+' || s.peek(0) == '-' {
		// ignore this char
		s.next()
	}

	// at least we have 1 digit-char
	if cs.IsDigit(s.peek(0)) {
		for ; cs.IsDigit(s.peek(0)); s.next() {
		}

		lval.bytes = s.buf[s.tokStart:s.pos]
		return FLOAT_NUM, true
	}

	return 0, false
}

func (s *Lexer) getPureIdentifier() (int, []byte) {
	identMap := s.charset.IdentMap
	c := s.peek(0)
	rs := int(c)

	for identMap[s.peek(0)] != 0 {
		rs |= int(c)
		c = s.next()
	}

	if rs&0x80 != 0 {
		rs = IDENT_QUOTED
	} else {
		rs = IDENT
	}

	if s.peek(0) == '.' && identMap[int(s.peek(1))] != 0 {
		s.nextState = MyLexIdentSep
	}

	return rs, s.buf[s.tokStart:s.pos]
}

func (s *Lexer) scanText() ([]byte, error) {
	var dq bool
	var sep byte

	if sep = s.head(); sep == '"' {
		dq = true
	}
	for s.pos < len(s.buf) {
		c := s.next()
		if c == '\\' && !s.mode.ModeNoBackslashEscapes {
			if s.peek(0) == EOF {
				return nil, ErrStringFormat
			}
			s.next() // skip next cha
		} else if matchQuote(c, dq) {
			s.next()
			continue
		}
		return s.buf[s.tokStart:s.pos], nil

	}
	return nil, ErrStringFormat

}

func (s *Lexer) Error(err string) {
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	fmt.Fprintf(buf, "%s at position %v near '%s'", err, s.pos, string(s.buf[s.tokStart:s.pos]))
	s.LastError = buf.String()
}

func matchQuote(c byte, doubleQuote bool) bool {
	if doubleQuote {
		return c == '"'
	}

	return c == '\''
}
