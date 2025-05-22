package token

type Position struct {
	Line   int
	Column int
}

func (p *Position) isValid() bool { return p.Line > 0 }

type Pos int // current index
const NoPos int = 0
