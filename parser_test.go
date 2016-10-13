package sql

import (
	"log"
	"testing"
)

func TestPase(t *testing.T) {
	sql := "SELECT * from t2.t3"
	stmt, err := Parse(sql)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch st := stmt.(type) {
	case SelectStatement:
		log.Println(st.GetSchemas())
		log.Println("select")
	default:
		log.Println("unknown %v ", st)
	}

}
