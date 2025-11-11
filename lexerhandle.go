package lexpress

// Way of exposing the lexer functionality without exposing the struct itself
type LexerHandle interface {
	// Used internally by the lexer, but is also exposed in `LexerHandleFn` for advancing the lexer
	NextRune() rune
}

// Used internally by the lexer, but is also exposed in `LexerHandleFn` for advancing the lexer
func (l *lexer) NextRune() rune {
	return ' '
}


