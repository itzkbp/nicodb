package sqlparser

import (
	"strings"
	"unicode"
)

type _Lexer struct {
	query    string
	position int
	cur_Char rune
}

func NewLexer(query string) *_Lexer {
	l := &_Lexer{query: query}
	l.readChar()
	return l
}

func (l *_Lexer) readChar() {
	if l.position >= len(l.query) {
		l.cur_Char = ';'
	} else {
		l.cur_Char = rune(l.query[l.position])
	}

	l.position++
}

func (l *_Lexer) skipWhitespaces() {
	for unicode.IsSpace(l.cur_Char) {
		l.readChar()
	}
}

func (l *_Lexer) readIdentifier() string {
	start := l.position - 1
	for unicode.IsLetter(l.cur_Char) || unicode.IsDigit(l.cur_Char) || l.cur_Char == '_' {
		l.readChar()
	}
	return l.query[start : l.position-1]
}

func (l *_Lexer) readLiteral() string {
	start := l.position - 1
	if l.cur_Char == '\'' || l.cur_Char == '"' || l.cur_Char == '`' {
		l.readChar()
		for !(l.cur_Char == '\'' || l.cur_Char == '"' || l.cur_Char == '`') {
			l.readChar()
		}
		l.readChar()
	} else {
		for unicode.IsDigit(l.cur_Char) || l.cur_Char == '.' {
			l.readChar()
		}
	}

	return l.query[start : l.position-1]
}

func (l *_Lexer) nextToken() *_Token {
	l.skipWhitespaces()

	var tok *_Token

	switch l.cur_Char {
	case '*':
		tok = newToken(TK_ASTERIK, "*")
		l.readChar()
	case '=':
		tok = newToken(TK_EQUALS, "=")
		l.readChar()
	case ';':
		tok = newToken(TK_SEMICOLON, ";")
		l.readChar()
	case '(':
		tok = newToken(TK_LPAREN, "(")
		l.readChar()
	case ')':
		tok = newToken(TK_RPAREN, ")")
		l.readChar()
	case ',':
		tok = newToken(TK_COMMA, ",")
		l.readChar()
	case '.':
		tok = newToken(TK_DOT, ".")
		l.readChar()

	default:
		if unicode.IsLetter(l.cur_Char) || l.cur_Char == '_' {
			ident := l.readIdentifier()

			switch strings.ToUpper(ident) {
			case "CREATE":
				tok = newToken(TK_KW_CREATE, ident)
			case "TABLE":
				tok = newToken(TK_KW_TABLE, ident)
			case "SELECT":
				tok = newToken(TK_KW_SELECT, ident)
			case "FROM":
				tok = newToken(TK_KW_FROM, ident)
			case "WHERE":
				tok = newToken(TK_KW_WHERE, ident)
			case "INSERT":
				tok = newToken(TK_KW_INSERT, ident)
			case "INTO":
				tok = newToken(TK_KW_INTO, ident)
			case "VALUES":
				tok = newToken(TK_KW_VALUES, ident)

			case "NUMBER":
				tok = newToken(TK_DT_NUMBER, ident)
			case "TEXT":
				tok = newToken(TK_DT_TEXT, ident)
			case "FLOAT":
				tok = newToken(TK_DT_FLOAT, ident)
			case "BOOL":
				tok = newToken(TK_DT_BOOL, ident)

			case "PRIMARY":
				tok = newToken(TK_PRIMARY, ident)
			case "KEY":
				tok = newToken(TK_KEY, ident)
			case "NOT":
				tok = newToken(TK_NOT, ident)
			case "NULL":
				tok = newToken(TK_NULL, ident)

			default:
				tok = newToken(TK_IDENTIFIER, ident)
			} // switch identifier
		} else if unicode.IsDigit(l.cur_Char) || l.cur_Char == '\'' || l.cur_Char == '"' || l.cur_Char == '`' {
			tok = newToken(TK_VAL_LITERAL, l.readLiteral())
		} else {
			tok = newToken(TK_INVALID, string(l.cur_Char))
		}
	}

	// fmt.Println("Tokenization: ", tok.Type, tok.Value)
	return tok
}
