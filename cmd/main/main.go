package main

import (
	"fmt"
	"os"
	"github.com/mattellis91/glox/pkg/lexing"
	"github.com/mattellis91/glox/pkg/parsing"
	"github.com/mattellis91/glox/pkg/util"
)

func main() {

	argsLen := len(os.Args)

	printer := &parsing.AstPrinter{}
	printer.Print(parsing.NewLiteralExpression(123))

	//TESTING AST PRINTER
	ue := parsing.NewUnaryExpression(lexing.Token{TokenType: lexing.Minus, Lexeme: "-", Literal: nil, Line: 1}, parsing.NewLiteralExpression(1123));
	be := parsing.NewBinaryExpression(ue, lexing.Token{TokenType: lexing.Star, Lexeme: "*", Literal: nil, Line: 1}, parsing.NewGroupingExpression(parsing.NewLiteralExpression(45.67)))
	printer.Print(ue) 
	printer.Print(be)

	if argsLen > 2 {
		fmt.Println("Ussage: zima [script]")
	} else if argsLen == 2{
		runFile((os.Args[1]))
	} else {
		runPrompt()
	}
}

func runFile(filepath string) {
	dat, err := os.ReadFile(filepath)
	util.Check(err)
	run(string(dat))
}

func run(source string) {
	lexer := lexing.NewLexer(source)
	tokens := lexer.Tokenize()

	for _, token := range tokens {
		fmt.Println(token.ToString())
	} 
}

func runPrompt() {
	for {
		fmt.Print(">: ")
		var input string
		fmt.Scanln(&input)
		if (len(input) == 0 || input == "q" ) {
			break
		}
		fmt.Println(input)
	}
}

