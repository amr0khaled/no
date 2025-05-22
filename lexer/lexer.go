package lexer

import (
	"lang/io"
	"lang/token"
	"strings"
)

type Lex struct {
	Line   int
	Offset int
	Value  []string
}

var prev int = 0

var base int = 0

var current int = 0

type LexText string

var text LexText = ""

var next int = 1

var prev_w int = 0 //start of previous operator or litral or expression

var exprNum int = 0

var exprs []Lex
var offset int = 0

func Lexer(file *print.File) *[]token.Statement {
	for _, fline := range file.Lines {
		ResetCursor()
		var expr []string
		var word string = ""
		offset = 0
		lex := &Lex{
			Line:   fline.Offset,
			Offset: offset,
		}
		text = LexText(fline.Text)
		for {
			if current >= len(text) {
				AppendWord(&expr, &word)
				lex.Value = expr
				break
			}
			if isEOC(text.peek()) {
				AppendWord(&expr, &word)
				text.peekNext()
				AppendExpr(lex, &expr)
				AppendLex(&exprs, lex)
			}
			if isWhite(text.peek()) {
				AppendWord(&expr, &word)
			} else {
				word += text.consume()
			}
		}
		exprs = append(exprs, *lex)
	}
	s := Tokenize(&exprs)
	return s
}

func Tokenize(lexs *[]Lex) *[]token.Statement {
	var tokens []token.Statement
	for _, lex := range *lexs {
		var _tokens token.Statement
		for _, v := range lex.Value {
			tok := token.GetToken(v)
			keyword := token.Tokenize(tok)
			if strings.Contains(keyword, "token") {
				tok = token.IDENT
			}
			token := &token.Token{
				Type:   tok,
				Lexeme: v,
				Value:  token.GetValue(v),
			}
			AppendToken(&_tokens, token)
		}
		AppendTokens(&tokens, &_tokens)
	}
	return &tokens
}

// Utils

func AppendTokens(tokens *[]token.Statement, stmt *token.Statement) {
	*tokens = append(*tokens, *stmt)
}

func AppendToken(tokens *token.Statement, token *token.Token) {
	*tokens = append(*tokens, *token)
}

func AppendWord(expr *[]string, word *string) {
	if *word != "" {
		*expr = append(*expr, *word)
		*word = ""
	}
	nextChar()
}

func AppendExpr(lex *Lex, expr *[]string) {
	lex.Value = *expr
	offset++
	lex.Offset = offset
	*expr = []string{}
}
func AppendLex(exprs *[]Lex, lex *Lex) {
	*exprs = append(*exprs, *lex)

}

func (t *LexText) prev() string {
	return string((*t)[prev])
}
func nextChar() {
	prev++
	current++
	next++
}
func ResetCursor() {
	prev = -1
	current = 0
	next = 1
}

func (t *LexText) peek() string {
	return string((*t)[current])
}

func (t *LexText) consume() string {
	if current >= len(string(*t)) {
		return ""
	}
	v := string((*t)[current])
	nextChar()
	return v
}

func (t *LexText) peekNext() string {
	return string((*t)[next])
}

func isEOC(s string) bool   { return s == ";" || s == "\n" }
func isWhite(s string) bool { return s == " " || s == "\t" }
