package main

import (
	"fmt"
	"lang/io"
	"lang/lexer"
	"lang/parser"
	"os"
)

func main() {
	args := os.Args
	file := args[1]
	f := print.ReadFile(file)
	tokens := lexer.Lexer(f)
	fmt.Println(parser.ParseProgram(tokens).String())
}
