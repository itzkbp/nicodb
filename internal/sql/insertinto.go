package sqlparser

type _InsertStmt struct {
	tableName string
	columns   []string
	data      []string
}

func (t *_InsertStmt) Execute() *_Result {
	// Execute Insert Statement
	return &_Result{
		output: "nada",
	}
}

func parseInsertInto(p *_Parser) _SQLQuery {
	var stmt _InsertStmt

	return &stmt
}
