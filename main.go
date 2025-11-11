package main

import "github.com/MagnumTrader/lexpress/lexer"

type Token struct {
	value int
}

// TODO: what will i do
// remove generics
// - [ ] Make it a way to create your own Tokenkind as an int
// - [ ] Make helper functions for establishing tokens
// - [ ] Rune indexing into array instead of hashmap
// - [ ] Create a default lexer which registers C like things?
// - [ ] Add a way to provide keywords, and what tokentype they should produce
// 			"KEYWORD", KEYWORD (constant)
// - [ ] think about error handling. how will this be handled and returned?

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
