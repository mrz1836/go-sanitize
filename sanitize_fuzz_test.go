package sanitize_test

import (
	"testing"
	"unicode"

	"github.com/mrz1836/go-sanitize"
)

// FuzzAlphaNumeric validates that AlphaNumeric only returns letters, digits and optional spaces.
func FuzzAlphaNumeric_General(f *testing.F) {
	seed := []struct {
		input  string
		spaces bool
	}{
		{"Example 123!", false},
		{"Another Example 456?", true},
	}
	for _, tc := range seed {
		f.Add(tc.input, tc.spaces)
	}
	f.Fuzz(func(t *testing.T, input string, spaces bool) {
		out := sanitize.AlphaNumeric(input, spaces)
		for _, r := range out {
			if spaces && r == ' ' {
				continue
			}
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				t.Fatalf("invalid rune %q in %q", r, out)
			}
		}
	})
}
