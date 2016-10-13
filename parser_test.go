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
	log.Println(sql)
	switch st := stmt.(type) {
	case SelectStatement:
		log.Println("type: select")
		log.Println("Schemas:", st.GetSchemas())
	default:
		log.Println("unknown %v ", st)
	}

}
