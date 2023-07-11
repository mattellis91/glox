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