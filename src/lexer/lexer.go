package lexer

import "github.com/ythosa/pukiclang/src/token"

// Lexer is type for lexer which turns code into a sequence of tokens
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	// todo: "ch" could be of <rune> type in the feature for the all unicode support
}

// New returns new lexer
func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()

	return &l
}

// NextToken returns next token of the code
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = makeTwoCharComparisonToken(l.ch)
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = makeTwoCharComparisonToken(l.ch)
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = makeTwoCharComparisonToken(l.ch)
			l.readChar()
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			tok = makeTwoCharComparisonToken(l.ch)
			l.readChar()
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func makeTwoCharComparisonToken(current byte) token.Token {
	if current == '=' {
		return token.Token{Type: token.EQ, Literal: "=="}
	}
	if current == '!' {
		return token.Token{Type: token.NOTEQ, Literal: "!="}
	}
	if current == '>' {
		return token.Token{Type: token.GTEQ, Literal: ">="}
	}
	if current == '<' {
		return token.Token{Type: token.LTEQ, Literal: "<="}
	}

	return token.Token{Type: token.ILLEGAL, Literal: ""}
}
