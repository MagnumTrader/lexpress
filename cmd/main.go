package main

import lx "github.com/MagnumTrader/lexpress"

type Token struct {
	value int
}

// TODO: what will i do
// - [ ] Make it a way to create your own Tokenkind as an int
// - [ ] Make helper functions for establishing tokens
// - [ ] Rune indexing into array instead of hashmap
// - [ ] Create a default lexer which registers C like things?
// - [ ] Add a way to provide keywords, and what tokentype they should produce
// 			"KEYWORD", KEYWORD (constant)
// - [ ] think about error handling. how will this be handled and returned?
// - [ ] EOF should be -

type myTokenKind = int

const (
	MYKIND myTokenKind = iota
	OTHERKIND
	LESS
	LESSEQUAL
)

func main() {

	lex := lx.NewLexer(
		lx.RuneToToken('&', MYKIND),
	)
	lex.PrintAll()
}
