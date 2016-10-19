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
	Offset NumVal
	Count  NumVal
}

type SelectLockType int

// Select lock types.
const (
	SelectLockNone SelectLockType = iota
	SelectLockForUpdate
	SelectLockInShareMode
)

type SelectStatement interface {
	Select()
	Statement()
	GetGroupBy() *GroupBy
	GetOrderBy() *OrderBy
	GetLimit() *Limit
	GetSchemas() []string
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
func (s *SelectStmt) GetGroupBy() *GroupBy {
	return s.GroupBy
}

func (s *SelectStmt) GetOrderBy() *OrderBy {
	return s.OrderBy
}

func (s *SelectStmt) GetLimit() *Limit {
	return s.Limit
}

func (s *SelectStmt) GetSchemas() []string {
	return s.From.GetSchemas()
}

type CallStmt struct {
	Spname *Spname
}

func (*CallStmt) Statement() {}

type UnionStmt struct {
	Left  SelectStatement
	Right SelectStatement
}

func (*UnionStmt) Statement() {}
func (*UnionStmt) Select()    {}

func (u *UnionStmt) GetGroupBy() *GroupBy {
	return nil
}

func (u *UnionStmt) GetOrderBy() *OrderBy {
	return nil
}

func (u *UnionStmt) GetLimit() *Limit {
	return nil
}

func (u *UnionStmt) GetSchemas() []string {
	if u.Left == nil {
		panic("union must have left select statement")
	}
	if u.Right == nil {
		panic("union must have right select statement")
	}
	l := u.Left.GetSchemas()
	r := u.Right.GetSchemas()
	return append(l, r...)

}

type SubQueryStmt struct {
	SelectStmt SelectStatement
}

func (*SubQueryStmt) Statement() {}
func (*SubQueryStmt) Select()    {}
func (s *SubQueryStmt) GetGroupBy() *GroupBy {
	return s.SelectStmt.GetGroupBy()
}

func (s *SubQueryStmt) GetOrderBy() *OrderBy {
	return s.SelectStmt.GetOrderBy()
}

func (s *SubQueryStmt) GetLimit() *Limit {
	return s.SelectStmt.GetLimit()
}

func (s *SubQueryStmt) GetSchemas() []string {
	return s.SelectStmt.GetSchemas()
}
func (*SubQueryStmt) Expression()      {}
func (*SubQueryStmt) ValueExpression() {}

type DoStmt struct {
}

func (*DoStmt) Statement() {}

type LoadStmt struct {
}

func (*LoadStmt) Statement() {}

type HandlerStmt struct {
}

func (*HandlerStmt) Statement() {}

type ParenSelect struct {
	SelectStmt SelectStatement
}

func (*ParenSelect) Statement() {}

func (p *ParenSelect) GetLimit() *Limit {
	return p.SelectStmt.GetLimit()
}

func (p *ParenSelect) GetGroupBy() *GroupBy {
	return p.SelectStmt.GetGroupBy()
}

func (p *ParenSelect) GetOrderBy() *OrderBy {
	return p.SelectStmt.GetOrderBy()
}

func (p *ParenSelect) GetSchemas() []string {
	return p.SelectStmt.GetSchemas()
}

func (*ParenSelect) Select() {}
