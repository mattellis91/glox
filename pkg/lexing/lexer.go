package lexing

import (
	"strconv"
	"github.com/mattellis91/glox/pkg/util"
)

type Lexer struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewLexer(source string) *Lexer {

	return &Lexer{
		start:   0,
		current: 0,
		line:    1,
		source:  source,
	}
}

func (l *Lexer) Tokenize() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	l.tokens = append(l.tokens, Token{Eof, "", nil, l.line})
	return l.tokens
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) scanToken() {
	c := l.advance()
	switch c {
	case '(':
		l.addToken(LeftParen, nil)
	case ')':
		l.addToken(RightParen, nil)
	case '{':
		l.addToken(LeftBrace, nil)
	case '}':
		l.addToken(RightBrace, nil)
	case ',':
		l.addToken(Comma, nil)
	case '.':
		l.addToken(Dot, nil)
	case '-':
		l.addToken(Minus, nil)
	case '+':
		l.addToken(Plus, nil)
	case ';':
		l.addToken(Semicolon, nil)
	case '*':
		l.addToken(Star, nil)
	case '!':
		{
			if l.match('=') {
				l.addToken(BangEqual, nil)
			} else {
				l.addToken(Bang, nil)
			}
		}
	case '=':
		{
			if l.match('=') {
				l.addToken(EqualEqual, nil)
			} else {
				l.addToken(Equal, nil)
			}
		}
	case '<':
		{
			if l.match('=') {
				l.addToken(LessEqual, nil)
			} else {
				l.addToken(Less, nil)
			}
		}
	case '>':
		{
			if l.match('=') {
				l.addToken(GreaterEqual, nil)
			} else {
				l.addToken(Greater, nil)
			}
		}
	case '/':
		{
			if l.match('/') {
				//comment tokens go until end of line
				for l.peek() != '\n' && !l.isAtEnd() {
					l.advance()
				}
			} else {
				l.addToken(Slash, nil)
			}
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		l.line++
	case '"':
		l.string()
	default:
		if(isDigit(c)) {
			l.number()
		} else if isAlpha(c) {
			l.identifier()
		} else {
			util.ErrorMessage(l.line, "Unexpected Character")
		}
	}
}

func (l *Lexer) advance() byte {
	l.current++
	return l.source[l.current-1]
}

func (l *Lexer) addToken(tt TokenType, literal any) {
	text := l.source[l.start:l.current]
	l.tokens = append(l.tokens, Token{tt, text, literal, l.line})
}

func (l *Lexer) match(expected byte) bool {
	if l.isAtEnd() {
		return false
	}
	if l.source[l.current] != expected {
		return false
	}
	l.current++
	return true
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}
	return l.source[l.current]
}

func (l *Lexer) peekNext() byte {
	if l.current + 1 >= len(l.source) {
		return 0
	}
	return l.source[l.current + 1]
}

func (l *Lexer) string() {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	if l.isAtEnd() {
		util.ErrorMessage(l.line, "Unterminated String")
		return
	}

	l.advance()
	value := l.source[l.start+1 : l.current-1]
	l.addToken(String, value)
}

func (l *Lexer) number() {
	for isDigit(l.peek()) {
		l.advance()
	}

	if(l.peek() == '.' && isDigit(l.peekNext())) {
		l.advance()
		for isDigit(l.peek()) {
			l.advance()
		}
	}

	val, err := strconv.ParseFloat(l.source[l.start : l.current], 64)
	util.Check(err)

	l.addToken(Number, val)
}

func (l *Lexer) identifier() {
	for isAlphaNumeric(l.peek()) {
		l.advance()
	}

	text := l.source[l.start : l.current]
	tt, exists := keywords[text]
	if !exists {
		tt = Identifier
	}
	l.addToken(tt, nil)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		   (c >= 'A' && c <= 'Z') ||
		    c == '_';
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
