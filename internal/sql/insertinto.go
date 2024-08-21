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
		Output: "nada",
	}
}

func (p *_Parser) parseColumnNames() []string {
	var columns []string
	// var column string

	if p.token.Type == TK_LPAREN {

	} else {
		// fill columns array with the columns from table definition
	}

	return columns
}

func parseInsertInto(p *_Parser) _SQLQuery {
	var stmt _InsertStmt

	p.nextToken() // Table Name
	p.expect(TK_IDENTIFIER, "Table Name")
	stmt.tableName = p.token.Value
	p.nextToken() // Columns (column1, column2, ...)
	stmt.columns = p.parseColumnNames()
	p.expect(TK_KW_VALUES, "VALUES")

	return &stmt
}
