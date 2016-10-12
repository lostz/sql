package sql

type JoinType int

const (
	// CrossJoin is cross join type.
	CrossJoin JoinType = iota + 1
	// LeftJoin is left Join type.
	LeftJoin
	// RightJoin is right Join type.
	RightJoin
)

// Join represents table join.
type Join struct {
	Left  Table
	Right Table
	Tp    JoinType
}

// GetSchemas return table schema
func (j *Join) GetSchemas() []string {
	if j.Left == nil {
		panic("join table must have left value")
	}
	if j.Right == nil {
		panic("join table must have left value")
	}
	l := j.Left.GetSchemas()
	r := j.Right.GetSchemas()
	if l == nil && r == nil {
		return nil
	} else if l == nil {
		return r
	} else if r == nil {
		return l
	}
	return append(l, r...)
}

type ValExpr interface {
	Expr()
}

type GroupBy struct {
	Items []ValExpr
}

type OrderBy struct {
	Items []ValExpr
}

type Limit struct {
	Offset []byte
	Count  []byte
}

type SelectLockType int

// Select lock types.
const (
	SelectLockNone SelectLockType = iota
	SelectLockForUpdate
	SelectLockInShareMode
)

type Select interface {
	Select()
	Statement()
}

type SelectStmt struct {
	From    Table
	GroupBy *GroupBy
	OrderBy *OrderBy
	Limit   *Limit
	LockTp  SelectLockType
}

func (*SelectStmt) Statement() {}
func (*SelectStmt) Select()    {}

type CallStmt struct {
	Spname *Spname
}

func (*CallStmt) Statement() {}

type UnionStmt struct {
	SelectList Select
	OrderBy    *OrderBy
	Limit      *Limit
}

func (*UnionStmt) Statement() {}
func (*UnionStmt) Select()    {}

type SubQueryStmt struct {
	SelectStmt Select
}

func (*SubQueryStmt) Statement()       {}
func (*SubQueryStmt) Select()          {}
func (*SubQueryStmt) Expression()      {}
func (*SubQueryStmt) ValueExpression() {}

type DoStmt struct {
}

func (*DoStmt) Statement() {}
