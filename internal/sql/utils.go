package sqlparser

import "fmt"

type _Operator uint8

const (
	OP_EQUALS _Operator = iota
	OP_NOT_EQUALS
	OP_GT
	OP_LT
	OP_GT_EQUALS
	OP_LT_EQUALS
)

type _Field struct {
	name      string
	isLiteral bool
}

type _Condition struct {
	left     *_Field
	operator _Operator
	right    *_Field
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

func (p *_Parser) parseOperator() _Operator {
	var operator _Operator

	switch p.token.Type {
	case TK_EQUALS:
		p.expect(TK_EQUALS, "=")
		operator = OP_EQUALS
		p.nextToken()
	case TK_EXCLAMATION:
		p.expect(TK_EXCLAMATION, "!")
		p.nextToken()
		if p.token.Type == TK_EQUALS {
			p.expect(TK_EQUALS, "=")
			operator = OP_NOT_EQUALS
			p.nextToken()
		}
	case TK_LT:
		p.expect(TK_LT, "<")
		p.nextToken()
		if p.token.Type == TK_EQUALS {
			p.expect(TK_EQUALS, "=")
			operator = OP_LT_EQUALS
			p.nextToken()
		} else {
			operator = OP_LT
		}
	case TK_GT:
		p.expect(TK_GT, ">")
		p.nextToken()
		if p.token.Type == TK_EQUALS {
			p.expect(TK_EQUALS, "=")
			operator = OP_GT_EQUALS
			p.nextToken()
		} else {
			operator = OP_GT
		}

	}
	return operator
}

func (p *_Parser) parseCondition() *_Condition {
	var condition _Condition

	condition.left = p.parseField()
	p.nextToken() // operator
	condition.operator = p.parseOperator()
	condition.right = p.parseField()

	p.nextToken()
	p.expect(TK_SEMICOLON, ";")

	return &condition
}
