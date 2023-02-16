package main

import (
	"fmt"
	"os"
	"github.com/mattellis91/zima/pkg/lexing"
)

func main() {

	argsLen := len(os.Args)

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
	check(err)
	run(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	} 
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

