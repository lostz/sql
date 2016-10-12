package sql

import "fmt"

// Table  represents a table
type Table interface {
	GetSchemas() []string
}

type Tables []Table

func (ts Tables) GetSchemas() []string {
	if ts == nil && len(ts) == 0 {
		return nil
	}
	var ret []string
	for _, v := range ts {
		if r := v.GetSchemas(); r != nil && len(r) != 0 {
			ret = append(ret, r...)
		}
	}
	return ret
}

type TableToTable struct {
	From Table
	To   Table
}

// TableIdent table_ident
type TableIdent struct {
	Schema []byte
	Name   []byte
}

// GetSchemas return table schema
func (t *TableIdent) GetSchemas() []string {
	return []string{string(t.Schema)}
}

type Spname struct {
	Schema []byte
	Name   []byte
}

func (s *Spname) GetSchemas() []string {
	return string(s.Schema)
}

type AliasedTable struct {
	TableOrSubQuery interface{}
	As              []byte
}

func (a *AliasedTable) GetSchemas() []string {
	if t, ok := a.TableOrSubQuery.(Table); ok {
		return t.GetSchemas()
	}
	if s, can := a.TableOrSubQuery.(*SubQueryStmt); can {
		return s.GetSchemas()
	}
	panic(fmt.Sprintf("alias table has no table_factor or subquery, element type[%T]", a.TableOrSubQuery))

}
