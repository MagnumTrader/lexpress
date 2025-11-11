package main

import "github.com/MagnumTrader/lexpress/lexer"

type Token struct {
	value int
}

func main() {

	lex := lexer.NewLexer(
		lexer.RuneToToken('&', Token{}),
		lexer.RuneToTokenWith('|', NextEqualElse()),
		lexer.RuneToTokenWith('¤', HandleOther),
	)
	lex.PrintAll()
}
func HandleOther(h lexer.LexerHandle) *Token { return &Token{} }
func NextEqualElse() lexer.LexerHandleFn[Token] {
	return func(h lexer.LexerHandle) *Token {
		if h.NextRune() == '=' {
			h.NextRune()
		}
		return &Token{}
	}

}
