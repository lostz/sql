package sql

type Statement interface {
	IsDDL() bool
	IsDML() bool
}

func setParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Lexer).ParseTree = stmt
}
