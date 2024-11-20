package sqlparser

import "fmt"

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
