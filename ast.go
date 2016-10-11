package sql

type Statement interface {
	Statement()
}

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Lexer).ParseTree = stmt
}
