package sqlparser

import "fmt"

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
		fmt.Println("fill columns array with the columns from table definition")
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
