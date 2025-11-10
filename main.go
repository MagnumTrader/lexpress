package main

import "testinggrounds/lexer"

type Token struct {
	value int
}

func main() {
	// Alternative is that register take a function, that returns a
	// and then we have quick token
	// We move the cursor to where we want, and then the lexer creates the token?
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
