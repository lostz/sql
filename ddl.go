package sql

// AlterTableStmt  change  structure of a table.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-table.html
type AlterTableStmt struct {
	Table Table
}

func (*AlterTableStmt) Statement() {}

type AlterDatabaseStmt struct {
	Name []byte
}

func (*AlterDatabaseStmt) Statement() {}

type AlterProcedureStmt struct {
	Procedure Table
}

func (*AlterProcedureStmt) Statement() {}

type AlterFunctionStmt struct {
	Function Table
}

func (*AlterFunctionStmt) Statement() {}

type AlterViewStmt struct {
	View *TableIdent
	As   SelectStmt
}

func (*AlterViewStmt) Statement() {}

type AlterEventStmt struct {
	Event  Table
	Rename *Spname
}

func (*AlterEventStmt) Statement() {}

type AlterTablespaceStmt struct {
}

func (*AlterTablespaceStmt) Statement() {}

type AlterLogfileStmt struct{}

func (*AlterLogfileStmt) Statement() {}

type AlterServerStmt struct{}

func (*AlterServerStmt) Statement() {}

type AlterUserStmt struct{}

func (*AlterUserStmt) Statement() {}

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
	View *TableIdent
	As   SelectStmt
}

func (*CreateViewStmt) Statement() {}

type CreateTriggerStmt struct {
	Trigger Table
}

func (*CreateTriggerStmt) Statement() {}

type CreateProcedureStmt struct {
	Procedure Table
}

func (*CreateProcedureStmt) Statement() {}

type CreateFunctionStmt struct {
	Function Table
}

func (*CreateFunctionStmt) Statement() {}

type CreateUDFStmt struct {
	Function *TableIdent
}

func (*CreateUDFStmt) Statement() {}

type CreateEventStmt struct {
	Event Table
}

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
	View *TableIdent
	As   SelectStmt
}
type EventTail struct {
	Event Table
}

type TriggerTail struct {
	Trigger Table
}

type SfTail struct {
	Function Table
}

type SpTail struct {
	Procedure Table
}

type DropTableStmt struct {
	Tables Tables
}

func (*DropTableStmt) Statement() {}

type DropIndexStmt struct {
	On Table
}

func (*DropIndexStmt) Statement() {}

type DropDatabaseStmt struct {
	Name []byte
}

func (*DropDatabaseStmt) Statement() {}

type DropFunctionStmt struct {
	Function *Spname
}

func (*DropFunctionStmt) Statement() {}

type DropProcedureStmt struct {
	Procedure *Spname
}

func (*DropProcedureStmt) Statement() {}

type DropUserStmt struct {
}

func (*DropUserStmt) Statement() {}

type DropViewStmt struct{}

func (*DropViewStmt) Statement() {}

type DropEventStmt struct{}

func (*DropEventStmt) Statement() {}

type DropTriggerStmt struct{}

func (*DropTriggerStmt) Statement() {}

type DropTablespaceStmt struct{}

func (*DropTablespaceStmt) Statement() {}

type DropLogfileStmt struct{}

func (*DropLogfileStmt) Statement() {}

type DropServerStmt struct{}

func (*DropServerStmt) Statement() {}

type InsertStmt struct {
	Table Table
}

func (*InsertStmt) Statement() {}

type ReplaceStmt struct {
	Table Table
}

func (*ReplaceStmt) Statement() {}

type UpdateStmt struct {
	Tables Tables
}

func (*UpdateStmt) Statement() {}

type DeleteStmt struct {
	Tables Tables
}

func (*DeleteStmt) Statement() {}

type TruncateTableStmt struct {
	Tables Tables
}

func (*TruncateTableStmt) Statement() {}
