package sql

import (
	"log"
	"testing"
)

func TestPase(t *testing.T) {
	sql := "SELECT DEPT, MAX(SALARY) AS MAXIMUM FROM STAFF GROUP BY DEPT limit 2"
	stmt, err := Parse(sql)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(sql)
	switch st := stmt.(type) {
	case SelectStatement:
		log.Println("type: %v", st)
		log.Println("Schemas:", st.GetSchemas())
		log.Println("group by %v", st.GetOrderBy())
		limit, err := st.GetLimit().Offset.ParseInt()
		log.Println(limit, err)
	case *SelectStmt:
		log.Println("group by %v", st.GetOrderBy())
	default:
		log.Println("unknown %v ", st)
	}

}
