package sql

type TableIndex struct {
	Table Table
}

type TableIndexes []*TableIndex

type UdfTail struct {
	Function *TableIdent
}

type LifeType int

const (
	LifeUnknown = 0
	LifeGlobal  = 1
	LifeLocal   = 2
	LifeSession = 3
)

type Variable struct {
	Life  LifeType
	Name  string
	Value Expr
}

func (*Variable) Expression()      {}
func (*Variable) ValueExpression() {}

type Variables []*Variable

func (Variables) Expression()      {}
func (Variables) ValueExpression() {}

type RenameTableStmt struct {
	ToList []*TableToTable
}

func (*RenameTableStmt) Statement() {}

type RenameUserStmt struct {
}

func (*RenameUserStmt) Statement() {}

type CacheIndexStmt struct {
	TableIndexList TableIndexes
}

func (*CacheIndexStmt) Statement() {}

type LoadIndexStmt struct {
	TableIndexList TableIndexes
}

func (*LoadIndexStmt) Statement() {}

type ShowType int

const (
	ShowDatabases = iota
	ShowTables
	ShowTriggers
	ShowEvents
	ShowTableStatus
	ShowPlugins
	ShowEngines
	ShowColumns
	ShowBinlogs
	ShowSlaveHosts
	ShowBinlogEvents
	ShowRelaylogEvents
	ShowKeys
	ShowPrivileges
	ShowWarnings
	ShowErrors
	ShowProfiles
	ShowStatus
	ShowProcesslist
	ShowFullProcesslist
	ShowVariables
	ShowCharset
	ShowCollation
	ShowGrants
	ShowTable
	ShowView
	ShowMasterStatus
	ShowSlaveStatus
	ShowProcedureStatus
	ShowFunction
	ShowTrigger
	ShowUser
)

type ShowStmt struct {
	Tp ShowType
}

func (*ShowStmt) Statement() {}

type UseStmt struct {
	Name []byte
}

func (*UseStmt) Statement() {}

type SetStmt struct {
	Variables Variables
}

func (*SetStmt) Statement() {}

type SetPasswordStmt struct {
}

func (*SetPasswordStmt) Statement() {}

type SetTransStmt struct{}

func (*SetTransStmt) Statement() {}

type ShutdownStmt struct{}

func (*ShutdownStmt) Statement() {}

type GrantStmt struct{}

func (*GrantStmt) Statement() {}

type SavePointStmt struct{}

func (*SavePointStmt) Statement() {}
