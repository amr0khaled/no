package main

import (
	"fmt"
	"lang/ast"
	"lang/io"
	"os"
)

func main() {
	args := os.Args
	file := args[1]
	f := print.ReadFile(file)
	for _, line := range f.Lines {
		fmt.Println(ast.Lexer(line.Text))
	}
}
