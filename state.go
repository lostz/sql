package sql

import "fmt"

const (
	MyLexStart = iota
	MyLexChar
	MyLexIdent
	MyLexIdentSep
	MyLexIdentStart
	MyLexReal
	MyLexHexNumber
	MyLexBinNumber
	MyLexCmpOp
	MyLexLongCmpOp
	MyLexString
	MyLexComment
	MyLexEnd
	MyLexOperatorOrIdent
	MyLexNumberIdent
	MyLexIntOrReal
	MyLexRealOrPoint
	MyLexBool
	MyLexEol
	MyLexEscape
	MyLexLongComment
	MyLexEndLongComment
	MyLexSemicolon
	MyLexSetVar
	MyLexUserEnd
	MyLexHostname
	MyLexSkip
	MyLexUserVariableDelimiter
	MyLexSystemVar
	MyLexIdentOrKeyword
	MyLexIdentOrHex
	MyLexIdentOrBin
	MyLexIdentOrNchar
	MyLexStringOrDelimiter
)

var statusMap map[uint]string = map[uint]string{

	MyLexStart:                 "MY_LEX_START",
	MyLexChar:                  "MY_LEX_CHAR",
	MyLexIdent:                 "MY_LEX_IDENT",
	MyLexIdentSep:              "MY_LEX_IDENT_SEP",
	MyLexIdentStart:            "MY_LEX_IDENT_START",
	MyLexReal:                  "MY_LEX_REAL",
	MyLexHexNumber:             "MY_LEX_HEX_NUMBER",
	MyLexBinNumber:             "MY_LEX_BIN_NUMBER",
	MyLexCmpOp:                 "MY_LEX_CMP_OP",
	MyLexLongCmpOp:             "MY_LEX_LONG_CMP_OP",
	MyLexString:                "MY_LEX_STRING",
	MyLexComment:               "MY_LEX_COMMENT",
	MyLexEnd:                   "MY_LEX_END",
	MyLexOperatorOrIdent:       "MY_LEX_OPERATOR_OR_IDENT",
	MyLexNumberIdent:           "MY_LEX_NUMBER_IDENT",
	MyLexIntOrReal:             "MY_LEX_INT_OR_REAL",
	MyLexRealOrPoint:           "MY_LEX_REAL_OR_POINT",
	MyLexBool:                  "MY_LEX_BOOL",
	MyLexEol:                   "MY_LEX_EOL",
	MyLexEscape:                "MY_LEX_ESCAPE",
	MyLexLongComment:           "MY_LEX_LONG_COMMENT",
	MyLexEndLongComment:        "MY_LEX_END_LONG_COMMENT",
	MyLexSemicolon:             "MY_LEX_SEMICOLON",
	MyLexSetVar:                "MY_LEX_SET_VAR",
	MyLexUserEnd:               "MY_LEX_USER_END",
	MyLexHostname:              "MY_LEX_HOSTNAME",
	MyLexSkip:                  "MY_LEX_SKIP",
	MyLexUserVariableDelimiter: "MY_LEX_USER_VARIABLE_DELIMITER",
	MyLexSystemVar:             "MY_LEX_SYSTEM_VAR",
	MyLexIdentOrKeyword:        "MY_LEX_IDENT_OR_KEYWORD",
	MyLexIdentOrHex:            "MY_LEX_IDENT_OR_HEX",
	MyLexIdentOrBin:            "MY_LEX_IDENT_OR_BIN",
	MyLexIdentOrNchar:          "MY_LEX_IDENT_OR_NCHAR",
	MyLexStringOrDelimiter:     "MY_LEX_STRING_OR_DELIMITER",
}

//GetLexString  uint to string
func GetLexString(which uint) string {
	if v, ok := statusMap[which]; ok {
		return v
	}

	return fmt.Sprint("Unknow Status[%d]", which)
}
