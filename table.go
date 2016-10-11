package sql

// Table  represents a table
type Table interface {
	GetSchemas() []string
}

type Tables []Table

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
