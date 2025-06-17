package sanitize_test

import (
	"testing"
	"unicode"

	"github.com/mrz1836/go-sanitize"
	"github.com/stretchr/testify/require"
)

// FuzzAlphaNumeric_General validates that AlphaNumeric only returns letters, digits, and optional spaces.
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
			require.Truef(t, unicode.IsLetter(r) || unicode.IsDigit(r),
				"invalid rune %q in %q (input: %q, spaces: %v)", r, out, input, spaces)
		}
	})
}

// FuzzAlpha_General validates that Alpha only returns letters and optional spaces.
func FuzzAlpha_General(f *testing.F) {
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
		out := sanitize.Alpha(input, spaces)
		for _, r := range out {
			if spaces && r == ' ' {
				continue
			}
			require.Truef(t, unicode.IsLetter(r),
				"invalid rune %q in %q (input: %q, spaces: %v)", r, out, input, spaces)
		}
	})
}

// FuzzBitcoinAddress_General validates that BitcoinAddress only returns valid Bitcoin address characters.
func FuzzBitcoinAddress_General(f *testing.F) {
	seed := []string{
		":1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!",
		"OIl01K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.BitcoinAddress(input)
		for _, r := range out {
			valid := (r >= 'a' && r <= 'k') ||
				(r >= 'm' && r <= 'z') ||
				(r >= 'A' && r <= 'H') ||
				(r >= 'J' && r <= 'N') ||
				(r >= 'P' && r <= 'Z') ||
				(r >= '1' && r <= '9')
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}
