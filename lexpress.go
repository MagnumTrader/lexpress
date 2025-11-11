package lexpress

import (
	"fmt"
	"log"
)

type Token struct {
	Kind    int
	line    int
	linePos int
	byte    int
}

type lexer struct {
	singleRune   map[rune]Token
	runeFunction map[rune]LexerHandleFn
}

func (l *lexer) NextRune() rune {
	return ' '
}

func (l *lexer) Next() Token {
	rune := '%'
	t, ok := l.singleRune[rune]
	if !ok {
		fn, ok := l.runeFunction[rune]

		if !ok {
			// TODO: this should return unknown token
			log.Fatalf("No handler for %c", rune)
		}
		return fn(l)
	}
	return t
}

type LexerHandle interface {
	NextRune() rune
}

// A function that should advance until your token end is found
// returns the kind of token the lexer is currently posititioned at
type LexerHandleFn = func(h LexerHandle) Token
type RuneHandler = func(l *lexer)

func NewLexer(handlers ...RuneHandler) *lexer {
	l := &lexer{
		singleRune:   map[rune]Token{},
		runeFunction: map[rune]LexerHandleFn{},
	}

	for _, handler := range handlers {
		handler(l)
	}
	return l
}

func RuneToTokenWith[T any](r rune, f LexerHandleFn) func(*lexer) {
	return func(l *lexer) {
		l.registerWith(r, f)
	}
}

func RuneToToken(r rune, token Token) func(*lexer) {
	return func(l *lexer) {
		l.register(r, token)
	}
}

func AllRunesToToken(token Token, runes ...rune) func(*lexer) {
	return func(l *lexer) {
		for _, r := range runes {
			// TODO: When we collect this token, enrich it with position and line for example?
			l.register(r, token)
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

func (l *lexer) register(r rune, token Token) {
	_, ok := l.singleRune[r]
	_, ok2 := l.runeFunction[r]
	if ok || ok2 {
		log.Fatalf("Handler for rune: '%c' already exists!", r)
	}
	l.singleRune[r] = token
}

func (l *lexer) PrintAll() {
	for key, value := range l.singleRune {
		fmt.Printf("'%c' registered as %v\n", key, value)
	}
	for key, value := range l.runeFunction {
		fmt.Printf("'%c' registered as %v\n", key, value)
	}
}
