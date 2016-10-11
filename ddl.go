package sql

// AlterTableStmt  change  structure of a table.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-table.html
type AlterTableStmt struct {
	Table Table
}

type CreateTableStmt struct {
	Table Table
}

func (*CreateTableStmt) Statement() {}

type CreateDatabaseStmt struct {
	Name []byte
}

func (*CreateDatabaseStmt) Statement() {}

type CreateIndexStmt struct {
}

func (*CreateIndexStmt) Statement() {}

type CreateViewStmt struct {
}

func (*CreateViewStmt) Statement() {}

type CreateTriggerStmt struct {
}

func (*CreateTriggerStmt) Statement() {}

type CreateProcedureStmt struct{}

func (*CreateProcedureStmt) Statement() {}

type CreateFunctionStmt struct {
}

func (*CreateFunctionStmt) Statement() {}

type CreateUDFStmt struct{}

func (*CreateUDFStmt) Statement() {}

type CreateEventStmt struct{}

func (*CreateEventStmt) Statement() {}

type CreateUserStmt struct{}

func (*CreateUserStmt) Statement() {}

type CreateLogStmt struct{}

func (*CreateLogStmt) Statement() {}

type CreateTablespaceStmt struct{}

func (*CreateTablespaceStmt) Statement() {}

type CreateServerStmt struct{}

func (*CreateServerStmt) Statement() {}

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
