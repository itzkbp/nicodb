package sqlparser

type _TokenKind uint8

const (
	TK_INVALID _TokenKind = iota

	TK_IDENTIFIER

	TK_ASTERIK

	TK_SEMICOLON

	TK_LPAREN
	TK_RPAREN

	TK_COMMA

	TK_GT
	TK_LT
	TK_EXCLAMATION
	TK_EQUALS

	TK_DOT

	TK_VAL_LITERAL

	// Data Types
	TK_DT_NUMBER
	TK_DT_TEXT
	TK_DT_FLOAT
	TK_DT_BOOL
	TK_DT_ANY

	// Constraints
	TK_PRIMARY
	TK_KEY
	TK_NOT
	TK_NULL
	TK_CONSTRAINT_ANY

	// SQL Keywords
	TK_KW_CREATE
	TK_KW_TABLE
	TK_KW_SELECT
	TK_KW_FROM
	TK_KW_WHERE
	TK_KW_INSERT
	TK_KW_INTO
	TK_KW_VALUES
	TK_KW_UPDATE
	TK_KW_SET
	TK_KW_DELETE
)

type _Token struct {
	Type  _TokenKind
	Value string
}

func newToken(t _TokenKind, v string) *_Token {
	return &_Token{t, v}
}
