sql.go: sql_yacc.yy
		goyacc -o sql_yacc.go -p MySQL sql_yacc.yy
			gofmt -w sql_yacc.go

clean:
		rm -f y.output sql_yacc.g
