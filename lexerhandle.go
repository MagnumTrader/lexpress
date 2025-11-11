package lexpress

import (
	"io"
	"log"
	"unicode/utf8"
)

// Way of exposing the lexer functionality without exposing the struct itself
type LexerHandle interface {
	// Used internally by the lexer, but is also exposed in `LexerHandleFn` for advancing the lexer
	NextRune() (rune, error)
	PeekRune() (rune, error)
	AdvanceUntil(until rune) error
}

func (l *lexer) nextRune() (rune, int, error) {
	if l.bytePos >= len(l.input) {
		return utf8.RuneError, 0, io.EOF
	}

	r, size := utf8.DecodeRuneInString(l.input[l.bytePos:])

	if r == utf8.RuneError && size == 1 {
		log.Fatalf("Invalid utf8 encoding in string at byte %d and line %d", l.bytePos, l.curLine)
	}
	return r, size, nil
}

// Used internally by the lexer, but is also exposed in `LexerHandleFn` for advancing the lexer
func (l *lexer) NextRune() (rune, error) {
	if l.peekedRune != 0 {
		// Reseting Peek
		r := l.peekedRune
		l.peekedRune = 0

		l.bytePos += utf8.RuneLen(r)
		return r, nil
	}

	r, size, err := l.nextRune()
	if err != nil {
		return r, err
	}

	l.bytePos += size
	return r, nil
}

// Used internally by the lexer, but is also exposed in `LexerHandleFn` for advancing the lexer
func (l *lexer) PeekRune() (rune, error) {
	if l.peekedRune != 0 {
		return l.peekedRune, nil
	}

	r, _, err := l.nextRune()
	if err != nil {
		// eof
		return r, err
	}

	l.peekedRune = r
	return r, nil
}

func (l *lexer) AdvanceUntil(until rune) error {
	for {
		r, err := l.NextRune()
		if err != nil {
			return err
		}
		if r == until {
			break
		}
	}
	return nil
}

