package parser

import (
	"fmt"
	"lang/ast"
	"lang/token"
)

// Parser For each Statement
type Parser struct {
	current int
	tokens  *token.Statement
}

func ParseProgram(stmts *[]token.Statement) ast.ASTNode {
	n := &ast.ProgramNode{}
	for _, stmt := range *stmts {
		parse := &Parser{
			current: 0,
			tokens:  &stmt,
		}
		AppendNode(&n.Statements, parse.ParseStatement())
	}
	return n
}

func (p *Parser) ParseStatement() ast.ASTNode {
	var n ast.ASTNode
	if p.isDecl() {
		n = p.ParseDecl()
	}
	fmt.Printf("")
	return n
}

func (p *Parser) ParseDecl() ast.ASTNode {
	n := &ast.DeclNode{}
	n.Kind = p.consume().Type
	n.Ident = p.consume().Lexeme
	p.consume()
	lit := &ast.LitNode{
		Value: p.peek().Value,
		Kind:  p.peek().Type,
		Raw:   p.peek().Lexeme,
	}
	p.consume()
	n.Value = lit
	return n
}

func (p *Parser) isDecl() bool {
	val := isVarDecl(p.peek().Type)
	if p.peek(p.current+1).Type == token.IDENT {
		val = true
	} else {
		val = false
	}
	return val
}
func isVarDecl(t token.TokenType) bool {
	return t == token.VAR
}

// UTILS

func (p *Parser) peek(t ...int) token.Token {
	if len(t) > 0 {
		if t[0] >= len(*p.tokens) {
			return token.Token{}
		}
		return (*p.tokens)[t[0]]

	} else {
		if p.current >= len(*p.tokens) {
			return token.Token{}
		}
		return (*p.tokens)[p.current]
	}

}

func (p *Parser) consume() token.Token {
	n := (*p.tokens)[p.current]
	p.current++
	return n
}

func AppendNode(nodes *[]ast.ASTNode, node ast.ASTNode) {
	*nodes = append(*nodes, node)
}
