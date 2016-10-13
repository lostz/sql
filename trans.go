package sql

type BeginStmt struct{}

func (*BeginStmt) Statement() {}

type ExecuteStmt struct{}

func (*ExecuteStmt) Statement() {}

type PrepareStmt struct{}

func (*PrepareStmt) Statement() {}

type StartTransStmt struct{}

func (*StartTransStmt) Statement() {}

type RevokeStmt struct{}

func (*RevokeStmt) Statement() {}

type CommitStmt struct{}

func (*CommitStmt) Statement() {}

type RollBackStmt struct{}

func (*RollBackStmt) Statement() {}

type ReleaseStmt struct{}

func (*ReleaseStmt) Statement() {}

type LockStmt struct {
	Tables Tables
}

func (*LockStmt) Statement() {}

type UnlockStmt struct {
}

func (*UnlockStmt) Statement() {}
