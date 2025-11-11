package lexpress

import (
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
	EOF TokenKind = iota - 1
)

// A function that should advance until your token end is found
// returns the kind of token the lexer is currently posititioned at
type LexerHandleFn = func(h LexerHandle) TokenKind
type LexerOption = func(l *lexer)

type lexer struct {
	input      string
	lookup     map[rune]LexerHandleFn
	peeked     *Token
	peekedRune rune

	// maybe byte instead
	bytePos    int
	tokenStart int
	tokenEnd   int
	curLine    int
	curLinePos int
}

// Main function used by your parser
// will cusome runes lazily and produce the next token
func (l *lexer) Next() *Token {
	if l.atEof() {
		return &Token{
			Kind:    EOF,
			line:    l.curLine,
			linePos: l.curLinePos,
			byte:    l.bytePos,
		}
	}

	return l.takeToken()
}

func NewLexer(input string, handlers ...LexerOption) *lexer {
	l := &lexer{
		input:      input,
		lookup:     map[rune]LexerHandleFn{},
		peeked:     nil,
		peekedRune: 0,
		bytePos:    0,
		tokenEnd:   1,
	}

	// register all the routes
	for _, handler := range handlers {
		handler(l)
	}
	return l
}
func (l *lexer) Peek() *Token {
	panic("unimplemented")
}

// Pointers into the string, or should we use a string reader?
func (l *lexer) takeToken() *Token {
	if l.bytePos >= len(l.input) {

	}

	return nil
}

func (l *lexer) atEof() bool {
	return l.bytePos >= len(l.input)
}


func MapRune(r rune, kind TokenKind) LexerOption {
	return MapRuneFunc(r, func(h LexerHandle) TokenKind { return kind })
}

func MapRuneFunc(r rune, f LexerHandleFn) LexerOption {
	return func(l *lexer) {
		l.expectNotRegistered(r)
		l.lookup[r] = f
	}
}

// TODO: capture to including the token
// TODO: capture inside including the token = capture to [1:pos-1]

func (l *lexer) expectNotRegistered(r rune) {
	_, ok := l.lookup[r]
	if ok {
		// TODO: collect all these errors instead
		log.Fatalf("Handler for rune: '%c' already exists!", r)
	}
}

// Print all registered routes
func (l *lexer) PrintAll() {
	//	for key, value := range l.singleRune {
	//		fmt.Printf("'%c' registered as %v\n", key, value)
	//	}
	//
	//	for key, value := range l.runeFunction {
	//		fmt.Printf("'%c' registered as %v\n", key, value)
	//	}
}
