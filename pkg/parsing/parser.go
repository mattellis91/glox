package parsing

import (
	"github.com/mattellis91/glox/pkg/lexing"
)

type Parser struct {
	tokens []lexing.Token
	current int
}

func NewParser(tokens []lexing.Token) *Parser {
	return &Parser {
		tokens: tokens,
		current: 0,
	}
}

func (p *Parser) expression() Expression {
	return p.equality()
}

func (p *Parser) equality() Expression {
	return NewLiteralExpression(1)
}

func (p *Parser) match(tokenTypes ...lexing.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if(p.check(tokenType)) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType lexing.TokenType) bool {
	if (p.isAtEnd()) {
		return false
	}
	return p.peek().TokenType == tokenType
}

func (p *Parser) advance() lexing.Token {
	if(!p.isAtEnd()) {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == lexing.Eof
}

func (p *Parser) peek() lexing.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() lexing.Token {
	return p.tokens[p.current - 1]
}

func (p *Parser) comparison() Expression {
	exp := p.term()

	for p.match(
		lexing.Greater, 
		lexing.GreaterEqual,
		lexing.Less,
		lexing.LessEqual,) {
			operator := p.previous()
			right := p.term()
			exp = NewBinaryExpression(exp, operator, right) 
	}

	return exp
}

func (p *Parser) term() Expression {
	exp := p.factor()

	for p.match(lexing.Minus, lexing.Plus) {
		operator := p.previous()
		right := p.factor()
		exp = NewBinaryExpression(exp, operator, right)
	}

	return exp
}

func (p *Parser) factor() Expression {
	exp := p.unary()

	for p.match(lexing.Slash, lexing.Star) {
		operator := p.previous()
		right := p.unary()
		exp = NewBinaryExpression(exp, operator, right)
	}

	return exp
}

func (p *Parser) unary() Expression {
	if p.match(lexing.Bang, lexing.Minus) {
		operator := p.previous()
		right := p.unary()
		return NewUnaryExpression(operator, right)
	}

	return p.primary()
}

func (p *Parser) primary() Expression {
	if p.match(lexing.False){
		return NewLiteralExpression(false)
	}
	if p.match(lexing.True){
		return NewLiteralExpression(true)
	}
	if p.match(lexing.Nil){
		return NewLiteralExpression(nil)
	}
	if p.match(lexing.Number, lexing.String) {
		return NewLiteralExpression(p.previous().Literal)
	}
	if p.match(lexing.LeftParen) {
		exp := p.expression()
		p.consume(lexing.RightParen, "Expected ')' after expression")
		return NewGroupingExpression(exp)
	}
	return NewLiteralExpression(p.previous().Literal)
}

func (p *Parser) consume(tokenType lexing.TokenType, message string) lexing.Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(message) 
}




