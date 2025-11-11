package lexpress

import (
	"fmt"
	"log"
)

type Token struct {
	Kind    TokenKind
	line    int
	linePos int
	byte    int
}

// Specify your own kinds, EOF is declared as -1 so consumer should start with consts from iota >
type TokenKind = int

const (
	EOF TokenKind = iota -1
)

// A function that should advance until your token end is found
// returns the kind of token the lexer is currently posititioned at
type LexerHandleFn = func(h LexerHandle) TokenKind
type RuneHandler = func(l *lexer)


type lexer struct {
	singleRune   map[rune]TokenKind
	runeFunction map[rune]LexerHandleFn
}

// Main function used by your parser
// will cusome runes lazily and produce the next token
func (l *lexer) Next() Token {
	panic("Not implemented, should get the next rune, parse it through the registered runes and handle it. returning the Token")
}

func NewLexer(handlers ...RuneHandler) *lexer {
	l := &lexer{
		singleRune:   map[rune]TokenKind{},
		runeFunction: map[rune]LexerHandleFn{},
	}

	// register all the routes
	for _, handler := range handlers {
		handler(l)
	}
	return l
}

func MapRune(r rune, kind TokenKind) func(*lexer) {
	return func(l *lexer) {
		l.register(r, kind)
	}
}

func MapRuneFunc(r rune, f LexerHandleFn) func(*lexer) {
	return func(l *lexer) {
		l.registerWith(r, f)
	}
}

func MapRunes(kind TokenKind, runes ...rune) func(*lexer) {
	return func(l *lexer) {
		for _, r := range runes {
			l.register(r, kind)
		}
	}
}

func (l lexer) registerWith(r rune, f LexerHandleFn) {
	_, ok := l.runeFunction[r]
	_, ok2 := l.singleRune[r]
	if ok || ok2 {
		log.Fatalf("Handler for rune: '%c' already exists!", r)
	}
	l.runeFunction[r] = f
}

func (l *lexer) register(r rune, kind TokenKind) {
	_, ok := l.singleRune[r]
	_, ok2 := l.runeFunction[r]
	if ok || ok2 {
		log.Fatalf("Handler for rune: '%c' already exists!", r)
	}
	l.singleRune[r] = kind
}

func (l *lexer) PrintAll() {
	for key, value := range l.singleRune {
		fmt.Printf("'%c' registered as %v\n", key, value)
	}
	for key, value := range l.runeFunction {
		fmt.Printf("'%c' registered as %v\n", key, value)
	}
}
