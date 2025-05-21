package token

type Position struct {
	Filename string
	Line     int
	Column   int
}

func (p *Position) isValid() bool { return p.Line > 0 }

type Pos int // current index
const NoPos int = 0

type Node interface {
	Pos() Pos
	End() Pos
}

type Expr interface {
	Node
	exprNode()
}

type Lit struct {
	Pos   Pos
	Type  Token
	Value string
}

type AssExpr struct {
	Ass Pos
	X   Expr
}
