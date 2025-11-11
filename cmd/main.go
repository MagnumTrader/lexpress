package main

import (

	lx "github.com/MagnumTrader/lexpress"
)

type Token struct {
	value int
}

// TODO: what will i do
// - [x] Make it a way to create your own Tokenkind as an int
// - [ ] Make helper functions for establishing tokens
// - [ ] Create a default lexer which registers C like things?
// - [ ] Add a way to provide keywords, and what tokentype they should produce
// 			"KEYWORD", KEYWORD (constant)
// - [ ] think about error handling. how will this be handled and returned?
// - [ ] add a function for format TokenKind function
//       where you can create function that takes a token kind and return a string for formatting.

type myTokenKind = int

const (
	MYKIND myTokenKind = iota
	OTHERKIND
	LESS
	LESSEQUAL
	QUOTE
	STRING
)

func main() {

	lx := lx.NewLexer(
		"&|=",
		lx.MapRune('&', MYKIND),
		lx.MapRune('|', MYKIND),
		// how should this have a predefined string thing*
		lx.MapRuneFunc('"', func(h lx.LexerHandle) lx.TokenKind {
			// Hey gpt: Here, should we assume that we just consumed "
			// or is it up to the user to consume and proceed.
			// also should h have some consume token thing or should the lexer consume the token and move the cursor in the background?
			h.NextRune()
			// logic here for consuming until "

			return STRING
		}),
	)
	lx.PrintAll()
}
