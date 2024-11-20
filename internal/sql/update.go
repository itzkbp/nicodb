package sqlparser

import "fmt"

type _DataKV map[string]string

type _UpdateStmt struct {
	tableName string
	datas     _DataKV
	condition *_Condition
}

func (t *_UpdateStmt) Execute() *_Result {
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

func (p *_Parser) parseUpdateKV() _DataKV {
	kvs := make(_DataKV)

	for p.token.Type != TK_KW_WHERE { // for multiple fields
		p.nextToken() // field (key)
		key := p.token.Value
		p.nextToken() // =
		p.nextToken() // data (value)
		kvs[key] = p.token.Value
		p.nextToken() // comma or WHERE keyword
	}

	return kvs
}

func parseUpdate(p *_Parser) _SQLQuery {
	var stmt _UpdateStmt

	p.nextToken()
	stmt.tableName = p.token.Value
	p.nextToken()
	p.expect(TK_KW_SET, "SET")
	stmt.datas = p.parseUpdateKV()
	p.expect(TK_KW_WHERE, "WHERE")
	p.nextToken()
	stmt.condition = p.parseCondition()

	return &stmt
}
