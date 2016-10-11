package sql

type HelpStmt struct{}

func (*HelpStmt) Statement() {}

type ChangeStmt struct{}

func (*ChangeStmt) Statement() {}
