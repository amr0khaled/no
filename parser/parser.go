package parser

import (
	"lang/ast"
	"lang/token"
)

type Praser struct {
	level  int
	tokens *[][]token.Token
}

func parseProgram(tokens *[][]token.Token) {
	n := ast.NewNode(nil)
	n.Lexeme = "Program"
	n.Level = 0
	var nodes []ast.Node
	for _, toks := range *tokens {
		for i, tok := range toks {
			if tok.Type.IsLit() {
				node := ast.Lit(&tok)
				node.Level = i
			} else if tok.Type.IsDecl() {
				node := ast.Expr(&tok)
				node.Level = i
			}
		}
	}
}
func parseNode(t *token.Token,){
  switch e := t.Type {
  case VAR:


  }
}
