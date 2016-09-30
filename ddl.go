package sql

// AlterTableStmt  change  structure of a table.
// See https://dev.mysql.com/doc/refman/5.7/en/alter-table.html
type AlterTableStmt struct {
	Table Table
}

//IsDDL  true
func (a *AlterTableStmt) IsDDL() bool {
	return true
}

//IsDML false
func (a *AlterTableStmt) IsDML() bool {
	return false
}
