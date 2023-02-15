package main

import (
	"fmt"
	"os"
)

func main() {

	args_len := len(os.Args)

	if args_len > 2 {
		fmt.Println("Ussage: zima [script]")
	} else if args_len == 2{
		runFile((os.Args[1]))
	} else {
		runPrompt()
	}
}

func runFile(filepath string) {
	dat, err := os.ReadFile(filepath)
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
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