package sql

import (
	"bytes"
	"fmt"
	"strings"
)

const eofChar = 0x100

type lex struct {
	InStream    *strings.Reader
	ForceEOF    bool
	lastChar    uint16
	Position    int
	posVarIndex int
}

func NewLex(sql string) *lex {
	return &lex{InStream: strings.NewReader(sql)}
}

func (l *lex) Lex(lval *yySymType) int {
	typ, val := l.Scan()

}

func (l *lex) Scan() (int, []byte) {
	if l.ForceEOF {
		return 0, nil
	}

	if l.lastChar == 0 {
		l.next()
	}
	l.skipBlank()
	switch ch := l.lastChar; {
	case isLetter(ch):
		return l.scanIdentifier()
	case ch == ':':
		return l.scanBindVar()
	default:
		l.next()
		switch ch {
		case eofChar:
			return 0, nil
		case '=', ',', ';', '(', ')', '+', '*', '%', '&', '|', '^', '~':
			return int(ch), nil
		case '?':
			l.posVarIndex++
			buf := new(bytes.Buffer)
			fmt.Fprintf(buf, ":v%d", l.posVarIndex)
			return TEXT_STRING, buf.Bytes()
		case '.':
			if isDigit(l.lastChar) {
				return l.scanNumber(true)
			}
			return int(ch), nil
		case '/':
			switch l.lastChar {
			case '/':
				l.next()
				return l.scanCommentType1("//")
			case '*':
				l.next()
				return l.scanCommentType2()
			default:
				return int(ch), nil
			}
		case '-':
			if l.lastChar == '-' {
				l.next()
				return l.scanCommentType1("--")
			}
			return int(ch), nil
		case '<':
			switch l.lastChar {
			case '>':
				l.next()
				return NE, nil
			case '<':
				l.next()
				return SHIFT_LEFT, nil
			case '=':
				l.next()
				switch l.lastChar {
				case '>':
					l.next()
					return EQUAL_SYM, nil
				default:
					return LE, nil
				}
			default:
				return int(ch), nil
			}
		case '>':
			switch l.lastChar {
			case '=':
				l.next()
				return GE, nil
			case '>':
				l.next()
				return SHIFT_RIGHT, nil
			default:
				return int(ch), nil
			}
		case '!':
			if l.lastChar == '=' {
				l.next()
				return NE, nil
			}
			return LEX_ERROR, []byte("!")
		case '\'', '"':
			return l.scanString(ch, TEXT_STRING)
		case '`':
			return l.scanLiteralIdentifier()
		default:
			return LEX_ERROR, []byte{byte(ch)}

		}

	}

}

func (l *lex) next() {
	if ch, err := l.InStream.ReadByte(); err != nil {
		// Only EOF is possible.
		l.lastChar = eofChar
	} else {
		l.lastChar = uint16(ch)
	}
	l.Position++
}
func (l *lex) skipBlank() {
	ch := l.lastChar
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		l.next()
		ch = l.lastChar
	}
}

func (l *lex) scanIdentifier() (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteByte(byte(l.lastChar))
	for tkn.next(); isLetter(l.lastChar) || isDigit(l.lastChar); l.next() {
		buffer.WriteByte(byte(l.lastChar))
	}
	lowered := bytes.ToLower(buffer.Bytes())
	loweredStr := string(lowered)
	if keywordID, found := keywords[loweredStr]; found {
		return keywordID, lowered
	}
	if loweredStr == "dual" {
		return ident, lowered
	}
	return ident, buffer.Bytes()

}

func (l *lex) scanBindVar() (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteByte(byte(l.lastChar))
	token := TEXT_STRING
	l.next()
	if !isLetter(l.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for isLetter(l.lastChar) || isDigit(l.lastChar) || l.lastChar == '.' {
		buffer.WriteByte(byte(l.lastChar))
		tkn.next()
	}
	return token, buffer.Bytes()

}

func (l *lex) scanNumber(seenDecimalPoint bool) (int, []byte) {
	buffer := &bytes.Buffer{}
	if seenDecimalPoint {
		buffer.WriteByte('.')
		l.scanMantissa(10, buffer)
		goto exponent
	}
	if l.lastChar == '0' {
		// int or float
		l.consumeNext(buffer)
		if l.lastChar == 'x' || l.lastChar == 'X' {
			// hexadecimal int
			l.consumeNext(buffer)
			l.scanMantissa(16, buffer)
		} else {
			seenDecimalDigit := false
			l.scanMantissa(8, buffer)
			if l.lastChar == '8' || l.lastChar == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				l.scanMantissa(10, buffer)
			}
			if l.lastChar == '.' || l.lastChar == 'e' || l.lastChar == 'E' {
				goto fraction
			}
			if seenDecimalDigit {
				return LEX_ERROR, buffer.Bytes()
			}
		}
		goto exit
	}
	l.scanMantissa(10, buffer)

fraction:
	if l.lastChar == '.' {
		l.consumeNext(buffer)
		l.scanMantissa(10, buffer)
	}
exponent:
	if l.lastChar == 'e' || l.lastChar == 'E' {
		l.consumeNext(buffer)
		if l.lastChar == '+' || l.lastChar == '-' {
			l.consumeNext(buffer)
		}
		l.scanMantissa(10, buffer)
	}
exit:
	return NUMBER_SYM, buffer.Bytes()

}

func (l *lex) scanCommentType1(prefix string) (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteString(prefix)
	for l.lastChar != eofChar {
		if l.lastChar == '\n' {
			l.consumeNext(buffer)
			break
		}
		l.consumeNext(buffer)
	}
	return COMMENT_SYM, buffer.Bytes()
}
func (l *lex) scanCommentType2() (int, []byte) {
	buffer := &bytes.Buffer{}
	buffer.WriteString("/*")
	for {
		if l.lastChar == '*' {
			l.consumeNext(buffer)
			if l.lastChar == '/' {
				l.consumeNext(buffer)
				break
			}
			continue
		}
		if l.lastChar == eofChar {
			return LEX_ERROR, buffer.Bytes()
		}
		l.consumeNext(buffer)
	}
	return COMMENT_SYM, buffer.Bytes()
}

func (l *lex) scanString(delim uint16, typ int) (int, []byte) {
	buffer := &bytes.Buffer{}
	for {
		ch := tkn.lastChar
		tkn.next()
		if ch == delim {
			if tkn.lastChar == delim {
				tkn.next()
			} else {
				break
			}
		}
	}
}

func isLetter(ch uint16) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}
func digitVal(ch uint16) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}
func isDigit(ch uint16) bool {
	return '0' <= ch && ch <= '9'
}
