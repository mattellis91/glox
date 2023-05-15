package main

import (
	"fmt"
	"os"
	"github.com/mattellis91/zima/pkg/lexing"
	"github.com/mattellis91/zima/pkg/parsing"
	"github.com/mattellis91/zima/pkg/util"
)

func main() {

	argsLen := len(os.Args)

	printer := &parsing.AstPrinter{}
	printer.Print(parsing.NewLiteralExpression(123))

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

