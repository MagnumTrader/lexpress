package lexer

import (
	"fmt"
	"log"
)

type lexer[T any] struct {
	singleRune   map[rune]*T
	runeFunction map[rune]LexerHandleFn[T]
}

func (l *lexer[T]) NextRune() rune {
	return ' '
}

func (l *lexer[T]) Next() *T {
	rune := '%'
	t, ok := l.singleRune[rune]
	if !ok {
		fn, ok := l.runeFunction[rune]

		if !ok {
			log.Fatalf("No handler for %c", rune)
		}

		return fn(l)	
	}
	return t
}

type LexerHandle interface {
	NextRune() rune
}

type LexerHandleFn[T any] = func(h LexerHandle) *T
type RuneHandler[T any] = func(l *lexer[T])

func NewLexer[T any](handlers ...RuneHandler[T]) *lexer[T] {
	l := &lexer[T]{
		singleRune:   map[rune]*T{},
		runeFunction: map[rune]LexerHandleFn[T]{},
	}

	for _, handler := range handlers {
		handler(l)
	}
	return l
}

func RuneToTokenWith[T any](r rune, f LexerHandleFn[T]) func(*lexer[T]) {
	return func(l *lexer[T]) {
		l.registerWith(r, f)
	}
}
func RuneToToken[T any](r rune, token T) func(*lexer[T]) {
	return func(l *lexer[T]) {
		l.register(r, token)
	}
}

func AllRunesToToken[T any](token T, runes ...rune) func(*lexer[T]) {
	return func(l *lexer[T]) {
		for _, r := range runes {
			// TODO: When we collect this token, enrich it with position and line for example?
			l.register(r, token)
		}
	}
}

func (l lexer[T]) registerWith(r rune, f LexerHandleFn[T]) {
	_, ok := l.runeFunction[r]
	_, ok2 := l.singleRune[r]
	if ok || ok2 {
		log.Fatalf("Handler for rune: '%c' already exists!", r)
	}
	l.runeFunction[r] = f
}

func (l *lexer[T]) register(r rune, token T) {
	_, ok := l.singleRune[r]
	_, ok2 := l.runeFunction[r]
	if ok || ok2 {
		log.Fatalf("Handler for rune: '%c' already exists!", r)
	}
	l.singleRune[r] = &token
}

func (l *lexer[T]) PrintAll() {
	for key, value := range l.singleRune {
		fmt.Printf("'%c' registered as %v\n", key, *value)
	}
	for key, value := range l.runeFunction {
		fmt.Printf("'%c' registered as %v\n", key, value)
	}
}
