package sql

type TableIndex struct {
	Table Table
}

type UdfTail struct {
	Function TableIdent
}

type LifeType int

const (
	LifeUnknown = 0
	LifeGlobal  = 1
	LifeLocal   = 2
	LifeSession = 3
)

type Variable struct {
	Life LifeType
	Name string
}
