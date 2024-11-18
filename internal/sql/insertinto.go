package sqlparser

import "fmt"

type _InsertStmt struct {
	rows      uint
	tableName string
	columns   []string
	data      [][]string
}

func (t *_InsertStmt) Execute() *_Result {
	fmt.Println(t)
	// Execute Insert Statement
	return &_Result{
		Output: "nada",
	}
}

func parseInsertInto(p *_Parser) _SQLQuery {
	var stmt _InsertStmt
	var dataRow []string
	rows := uint(1)

	p.nextToken() // Table Name
	p.expect(TK_IDENTIFIER, "Table Name")
	stmt.tableName = p.token.Value
	p.nextToken() // Columns (column1, column2, ...)
	stmt.columns = p.parseColumnNames()
	p.expect(TK_KW_VALUES, "VALUES")
	p.nextToken() // lparen
	for p.token.Type != TK_SEMICOLON {
		dataRow = p.parseColumnDatas()
		stmt.data = append(stmt.data, dataRow)

		if p.token.Type == TK_COMMA {
			p.expect(TK_COMMA, ",")
			p.nextToken()
			rows++
		}
	}
	stmt.rows = rows
	p.expect(TK_SEMICOLON, ";")

	return &stmt
}
