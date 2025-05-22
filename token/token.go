package token

import "strconv"

type TokenType int

type Statement []Token

type Token struct {
	Type   TokenType
	Lexeme string
	Value  interface{}
}

const (
	IDENT = iota
	keyword_beg
	decl_beg
	VAR
	decl_end
	PRINT
	keyword_end
	lit_beg
	INT //int
	lit_end
	stmt_beg
	ASS
	stmt_end
	token_end
)

var tokens = [...]string{
	VAR:   "var",
	PRINT: "print",
	ASS:   "=",
}

var keywords map[string]TokenType

func Tokenize(t TokenType) string {
	s := ""
	if t > IDENT && t < token_end {
		s = tokens[t]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}
func init() {
	keywords = make(map[string]TokenType, len(tokens))
	for i, t := range tokens {
		i := TokenType(i)
		keywords[t] = i
	}
}

func GetToken(word string) TokenType {
	return keywords[word]
}

func (t *TokenType) IsKeyword() bool { return *t < keyword_end && *t > keyword_beg }

func (t *TokenType) IsDecl() bool { return *t < decl_end && *t > decl_beg }

func (t *TokenType) IsStatment() bool { return *t < stmt_end && *t > stmt_beg }

func (t *TokenType) IsLit() bool { return *t < lit_end && *t > lit_beg }
