package sql

//Select select_stmt
type Select struct {
	QueryExpression QueryExpression
}

//QueryExpression ...
type QueryExpression struct {
	QueryExpressionBody QueryExpressionBody
	Order               Order
	Limit               Limit
	Lock                Lock
}

//QueryExpressionBody ...
type QueryExpressionBody struct {
}

//Order ...
type Order struct {
}

//Limit ...
type Limit struct {
}

//Lock ...
type Lock struct {
}

//Table ...
type Table interface {
}

// TableName  table name
type TableName struct {
	Schema string
	Name   string
}
