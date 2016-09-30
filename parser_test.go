package sql

import (
	"log"
	"testing"
)

func TestPase(t *testing.T) {
	sql := "alter table t1"
	stmt, err := Parse(sql)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(stmt.IsDDL())

}
