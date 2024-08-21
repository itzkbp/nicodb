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

type SQLQuery interface {
	Execute() _Result
}

type ColumnDefinition struct {
	columnName   string
	dataType     _TokenKind
	isPrimaryKey bool
	isNullable   bool
}

type CreateTableStmt struct {
	tableName string
	columns   []ColumnDefinition
}

func (t *CreateTableStmt) Execute() _Result {
	// Execute Create Table
	return _Result{
		output: "nada",
	}
}

type InsertStmt struct {
	tableName string
	columns   []string
	data      []string
}

func (t *InsertStmt) Execute() {
	// Execute Insert Statement
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

func (p *_Parser) parseColumnDefinitions() []ColumnDefinition {
	var columns []ColumnDefinition
	var column ColumnDefinition
	hasPK := false

	// continue from ( to ) & parse each column spearated by ,
	p.expect(TK_LPAREN, "(")

	for p.token.Type != TK_RPAREN {

		p.nextToken()
		p.expect(TK_IDENTIFIER, "Column Name")
		column.columnName = p.token.Value
		column.isNullable = true
		column.isPrimaryKey = false

		p.nextToken()
		switch p.token.Type {
		case TK_DT_NUMBER:
			p.expect(TK_DT_NUMBER, "NUMBER")
			column.dataType = p.token.Type
		case TK_DT_TEXT:
			p.expect(TK_DT_TEXT, "TEXT")
			column.dataType = p.token.Type
		case TK_DT_FLOAT:
			p.expect(TK_DT_FLOAT, "FLOAT")
			column.dataType = p.token.Type
		case TK_DT_BOOL:
			p.expect(TK_DT_BOOL, "BOOL")
			column.dataType = p.token.Type
		default:
			p.expect(TK_DT_ANY, "Data Type: {NUMBER | TEXT | FLOAT | BOOL}")
		}

		p.nextToken()
		if !(p.token.Type == TK_COMMA || p.token.Type == TK_RPAREN) {
			// parsing additional constraints
			switch p.token.Type {
			case TK_NOT:
				p.expect(TK_NOT, "NOT")
				p.nextToken()

				if p.token.Type == TK_NULL {
					p.expect(TK_NULL, "NULL")
					column.isNullable = false
					p.nextToken()
				}
			case TK_PRIMARY:
				p.expect(TK_PRIMARY, "PRIMARY")
				p.nextToken()

				if p.token.Type == TK_KEY {
					p.expect(TK_KEY, "KEY")
					column.isPrimaryKey = true

					if hasPK {
						log.Fatal("A Table cannot have multiple PRIMARY KEY.")
					}

					hasPK = true
					p.nextToken()
				}
			default:
				p.expect(TK_CONSTRAINT_ANY, "Constraint: {PRIMARY KEY | NUT NULL}")
			}
		}

		if p.token.Type == TK_COMMA {
			p.expect(TK_COMMA, ",")
			p.nextToken()
		}

		columns = append(columns, column)
	}

	if !hasPK {
		log.Fatal("A Table must have a PRIMARY KEY.")
	}
	return columns
}

func parseCreateTable(p *_Parser) SQLQuery {
	var stmt CreateTableStmt

	p.nextToken()
	p.expect(TK_IDENTIFIER, "Table Name")
	stmt.tableName = p.token.Value
	p.nextToken()
	stmt.columns = p.parseColumnDefinitions()
	p.nextToken()
	p.expect(TK_SEMICOLON, ";")

	return &stmt
}

func (p *_Parser) Parse() SQLQuery {
	p.nextToken()

	var query SQLQuery

	switch p.token.Type {
	case TK_KW_CREATE:
		p.expect(TK_KW_CREATE, "CREATE")
		p.nextToken()

		if p.token.Type == TK_KW_TABLE {
			p.expect(TK_KW_TABLE, "TABLE")

			query = parseCreateTable(p)
		}
	}

	return query
}
