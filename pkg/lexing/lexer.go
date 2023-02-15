package lexing

import (
	"github.com/mattellis91/zima/pkg/reporting"
)

type Lexer struct {
	source string
	tokens []Token
	start int
	current int
	line int
}

func (l *Lexer) init(source string) {
	l.start = 0
	l.current = 0
	l.line = 1
	l.source = source
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
	switch(c) {
		case '(' : l.addToken(LeftParen, nil)
		case ')' : l.addToken(RightParen, nil)
		case '{' : l.addToken(LeftBrace, nil)
		case '}' : l.addToken(RightBrace, nil)
		case ',' : l.addToken(Comma, nil)
		case '.' : l.addToken(Dot, nil)
		case '-' : l.addToken(Minus, nil)
		case '+' : l.addToken(Plus, nil) 
		case ';' : l.addToken(Semicolon, nil)
		case '*' : l.addToken(Star, nil)
		case '!' : {
			if l.match('=') {
				l.addToken(BangEqual, nil)
			} else {
				l.addToken(Bang, nil)
			}
		}
		case '=' : {
			if l.match('=') {
				l.addToken(EqualEqual, nil)
			} else {
				l.addToken(Equal, nil)
			}
		} 
		case '<' : {
			if l.match('=') {
				l.addToken(LessEqual, nil)
			} else {
				l.addToken(Less, nil)
			}
		} 
		case '>' : {
			if l.match('=') {
				l.addToken(GreaterEqual, nil)
			} else {
				l.addToken(Greater, nil)
			}
		} 
		default:
			reporting.ErrorMessage(l.line, "Unexpected Character")
			
	}
}

func (l *Lexer) advance() byte{
	l.current++
	return l.source[l.current]
}

func (l *Lexer) addToken(tt TokenType, literal any) {
	text := l.source[l.start : l.current]
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

