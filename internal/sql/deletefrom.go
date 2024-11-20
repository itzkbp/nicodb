package sqlparser

import "fmt"

type _DeleteStmt struct {
	tableName string
	condition *_Condition
}

func (t *_DeleteStmt) Execute() *_Result {
	fmt.Println(t)
	if t.condition != nil {
		fmt.Println(t.condition.left)
		fmt.Println(t.condition.operator)
		fmt.Println(t.condition.right)
	}

	return &_Result{
		Output: "nada",
	}
}

func parseDeleteFrom(p *_Parser) _SQLQuery {
	var stmt _DeleteStmt

	p.nextToken() // tableName
	stmt.tableName = p.token.Value

	p.nextToken() // WHERE
	p.expect(TK_KW_WHERE, "WHERE")

	p.nextToken()
	stmt.condition = p.parseCondition()

	return &stmt
}
