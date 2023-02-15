package lexing

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme string
	literal any
	line int
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

func (t Token) ToString() string{
	return fmt.Sprintf("%s %s %v", TokenToString(t.tokenType), t.lexeme, t.literal)
}