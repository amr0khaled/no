package ast

import (
	"fmt"
	"lang/token"
)

// Basic Node Types
// type Node interface {
// 	Level() int
// 	Node() *Node
// }
//
// type Decl interface {
// 	Node
// 	declNode()
// }
//
// type Stmt interface {
// 	Node
// 	stmtNode()
// }
// type Expr interface {
// 	Node
// 	exprNode()
// }
//
// // Declerations Structures
// type Spec interface {
// 	Node
// 	specNode()
// }
//
// type TypeSpec struct {
// 	Spec
// 	Token token.TokenType
// 	Value string // name of variable
// }
// type LitSpec struct {
// 	Spec
// 	Token  token.TokenType
// 	Lexeme string
// 	Value  interface{}
// }
//
// // the whole decleration
// // Type -> var
// // Value -> string("Helo")
// type DeclStmt struct {
// 	Stmt Stmt
// 	Type *TypeSpec
// 	Lit  *LitSpec
// }
//
// // validation
// func (*DeclStmt) stmtNode() {}
// func (*TypeSpec) specNode() {}
// func (*LitSpec) specNode()  {}

type NodeType int

const (
	NODE NodeType = iota
	LIT
	EXPR
	DECL
	PROG
)

type ASTNode interface {
	Type() NodeType
	String() string
}

type LitNode struct {
	Kind  token.TokenType
	Raw   string
	Value interface{}
}

func (*LitNode) Type() NodeType { return LIT }
func (n *LitNode) String() string {
	return fmt.Sprintf(
		`Literal: {
  Token: %v
  NodeType: %v
  Value: %v
}`, token.Tokenize(n.Kind), n.Type(), n.Value)
}

type ExprNode struct {
	Kind  token.TokenType
	LNode ASTNode
	RNode ASTNode
}

func (*ExprNode) Type() NodeType { return EXPR }
func (n *ExprNode) String() string {
	return fmt.Sprintf(
		`Expression: {
  Token: %v
  NodeType: %v
  LNode: %v
  RNode: %v
}`,
		token.Tokenize(n.Kind), n.Type(), n.LNode.String(), n.RNode.String())
}

type DeclNode struct {
	Level int
	Kind  token.TokenType
	Ident string
	Value ASTNode
}

func (*DeclNode) Type() NodeType { return DECL }
func (n *DeclNode) String() string {
	return fmt.Sprintf(
		`Declaration: {
  Token: %v
  NodeType: %v
  Level: %v
  Ident: %v
  Value: %v
}`, token.Tokenize(n.Kind), n.Type(), n.Level, n.Ident, "\t"+n.Value.String())
}

type ProgramNode struct {
	Kind       token.TokenType
	Statements []ASTNode
}

func (*ProgramNode) Type() NodeType { return PROG }

func (n *ProgramNode) String() string {
	var _strs string = "[\n"
	for _, stat := range n.Statements {
		_strs += "\t"
		_strs += stat.String()
	}
	_strs += "\n]\n"
	return fmt.Sprintf(
		`Program: {
  Token: %v
  Statements: %v
}`, token.Tokenize(n.Kind), _strs)
}
