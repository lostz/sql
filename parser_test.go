package sql

import (
	"log"
	"testing"
)

func TestPase(t *testing.T) {
	sql := "db.test"
	stmt, err := Parse(sql)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(stmt)

}
