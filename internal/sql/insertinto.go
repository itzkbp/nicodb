package sqlparser

import "fmt"

type _InsertStmt struct {
	tableName string
	columns   []string
	data      []string
}

func (t *_InsertStmt) Execute() *_Result {
	fmt.Println(t)
	// Execute Insert Statement
	return &_Result{
		output: "nada",
	}
}

func parseInsertInto(p *_Parser) _SQLQuery {
	var stmt _InsertStmt

	p.nextToken() // Table Name

	return &stmt
}
