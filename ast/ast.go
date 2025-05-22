package ast

import (
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

const __expr = 0
const (
	NODE = __expr
	LIT  = 1
	EXPR = 2
)

type Node struct {
	Level  int
	Kind   token.TokenType
	Lexeme string
	Value  interface{}
	Type   func() int
}

type LitNode struct {
	Node
}

type ExprNode struct {
	Node
	LNode *LitNode
	RNode *LitNode
}

func (n *Node) isExpr() bool { return n.Type() == EXPR }
func (n *Node) isLit() bool  { return n.Type() == LIT }
func (n *Node) isNode() bool { return n.Type() == NODE }

func NewNode(t *token.Token) Node {
	e := Node{}
	e.Type = func() int { return NODE }
	if t != nil {
		e.Value = t.Value
		e.Lexeme = t.Lexeme
		e.Kind = t.Type
	}
	return e
}

func Lit(t *token.Token) LitNode {
	e := LitNode{}
	e.Type = func() int { return LIT }
	if t != nil {
		e.Value = t.Value
		e.Lexeme = t.Lexeme
		e.Kind = t.Type
	}
	return e
}

func Expr(t *token.Token) ExprNode {
	e := ExprNode{}
	e.Type = func() int { return EXPR }
	if t != nil {
		e.Value = t.Value
		e.Lexeme = t.Lexeme
		e.Kind = t.Type
	}
	return e
}
