package sql

type DescribeStmt struct {
	Table   Table
	Command Statement
}

func (*DescribeStmt) Statement() {}

type FlushStmt struct {
}

func (*FlushStmt) Statement() {}

type ResetStmt struct{}

func (*ResetStmt) Statement() {}

type PurgeStmt struct{}

func (*PurgeStmt) Statement() {}

type KillStmt struct{}

func (*KillStmt) Statement() {}
