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

// FuzzDecimal_General validates that Decimal only returns digits, hyphens, and dots.
func FuzzDecimal_General(f *testing.F) {
	seed := []string{
		"The price is -123.45 USD",
		"Balance: 0.001234",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.Decimal(input)
		for _, r := range out {
			valid := unicode.IsDigit(r) || r == '-' || r == '.'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzDomain_General validates that Domain only returns valid domain characters when no error.
func FuzzDomain_General(f *testing.F) {
	seed := []struct {
		input        string
		preserveCase bool
		removeWww    bool
	}{
		{"https://www.Example.com", false, true},
		{"example.COM", true, false},
	}
	for _, tc := range seed {
		f.Add(tc.input, tc.preserveCase, tc.removeWww)
	}
	f.Fuzz(func(t *testing.T, input string, preserveCase, removeWww bool) {
		out, err := sanitize.Domain(input, preserveCase, removeWww)
		if err != nil {
			return
		}
		for _, r := range out {
			valid := unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '-'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzEmail_General validates that Email only returns valid email characters.
func FuzzEmail_General(f *testing.F) {
	seed := []struct {
		input        string
		preserveCase bool
	}{
		{"mailto:Person@Example.COM", false},
		{"test+1@EXAMPLE.com", true},
	}
	for _, tc := range seed {
		f.Add(tc.input, tc.preserveCase)
	}
	f.Fuzz(func(t *testing.T, input string, preserveCase bool) {
		out := sanitize.Email(input, preserveCase)
		for _, r := range out {
			valid := unicode.IsLetter(r) || unicode.IsDigit(r) ||
				r == '-' || r == '_' || r == '.' || r == '@' || r == '+'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}
