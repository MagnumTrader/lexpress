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
		f func()(rune, error)
		  
	}{
		{
			expect:    '%',
			expectErr: nil,
			f: l.PeekRune,
		},
		{
			expect:    '%',
			expectErr: nil,
			f: l.PeekRune,
		},
		{
			expect:    '%',
			expectErr: nil,
			f: l.NextRune,
		},
		{
			expect:    '|',
			expectErr: nil,
			f: l.PeekRune,
		},
		{
			expect:    '|',
			expectErr: nil,
			f: l.PeekRune,
		},
		{
			expect:    '|',
			expectErr: nil,
			f: l.NextRune,
		},
		{
			expect:    utf8.RuneError,
			expectErr: io.EOF,
			f: l.NextRune,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			r, err := c.f()
			fmt.Printf("function yielded %c\n", r)
			if r!= c.expect || err != c.expectErr {
				t.Fatalf("failed, expected: '%c' and %v, got '%c' and %v", c.expect, c.expectErr, r, err)
			}
		})
	}
}
