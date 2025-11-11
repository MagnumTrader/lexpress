package lexpress

import (
	"fmt"
	"io"
	"testing"
	"unicode/utf8"
)

func TestNextRune(t *testing.T) {
	l := NewLexer("%|")

	cases := []struct {
		expect    rune
		expectErr error
		f         func() (rune, error)
	}{
		{
			expect:    '%',
			expectErr: nil,
			f:         l.PeekRune,
		},
		{
			expect:    '%',
			expectErr: nil,
			f:         l.PeekRune,
		},
		{
			expect:    '%',
			expectErr: nil,
			f:         l.NextRune,
		},
		{
			expect:    '|',
			expectErr: nil,
			f:         l.PeekRune,
		},
		{
			expect:    '|',
			expectErr: nil,
			f:         l.PeekRune,
		},
		{
			expect:    '|',
			expectErr: nil,
			f:         l.NextRune,
		},
		{
			expect:    utf8.RuneError,
			expectErr: io.EOF,
			f:         l.NextRune,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			r, err := c.f()
			if r != c.expect || err != c.expectErr {
				t.Fatalf("failed, expected: '%c' and %v, got '%c' and %v", c.expect, c.expectErr, r, err)
			}
		})
	}
}

func TestNextKind(t *testing.T) {
	l := NewLexer(
		`%|#"hello"`,
		MapRune('|', 1),
		MapRune('%', 2),
		MapRune('#', 3),
		// Rune is included, the pointer still includes
		// this should be a between function
		// we should have a function for that a LexerHandlerBetween -> LexerHandler
		// where we consume the first, then we generate the token
		// then we pass 
		MapRuneFunc('"', func(h LexerHandle) TokenKind {
			err := h.AdvanceUntil('"')
			if err != nil {
			  panic(err)
			}
			return 4
		}))

	cases := []struct {
		expectKind  TokenKind
		expectLexem string
		expectErr   error
	}{
		{
			expectKind:  2,
			expectLexem: "%",
			expectErr:   nil,
		},
		{
			expectKind:  1,
			expectLexem: "|",
			expectErr:   nil,
		},
		{
			expectKind:  3,
			expectLexem: "#",
			expectErr:   nil,
		},
		{
			expectKind:  4,
			expectLexem: `"hello"`,
			expectErr:   nil,
		},
		{
			expectKind:  EOF,
			expectLexem: "",
			expectErr:   io.EOF,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			// Should i include "" or not
			token := l.Next()
			if token == nil {
				return
			}
			if c.expectKind != token.Kind || c.expectLexem != token.Lexeme {
				t.Fatalf("failed, expected: %d '%s' , got %d '%s' ", c.expectKind, c.expectLexem, token.Kind, token.Lexeme)
			}
		})
	}
}
