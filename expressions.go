package sql

// Subquery Syntax
type SubQuery struct {
}

type TableExpression struct {
	From    Table
	GroupBy *GroupBy
	OrderBy *OrderBy
	Limit   *Limit
	LockTp  SelectLockType
}

type Expr interface {
	Expr()
}

type Exprs []Expr

type BoolExpr interface {
	Expr()
	BoolExpr()
}

type ValueExpr interface {
	Expr()
	ValueExpr()
}
