package sql

type BeginStmt struct{}

func (*BeginStmt) Statement() {}

type ExecuteStmt struct{}

func (*ExecuteStmt) Statement() {}

type PrepareStmt struct{}

func (*PrepareStmt) Statement() {}
