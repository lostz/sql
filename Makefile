build:
	goyacc -o sql_yacc.go -p MySQL sql_yacc.yy
	gofmt -w sql_yacc.go
.PHONY:  build
