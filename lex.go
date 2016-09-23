package sql

import (
	"bufio"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/lostz/sql/charset"
)

const NAMES_SEP_CHAR byte = '\377'
const MYSQL_VERSION_ID = 50109

type lex struct {
	reader    *bufio.Reader
	buf       []byte
	pos       int
	lastPos   int
	state     uint
	nextState uint
	tokStart  int
	charset   *charset.CharsetInfo
	mode      mode
	inComment uint
	ParseTree Statement
	LastError string
}

type mode struct {
	MODE_ANSI_QUOTES          bool
	MODE_HIGH_NOT_PRECEDENCE  bool
	MODE_PIPES_AS_CONCAT      bool
	MODE_NO_BACKSLASH_ESCAPES bool
	MODE_IGNORE_SPACE         bool
}

func NewLexer(sql string) *lex {
	s := &lex{
		reader:  bufio.NewReader(strings.NewReader(sql)),
		buf:     []byte(sql),
		state:   MY_LEX_START,
		charset: charset.CSUtf8GeneralCli,
	}
	return s
}

func (s *lex) Lex(lval *MySQLSymType) (retstate int) {
	var resultState int
	var c byte
	var length int
	state := s.state
	cs := s.charset
	stateMap := cs.StateMap
	identMap := cs.IdentMap
	for {
		switch state {
		case MY_LEX_OPERATOR_OR_IDENT, MY_LEX_START:
			for c = s.next(); stateMap[c] == MY_LEX_SKIP; c = s.next() {

			}
			s.tokStart = s.pos - 1
			state = stateMap[c]
		case MY_LEX_ESCAPE:
			if s.next() == 'N' {
				retstate = NULL_SYM
				return
			}
			fallthrough
		case MY_LEX_CHAR, MY_LEX_SKIP:
			if c == '-' && s.peek(0) == '-' && (cs.IsSpace(s.peek(1)) || cs.IsCntrl(s.peek(1))) {
				state = MY_LEX_COMMENT
				break
			}
			s.pos = s.tokStart

			c = s.next()
			if c != ')' {
				s.nextState = MY_LEX_START
			}

			if c == ',' {
				s.tokStart = s.pos
			} else if c == '?' && identMap[s.peek(0)] == 0 {
				retstate = PARAM_MARKER
				return
			}
			retstate = int(c)
			return
		case MY_LEX_IDENT_OR_NCHAR:
			if s.peek(0) != '\'' {
				state = MY_LEX_IDENT
				break
			}

			retstate, c = s.scanChar(lval)
			return
		case MY_LEX_IDENT_OR_HEX:
			if s.peek(0) == '\'' {
				state = MY_LEX_HEX_NUMBER
				break
			}
			fallthrough
		case MY_LEX_IDENT_OR_BIN:
			if s.peek(0) == '\'' {
				state = MY_LEX_BIN_NUMBER
				break
			}
			fallthrough
		case MY_LEX_IDENT:
			retstate, lval.bytes = s.scanIdentifier()
			return
		case MY_LEX_IDENT_SEP: // Found ident before
			// And Now '.'
			c = s.next()
			s.nextState = MY_LEX_IDENT_START
			if identMap[s.peek(0)] == 0 {
				s.nextState = MY_LEX_START
			}
			retstate = int(c)
			return
		case MY_LEX_NUMBER_IDENT: // number or ident which num-start
			for cs.IsDigit(s.peek(0)) {
				c = s.next()
			}
			c = s.peek(0)

			if identMap[c] == 0 {
				// Can't be identifier
				state = MY_LEX_INT_OR_REAL
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
		case MY_LEX_IDENT_START:
			retstate, lval.bytes = s.getPureIdentifier()
			return
		case MY_LEX_USER_VARIABLE_DELIMITER:
			quoteChar := c

			for c = s.next(); c != EOF; c = s.next() {
				if c == NAMES_SEP_CHAR {
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

			s.nextState = MY_LEX_START
			lval.bytes = s.buf[s.tokStart:s.pos]
			retstate = IDENT_QUOTED
			return

		case MY_LEX_INT_OR_REAL:
			if c != '.' {
				retstate = s.scanInt(lval)
				return
			}
			s.next()
			fallthrough
		case MY_LEX_REAL:
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

				state = MY_LEX_CHAR
				break
			}

			lval.bytes = s.buf[s.tokStart:s.pos]
			retstate = DECIMAL_NUM
			return
		case MY_LEX_HEX_NUMBER:
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

		case MY_LEX_BIN_NUMBER:
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
		case MY_LEX_CMP_OP:
			if stateMap[s.peek(0)] == MY_LEX_CMP_OP || stateMap[s.peek(0)] == MY_LEX_LONG_CMP_OP {
				s.next()
			}

			var ok bool
			if retstate, ok = findKeywords(s.buf[s.tokStart:s.pos], false); ok {
				s.nextState = MY_LEX_START
				return
			}
			state = MY_LEX_CHAR

		case MY_LEX_LONG_CMP_OP:
			if stateMap[s.peek(0)] == MY_LEX_CMP_OP || stateMap[s.peek(0)] == MY_LEX_LONG_CMP_OP {
				s.next()
				if stateMap[s.peek(0)] == MY_LEX_CMP_OP {
					s.next()
				}
			}

			var ok bool
			if retstate, ok = findKeywords(s.buf[s.tokStart:s.pos], false); ok {
				s.nextState = MY_LEX_START
				return
			}
			state = MY_LEX_CHAR

		case MY_LEX_BOOL:
			if c != s.peek(0) {
				state = MY_LEX_CHAR
			} else {
				s.next()
				retstate, _ = findKeywords(s.buf[s.tokStart:s.tokStart+2], false)
				s.nextState = MY_LEX_START
				return
			}

		case MY_LEX_STRING_OR_DELIMITER:
			if s.mode.MODE_ANSI_QUOTES {
				state = MY_LEX_USER_VARIABLE_DELIMITER
				break
			}
			fallthrough
		case MY_LEX_STRING:
			b, err := s.scanText()
			if err != nil {
				s.Error(err.Error())
				retstate = ABORT_SYM
			} else {
				lval.bytes = b
				retstate = TEXT_STRING
			}
			return
		case MY_LEX_COMMENT:
			c = s.next()
			n := s.peek(0)
			for c != '\n' && !(c == '\r' && n != '\n') {
				c = s.next()
				n = s.peek(0)
			}

			s.back() // Safety against eof
			state = MY_LEX_START
		case MY_LEX_LONG_COMMENT:
			if s.peek(0) != '*' {
				state = MY_LEX_CHAR
				break
			}

			s.next() // skip '*'
			if s.peek(0) == '!' {
				version := MYSQL_VERSION_ID
				s.next() // skip '!'
				state = MY_LEX_START
				if cs.IsDigit(s.peek(0)) {
					start := s.pos
					s.next() // skip first digit
					for c = s.peek(0); cs.IsDigit(c); c = s.peek(0) {
						s.next()
					}

					if i, err := strconv.Atoi(string(s.buf[start:s.pos])); err != nil {
						s.Error(err.Error())
						return ABORT_SYM
					} else {
						version = i
					}
				}

				if version <= MYSQL_VERSION_ID {
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

			state = MY_LEX_START

		case MY_LEX_END_LONG_COMMENT:
			if s.inComment != 0 && s.peek(0) == '/' {
				s.next()
				s.inComment = 0
				state = MY_LEX_START
			} else {
				state = MY_LEX_CHAR
			}
		case MY_LEX_SET_VAR:
			if s.peek(0) != '=' {
				state = MY_LEX_CHAR
			} else {
				s.next()
				retstate = SET_VAR
				return
			}

		case MY_LEX_SEMICOLON:
			if s.peek(0) != 0 {
				state = MY_LEX_CHAR
				break
			}
			fallthrough
		case MY_LEX_EOL:
			if s.pos >= len(s.buf) {
				s.nextState = MY_LEX_END
				retstate = END_OF_INPUT
				return
			}

			state = MY_LEX_CHAR
		case MY_LEX_END:
			s.nextState = MY_LEX_END
			retstate = 0
			return
		case MY_LEX_REAL_OR_POINT:
			if cs.IsDigit(s.peek(0)) {
				state = MY_LEX_REAL
			} else {
				state = MY_LEX_IDENT_SEP
				s.back()
			}
		case MY_LEX_USER_END: // end '@' of user@hostname
			switch stateMap[s.peek(0)] {
			case MY_LEX_STRING, MY_LEX_USER_VARIABLE_DELIMITER, MY_LEX_STRING_OR_DELIMITER:
			case MY_LEX_USER_END:
				s.nextState = MY_LEX_SYSTEM_VAR
			default:
				s.nextState = MY_LEX_HOSTNAME
			}

			retstate = int('@')
			return

		case MY_LEX_HOSTNAME:
			for c = s.next(); cs.IsAlnum(c) || c == '.' || c == '_' || c == '$'; c = s.next() {
			}

			retstate = LEX_HOSTNAME
			return
		case MY_LEX_SYSTEM_VAR:
			s.next()
			s.nextState = func() uint {
				if stateMap[s.peek(0)] == MY_LEX_USER_VARIABLE_DELIMITER {
					return MY_LEX_OPERATOR_OR_IDENT
				} else {
					return MY_LEX_IDENT_OR_KEYWORD
				}
			}()

			retstate = int('@')
			return
		case MY_LEX_IDENT_OR_KEYWORD:
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
				s.nextState = MY_LEX_IDENT_SEP
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

func (s *lex) next() (b byte) {
	if s.pos < len(s.buf) {
		b = s.buf[s.pos]
		s.pos++
	} else {
		b = EOF
	}
	return
}

func (s *lex) peek(pos int) (b byte) {
	if s.pos+pos < len(s.buf) {
		b = s.buf[s.pos+1]
	} else {
		b = EOF
	}
	return

}

func (s *lex) back() {
	s.pos -= 1
}

func (s *lex) head() (b byte) {
	return s.buf[s.pos-1]
}

func (s *lex) scanChar(lval *MySQLSymType) (int, byte) {
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

func (s *lex) scanIdentifier() (int, []byte) {
	identMap := s.charset.IdentMap
	log.WithFields(log.Fields{"identMap": identMap, "buf": s.buf}).Info()
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
		s.nextState = MY_LEX_IDENT_SEP
	} else if ret, ok := findKeywords(idc, c == '('); ok {
		s.nextState = MY_LEX_START
		return ret, idc
	}

	if idc[0] == '_' && charset.IsValidCharsets(idc[1:]) {
		return UNDERSCORE_CHARSET, idc
	}

	return rs, idc

}

func (s *lex) scanInt(lval *MySQLSymType) int {
	length := s.pos - s.tokStart
	lval.bytes = s.buf[s.tokStart:s.pos]

	if length < LONG_LEN {
		return NUM
	}

	neg := false
	start := s.tokStart
	if s.buf[start] == '+' {
		start += 1
		length -= 1
	} else if s.buf[start] == '-' {
		start += 1
		length -= 1
		neg = true
	}

	// ignore any '0' character
	for start < s.pos && s.buf[start] == '0' {
		start += 1
		length -= 1
	}

	if length < LONG_LEN {
		return NUM
	}

	var cmp []byte
	var smaller int
	var bigger int
	if neg {
		if length == LONG_LEN {
			cmp = SIGNED_LONG[1:len(SIGNED_LONG)]
			smaller = NUM
			bigger = LONG_NUM
		} else if length < SIGNED_LONGLONG_LEN {
			return LONG_NUM
		} else if length > SIGNED_LONGLONG_LEN {
			return DECIMAL_NUM
		} else {
			cmp = SIGNED_LONGLONG[1:len(SIGNED_LONGLONG)]
			smaller = LONG_NUM
			bigger = DECIMAL_NUM
		}
	} else {
		if length == LONG_LEN {
			cmp = LONG
			smaller = NUM
			bigger = LONG_NUM
		} else if length < LONGLONG_LEN {
			return LONG_NUM
		} else if length > LONGLONG_LEN {
			if length > UNSIGNED_LONGLONG_LEN {
				return DECIMAL_NUM
			}
			cmp = UNSIGNED_LONGLONG
			smaller = ULONGLONG_NUM
			bigger = DECIMAL_NUM
		} else {
			cmp = LONGLONG
			smaller = LONG_NUM
			bigger = ULONGLONG_NUM
		}
	}

	idx := 0
	for idx < len(cmp) && cmp[idx] == s.buf[start] {
		idx += 1
		start += 1
	}

	if idx == len(cmp) {
		return smaller
	}

	if s.buf[start] <= cmp[idx] {
		return smaller
	}
	return bigger
}

func (s *lex) scanFloat(lval *MySQLSymType, c *byte) (int, bool) {
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

func (s *lex) getPureIdentifier() (int, []byte) {
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
		s.nextState = MY_LEX_IDENT_SEP
	}

	return rs, s.buf[s.tokStart:s.pos]
}

func (s *lex) scanText() ([]byte, error) {
	var dq bool
	var sep byte

	if sep = s.head(); sep == '"' {
		dq = true
	}
	for s.pos < len(s.buf) {
		c := s.next()
		if c == '\\' && !s.mode.MODE_NO_BACKSLASH_ESCAPES {
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

func (s *lex) Error(err string) {
	s.LastError = err
}

func matchQuote(c byte, doubleQuote bool) bool {
	if doubleQuote {
		return c == '"'
	}

	return c == '\''
}
