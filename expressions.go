package sql

type TableExpression struct {
	From    Table
	GroupBy *GroupBy
	OrderBy *OrderBy
	Limit   *Limit
	LockTp  SelectLockType
}

type Expr interface {
	Expression()
}

type Exprs []Expr

func (Exprs) Expression()      {}
func (Exprs) ValueExpression() {}

type BoolExpr interface {
	Expression()
	BoolExpression()
}

type ValueExpr interface {
	Expression()
	ValueExpression()
}

type OrExpr struct {
	Left, Right Expr
}

func (*OrExpr) Expression() {}

type XorExpr struct {
	Left, Right Expr
}

func (*XorExpr) Expression() {}

type AndExpr struct {
	Left, Right Expr
}

func (*AndExpr) Expression() {}

type NotExpr struct {
	Expr Expr
}

func (*NotExpr) Expression() {}

type TrueExpr struct {
	Expr Expr
	Not  bool
}

func (*TrueExpr) Expression() {}

type FalseExpr struct {
	Expr Expr
	Not  bool
}

func (*FalseExpr) Expression() {}

type UnknownExpr struct {
	Expr Expr
	Not  bool
}

func (*UnknownExpr) Expression() {}

type IsNullPri struct {
	Expr BoolExpr
	Not  bool
}

func (*IsNullPri) BoolExpression() {}
func (*IsNullPri) Expression()     {}

type CompOp struct {
	Left  BoolExpr
	Right ValueExpr
}

func (*CompOp) BoolExpression() {}
func (*CompOp) Expression()     {}

type CompOpAll struct {
	Left  BoolExpr
	Right ValueExpr
}

func (*CompOpAll) BoolExpression() {}
func (*CompOpAll) Expression()     {}

type Predicate struct {
	Expr ValueExpr
}

func (*Predicate) Expression()     {}
func (*Predicate) BoolExpression() {}

type InCondition struct {
	Expr ValueExpr
	Not  bool
	List Exprs
}

func (*InCondition) Expression()      {}
func (*InCondition) ValueExpression() {}

type BetweenCondition struct {
	Not   bool
	Right ValueExpr
	Left  ValueExpr
}

func (*BetweenCondition) Expression()      {}
func (*BetweenCondition) ValueExpression() {}

type LikeCondition struct {
	Expr    ValueExpr
	Not     bool
	Pattern ValueExpr
}

func (*LikeCondition) Expression()      {}
func (*LikeCondition) ValueExpression() {}

type RgexpCondition struct {
	Expr    ValueExpr
	Not     bool
	Pattern ValueExpr
}

func (*RgexpCondition) Expression()      {}
func (*RgexpCondition) ValueExpression() {}

const (
	BITAND     = "&"
	BITOR      = "|"
	BITXOR     = "^"
	PLUS       = "+"
	MINUS      = "-"
	MULT       = "*"
	DIV        = "/"
	MOD        = "%"
	SHIFTLEFT  = "<<"
	SHIFTRIGHT = ">>"
	OPEQ       = "="
	OPEQUAL    = "=="
	OPLT       = "<"
	OPGT       = ">"
	OPLE       = "<="
	OPGE       = ">="
	OPNE       = "!="
	OPNSE      = "<=>"
	OROR       = "||"
	UPLUS      = "+"
	UMINUS     = "-"
	TILDA      = "~"
	NOT2       = "!"
	UBINARY    = "binary"
)

type BinaryOperationExpr struct {
	Operator    string
	Left, Right ValueExpr
}

func (*BinaryOperationExpr) Expression()      {}
func (*BinaryOperationExpr) ValueExpression() {}

type IntervalExpr struct {
	Expr     Expr
	Interval []byte
}

func (*IntervalExpr) Expression()      {}
func (*IntervalExpr) ValueExpression() {}

type FuncExpr struct{}

func (*FuncExpr) Expression()      {}
func (*FuncExpr) ValueExpression() {}

type CollateExpr struct {
	Expr    ValueExpr
	Collate []byte
}

func (*CollateExpr) Expression()      {}
func (*CollateExpr) ValueExpression() {}

type ExistsExpr struct {
	SubQuery *SubQueryStmt
}

func (*ExistsExpr) Expression()      {}
func (*ExistsExpr) ValueExpression() {}

type IdentExpr struct {
	Ident []byte
	Expr  Expr
}

func (*IdentExpr) Expression()      {}
func (*IdentExpr) ValueExpression() {}

type MatchExpr struct{}

func (*MatchExpr) Expression()      {}
func (*MatchExpr) ValueExpression() {}
