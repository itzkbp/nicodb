package sqlparser

import (
	"fmt"
	"log"
)

type _ColumnDefinition struct {
	columnName   string
	dataType     _TokenKind
	isPrimaryKey bool
	isNullable   bool
}

type _CreateTableStmt struct {
	tableName string
	columns   []_ColumnDefinition
}

func (t *_CreateTableStmt) Execute() *_Result {
	fmt.Println(t)
	// Execute Create Table
	return &_Result{
		Output: "nada",
	}
}

func (p *_Parser) parseColumnDefinitions() []_ColumnDefinition {
	var columns []_ColumnDefinition
	var column _ColumnDefinition
	hasPK := false

	// continue from ( to ) & parse each column spearated by ,
	p.expect(TK_LPAREN, "(")
	p.nextToken() // column definitions ( columnname datatype constraints, ...)

	for p.token.Type != TK_RPAREN {

		p.expect(TK_IDENTIFIER, "Column Name")
		column.columnName = p.token.Value
		column.isNullable = true
		column.isPrimaryKey = false

		p.nextToken() // data type
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
		for !(p.token.Type == TK_COMMA || p.token.Type == TK_RPAREN) { // possibility of multiple constraints
			// parsing additional constraints
			switch p.token.Type {
			case TK_NOT:
				p.expect(TK_NOT, "NOT")
				p.nextToken() // null

				if p.token.Type == TK_NULL {
					p.expect(TK_NULL, "NULL")
					column.isNullable = false
					p.nextToken() // comma or rparen
				}
			case TK_PRIMARY:
				p.expect(TK_PRIMARY, "PRIMARY")
				p.nextToken() // key

				if p.token.Type == TK_KEY {
					p.expect(TK_KEY, "KEY")
					column.isPrimaryKey = true

					if hasPK {
						log.Fatal("A Table cannot have multiple PRIMARY KEY.")
					}

					hasPK = true
					p.nextToken() // comma or rparen
				}
			default:
				p.expect(TK_CONSTRAINT_ANY, "Constraint: {PRIMARY KEY | NUT NULL}")
			}
		}

		if p.token.Type == TK_COMMA {
			p.expect(TK_COMMA, ",")
			p.nextToken() // next column
		}

		columns = append(columns, column)
	}

	if !hasPK {
		log.Fatal("A Table must have a PRIMARY KEY.")
	}
	return columns
}

func parseCreateTable(p *_Parser) _SQLQuery {
	var stmt _CreateTableStmt

	p.nextToken() // Table Name
	p.expect(TK_IDENTIFIER, "Table Name")
	stmt.tableName = p.token.Value
	p.nextToken() // Column Definitions
	stmt.columns = p.parseColumnDefinitions()
	p.nextToken() // Semicolon
	p.expect(TK_SEMICOLON, ";")

	return &stmt
}
