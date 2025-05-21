package token

type Token int

const (
	IDENT = iota
	keyword_beg
	VAR
	PRINT
	keyword_end
	stmt_beg
	ASS
	stmt_end
)

var tokens = [...]string{
	VAR:   "var",
	PRINT: "print",
	ASS:   "=",
}

func (t *Token) isKeyword() bool { return *t < keyword_end && *t > keyword_beg }

func (t *Token) isStatment() bool { return *t < stmt_end && *t > stmt_beg }
