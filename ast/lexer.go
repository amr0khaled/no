package ast

import "lang/token"

type Lex string

var prev int = -1

var base int = 0

var current int = 0

var prev_w int = 0 //start of previous operator or litral or expression

func Lexer(line string) []string {
	var expr []string
	var word string = ""
	for i, chr := range line {
		chr := string(chr)
		if isEOC(chr) {
			expr = append(expr, Lexer(line[i+1:])...)
		}

		if !isWhite(chr) {
			word += chr
		} else if isWhite(chr) {
			if word != "" {
				expr = append(expr, word)
			}
			word = ""
		}
	}
	return expr
}

//func (l *Lex) prev() string {
//	current--
//	return l[current]
//}
//
//func (l *Lex) next() string {
//	current++
//	return l[current]
//}

func isEOC(s string) bool   { return s == ";" || s == "\n" }
func isWhite(s string) bool { return s != " " || s != "\t" }
