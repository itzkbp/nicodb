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

func (l *_Lexer) nextToken() _Token {
	l.skipWhitespaces()

	var tok _Token

	switch l.cur_Char {
	case '*':
		tok = newToken(TK_ASTERIK, "*")
	case '=':
		tok = newToken(TK_EQUALS, "=")
	case ';':
		tok = newToken(TK_SEMICOLON, ";")
	case '(':
		tok = newToken(TK_LPAREN, "(")
	case ')':
		tok = newToken(TK_RPAREN, ")")
	case ',':
		tok = newToken(TK_COMMA, ",")

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
			default:
				tok = newToken(TK_IDENTIFIER, ident)
			}
		}
	}

	return tok
}
