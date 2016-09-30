package sql

type Table interface {
	GetSchemas() []string
}

//TableIdent table_ident
type TableIdent struct {
	Schema []byte
	Name   []byte
}

func (t *TableIdent) GetSchemas() []string {
	return []string{string(t.Schema)}
}
