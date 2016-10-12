package sql

type HelpStmt struct{}

func (*HelpStmt) Statement() {}

type ChangeStmt struct{}

func (*ChangeStmt) Statement() {}

type SignalStmt struct{}

func (*SignalStmt) Statement() {}

type ResignalStmt struct{}

func (*ResignalStmt) Statement() {}

type DiagnosticsStmt struct{}

func (*DiagnosticsStmt) Statement() {}

type GeneratedColumnStmt struct{}

func (*GeneratedColumnStmt) Statement() {}

type ChecksumStmt struct{}

func (*ChecksumStmt) Statement() {}

type RepairStmt struct{}

func (*RepairStmt) Statement() {}

type AnalyzeStmt struct{}

func (*AnalyzeStmt) Statement() {}

type BinlogSmt struct{}

func (*BinlogSmt) Statement() {}

type CheckStmt struct{}

func (*CheckStmt) Statement() {}

type OptimizeStmt struct{}

func (*OptimizeStmt) Statement() {}
