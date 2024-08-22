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
	var column string

	if p.token.Type == TK_LPAREN {
		p.expect(TK_LPAREN, "(")
		p.nextToken() // column name

		for p.token.Type != TK_RPAREN {
			p.expect(TK_IDENTIFIER, "Column Name")
			column = p.token.Value
			columns = append(columns, column)
			p.nextToken() // comma or rparen

			if p.token.Type == TK_COMMA {
				p.expect(TK_COMMA, ",")
				p.nextToken() // next column
			}
		}

		p.expect(TK_RPAREN, ")")
		p.nextToken()
	} else {
		// fill columns array with the columns from table definition
	}
	return columns
}

func (p *_Parser) parseColumnDatas() []string {
	var datas []string
	var data string

	p.expect(TK_LPAREN, "(")
	p.nextToken() // datas (data, ...)
	for p.token.Type != TK_RPAREN {
		p.expect(TK_VAL_LITERAL, "Literal Value")
		data = p.token.Value
		datas = append(datas, data)
		p.nextToken()

		if p.token.Type == TK_COMMA {
			p.expect(TK_COMMA, ",")
			p.nextToken() // next column
		}
	}

	p.expect(TK_RPAREN, ")")
	p.nextToken()
	return datas
}

func parseInsertInto(p *_Parser) _SQLQuery {
	var stmt _InsertStmt

	p.nextToken() // Table Name
	p.expect(TK_IDENTIFIER, "Table Name")
	stmt.tableName = p.token.Value
	p.nextToken() // Columns (column1, column2, ...)
	stmt.columns = p.parseColumnNames()
	p.expect(TK_KW_VALUES, "VALUES")
	p.nextToken() // lparen
	stmt.data = p.parseColumnDatas()
	p.expect(TK_SEMICOLON, ";")

	return &stmt
}
