package sql

// AlterTableStmt  change  structure of a table.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-table.html
type AlterTableStmt struct {
	Table Table
}

type ViewTail struct {
	View TableIdent
	As   SelectStmt
}
type EventTail struct {
	Event TableIdent
}

type TriggerTail struct {
	Trigger TableIdent
}

type SfTail struct {
	Function TableIdent
}

type SpTail struct {
	Procedure TableIdent
}
