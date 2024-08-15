package sqlparser

type _TokenKind string

const (
	TK_INVALID _TokenKind = "INVALID"

	TK_IDENTIFIER _TokenKind = "IDENTIFIER"

	TK_ASTERIK _TokenKind = "*"

	TK_SEMICOLON _TokenKind = ";"

	TK_LPAREN _TokenKind = "("
	TK_RPAREN _TokenKind = ")"

	TK_DOU_INV _TokenKind = "\""
	TK_SIN_INV _TokenKind = "'"
	TK_BK_TICK _TokenKind = "`"

	TK_COMMA _TokenKind = ","

	TK_EQUALS _TokenKind = "="

	TK_DOT _TokenKind = "."

	TK_VAL_NUMBER _TokenKind = "NUMERIC VALUE"
	TK_VAL_TEXT   _TokenKind = "STRING VALUE"
	TK_VAL_FLOAT  _TokenKind = "FLOAT VALUE"
	TK_VAL_BOOL   _TokenKind = "BOOLEAN VALUE"

	// Data Types
	TK_DT_NUMBER _TokenKind = "NUMBER"
	TK_DT_TEXT   _TokenKind = "TEXT"
	TK_DT_FLOAT  _TokenKind = "FLOAT"
	TK_DT_BOOL   _TokenKind = "BOOL"

	// Constraints
	TK_PRIMARY _TokenKind = "PRIMARY"
	TK_KEY     _TokenKind = "KEY"
	TK_NOT     _TokenKind = "NOT"
	TK_NULL    _TokenKind = "NULL"

	// SQL Keywords
	TK_KW_CREATE _TokenKind = "CREATE"
	TK_KW_TABLE  _TokenKind = "TABLE"
	TK_KW_SELECT _TokenKind = "SELECT"
	TK_KW_FROM   _TokenKind = "FROM"
	TK_KW_WHERE  _TokenKind = "WHERE"
)

type _Token struct {
	Type  _TokenKind
	Value string
}

func newToken(t _TokenKind, v string) _Token {
	return _Token{t, v}
}
