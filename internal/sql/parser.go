package sqlparser

import (
	"log"
	"strings"
)

type _Result struct {
	output string
}

type _Parser struct {
	lexer *_Lexer
	token *_Token
}

type _SQLQuery interface {
	Execute() *_Result
}

type Operator uint8

const (
	OP_EQUALS Operator = iota
	OP_NOT_EQUALS
	OP_GT
	OP_LT
	OP_GT_EQUALS
	OP_LT_EQUALS
)

type Condition struct {
	option1    string
	operator   Operator
	option2    string
	isOp2Field bool
}

type SelectStmt struct {
	tableName string
	columns   []string
	condition Condition
}

func (t *SelectStmt) Execute() _Result {
	// Execute Select Statement
	return _Result{
		output: "nada",
	}
}

func NewParser(l *_Lexer) *_Parser {
	return &_Parser{
		lexer: l,
	}
}

func (p *_Parser) nextToken() {
	p.token = p.lexer.nextToken()
}

func (p *_Parser) expect(tkKind _TokenKind, tkValue string) {
	if p.token.Type == tkKind && strings.ToUpper(p.token.Value) == tkValue {
		return
	}
	if p.token.Type == TK_IDENTIFIER && p.token.Type == tkKind {
		return
	}

	log.Fatalf("(Parser): Expected %s got %s\n", tkValue, p.token.Value)
}

func (p *_Parser) Parse() _SQLQuery {
	p.nextToken()

	var query _SQLQuery

	switch p.token.Type {
	case TK_KW_CREATE:
		p.expect(TK_KW_CREATE, "CREATE")
		p.nextToken()

		if p.token.Type == TK_KW_TABLE {
			p.expect(TK_KW_TABLE, "TABLE")

			query = parseCreateTable(p)
		}

	case TK_KW_INSERT:
		p.expect(TK_KW_INSERT, "INSERT")
		p.nextToken()

		if p.token.Type == TK_KW_INTO {
			p.expect(TK_KW_INTO, "INTO")

			query = parseInsertInto(p)
		}
	}

	return query
}
