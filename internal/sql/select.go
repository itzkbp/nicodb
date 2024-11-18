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

type _SelectStmt struct {
	tableName string
	columns   []string
	condition *_Condition
}

func (t *_SelectStmt) Execute() *_Result {
	fmt.Println(t)
	if t.condition != nil {
		fmt.Println(t.condition.left)
		fmt.Println(t.condition.operator)
		fmt.Println(t.condition.right)
	}
	// Execute Select Statement
	return &_Result{
		Output: "nada",
	}
}

func (p *_Parser) parseField() *_Field {
	var field _Field
	field.name = p.token.Value

	if p.token.Type == TK_VAL_LITERAL {
		field.isLiteral = true
	}

	return &field
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

func parseSelect(p *_Parser) _SQLQuery {
	var stmt _SelectStmt

	p.nextToken()
	stmt.columns = p.parseColumnNames()
	if p.token.Type == TK_ASTERIK {
		p.expect(TK_ASTERIK, "*")
		p.nextToken() // from
	}

	p.expect(TK_KW_FROM, "FROM")
	p.nextToken() // Table Name
	p.expect(TK_IDENTIFIER, "Table Name")
	stmt.tableName = p.token.Value
	p.nextToken() // where or semicolon
	if p.token.Type == TK_KW_WHERE {
		p.expect(TK_KW_WHERE, "WHERE")
		p.nextToken() // condition
		stmt.condition = p.parseCondition()
	}

	p.expect(TK_SEMICOLON, ";")

	return &stmt
}
