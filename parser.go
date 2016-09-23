package sql

import "errors"

//Parse sql
func Parse(sql string) (Statement, error) {
	lexer := NewLexer(sql)
	if MySQLParse(lexer) != 0 {
		return nil, errors.New(lexer.LastError)
	}
	return lexer.ParseTree, nil

}
