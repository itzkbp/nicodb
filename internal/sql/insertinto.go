package sqlparser

type InsertStmt struct {
	tableName string
	columns   []string
	data      []string
}

func (t *InsertStmt) Execute() {
	// Execute Insert Statement
}
