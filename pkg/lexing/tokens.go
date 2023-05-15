package lexing

import "fmt"

type Token struct {
	TokenType TokenType
	Lexeme string
	Literal any
	Line int
}

func (t *Token) ToString() string{
	return fmt.Sprintf("%s %s %v | line %d", TokenToString(t.TokenType), t.Lexeme, t.Literal, t.Line)
}

type TokenType int

const (
	// Single-character tokens.
	LeftParen TokenType = iota
	RightParen 
	LeftBrace  
	RightBrace
	Comma
	Dot
	Minus 
	Plus 
	Semicolon 
	Slash 
	Star
  
	// One or two character tokens.
	Bang 
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual
  
	// Literals.
	Identifier
	String 
	Number
  
	// Keywords.
	And
	Class 
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print 
	Return
	Super
	This
	True
	Var
	While
	Eof
)

var keywords = map[string]TokenType {
	"and": And,
	"class": Class,
	"else": Else,
	"false": False,
	"for": For,
	"fun": Fun,
	"if": If,
	"nil": Nil,
	"or": Or,
	"print": Print,
	"return": Return,
	"super": Super,
	"this": This,
	"true": True,
	"var": Var,
	"while": While,
}

func TokenToString(tt TokenType) string {
	return []string{
		"LeftParen",
		"RightParen",
		"LeftBrace",  
		"RightBrace",
		"Comma",
		"Dot",
		"Minus", 
		"Plus", 
		"Semicolon", 
		"Slash", 
		"Star",
		"Bang",
		"BangEqual",
		"Equal",
		"EqualEqual",
		"Greater",
		"GreaterEqual",
		"Less",
		"LessEqual",
		"Identifier",
		"String", 
		"Number",
		"And",
		"Class", 
		"Else",
		"False",
		"Fun",
		"For",
		"If",
		"Nil",
		"Or",
		"Print", 
		"Return",
		"Super",
		"This",
		"True",
		"Var",
		"While",
		"Eof",
	}[tt]
}