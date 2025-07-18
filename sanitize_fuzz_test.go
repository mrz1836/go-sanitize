package sanitize_test

import (
	"net"
	"regexp"
	"testing"
	"unicode"

	"github.com/mrz1836/go-sanitize"
	"github.com/stretchr/testify/require"
)

// FuzzAlpha_Basic validates that Alpha only returns letters and optional spaces.
func FuzzAlpha_Basic(f *testing.F) {
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

// FuzzAlphaNumeric validates that AlphaNumeric only returns letters, digits, and optional spaces.
func FuzzAlphaNumeric(f *testing.F) {
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

// FuzzBitcoinAddress validates that BitcoinAddress only returns valid Bitcoin address characters.
func FuzzBitcoinAddress(f *testing.F) {
	seed := []string{
		":1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!",
		"OIl01K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!",
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
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

// FuzzBitcoinCashAddress validates that BitcoinCashAddress only returns valid Bitcoin Cash address characters.
func FuzzBitcoinCashAddress(f *testing.F) {
	seed := []string{
		"$#:qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz!",
		"pqbq3728yw0y47sOqn6l2na30mcw6zm78idzq5ucqzc371",
		"1PUwPCNqKiC6La8wtbJEAhnBvtc8gdw19h",
	}
	for _, tc := range seed {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.BitcoinCashAddress(input)
		for _, r := range out {
			valid := r == '0' ||
				(r >= '2' && r <= '9') ||
				r == 'a' ||
				(r >= 'c' && r <= 'h') ||
				(r >= 'j' && r <= 'n') ||
				(r >= 'p' && r <= 'z') ||
				r == 'A' ||
				(r >= 'C' && r <= 'H') ||
				(r >= 'J' && r <= 'N') ||
				(r >= 'P' && r <= 'Z')
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzDecimal validates that Decimal only returns digits, hyphens, and dots.
func FuzzDecimal(f *testing.F) {
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

// FuzzDomain validates that Domain only returns valid domain characters when no error.
func FuzzDomain(f *testing.F) {
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

// FuzzEmail validates that Email only returns valid email characters.
func FuzzEmail(f *testing.F) {
	seed := []struct {
		input        string
		preserveCase bool
	}{
		{"mailto:Person@Example.COM", false},
		{"test+1@EXAMPLE.com", true},
		{"user.name+tag+sorting@example.co.uk", false},
		{"UPPER_lower-123@sub.domain.com", true},
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

// FuzzFirstToUpper validates that FirstToUpper capitalizes the first letter of the input string.
func FuzzFirstToUpper(f *testing.F) {
	seed := []string{"example", "Already Upper", ""}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.FirstToUpper(input)
		inRunes := []rune(input)

		outRunes := []rune(out)
		if len(inRunes) == 0 {
			require.Empty(t, outRunes)
			return
		}

		require.Len(t, outRunes, len(inRunes))
		require.Equal(t, unicode.ToUpper(inRunes[0]), outRunes[0])
		require.Equal(t, inRunes[1:], outRunes[1:])
	})
}

// FuzzFormalName validates that FormalName only returns valid characters for names.
func FuzzFormalName(f *testing.F) {
	seed := []string{
		"Mark Mc'Cuban-Host",
		"Does #Not Work!",
		"O'Leary-Brown",
		"d'Artagnan",
		"D’Angelo",
		"Van  der  Meer",
		"Émilie du Châtelet",
		"Björk Guðmundsdóttir",
	}
	for _, tc := range seed {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.FormalName(input)
		for _, r := range out {
			valid := unicode.IsLetter(r) || unicode.IsDigit(r) ||
				r == '-' || r == '\'' || r == ',' || r == '.' || unicode.IsSpace(r)
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzHTML validates that HTML removes all HTML tags.
func FuzzHTML(f *testing.F) {
	seed := []string{
		"<div>Hello <b>World</b></div>",
		"Plain <b>text</b>",
	}
	for _, tc := range seed {
		f.Add(tc)
	}

	htmlPattern := regexp.MustCompile(`(?i)<[^>]*>`)

	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.HTML(input)
		require.False(t, htmlPattern.MatchString(out), "output %q still contains HTML tags", out)
	})
}

// FuzzIPAddress validates that IPAddress only returns valid IP addresses.
func FuzzIPAddress(f *testing.F) {
	seed := []string{"192.168.0.1", "2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f", "::ffff:192.0.2.128", "bad"}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.IPAddress(input)
		if out == "" {
			return
		}

		ip := net.ParseIP(out)
		require.NotNilf(t, ip, "output %q is not a valid IP", out)
		require.Equal(t, ip.String(), out)
	})
}

// FuzzNumeric validates that Numeric only returns digits.
func FuzzNumeric(f *testing.F) {
	seed := []string{
		"Phone: 123-456-7890",
		"Order #987654321",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.Numeric(input)
		for _, r := range out {
			require.Truef(t, unicode.IsDigit(r),
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzPhoneNumber validates that PhoneNumber only returns digits and plus signs.
func FuzzPhoneNumber(f *testing.F) {
	seed := []string{
		"+1 (234) 567-8900",
		"(555)555-5555 ext.123",
		"tel:+44 20 7946 0958",
	}
	for _, tc := range seed {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.PhoneNumber(input)
		for _, r := range out {
			valid := unicode.IsDigit(r) || r == '+'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzPathName validates that PathName only returns valid pathname characters.
func FuzzPathName(f *testing.F) {
	seed := []string{
		"file:name/with*invalid|chars_week",
		"another path\\with spaces.txt",
	}
	for _, tc := range seed {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.PathName(input)
		for _, r := range out {
			valid := unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzPunctuation validates that Punctuation only returns letters, digits, spaces, and standard punctuation.
func FuzzPunctuation(f *testing.F) {
	seed := []string{
		"Hello, World! How's it going? (Good, I hope.) ``[]{}",
		"Testing #1 & #2: \"quotes\" and punctuation!",
	}
	for _, tc := range seed {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.Punctuation(input)
		for _, r := range out {
			valid := unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '\'' ||
				r == '"' || r == '#' || r == '&' || r == '!' || r == '?' || r == ',' ||
				r == '.' || unicode.IsSpace(r)
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzScientificNotation validates that ScientificNotation only returns digits, dots, and exponent characters.
func FuzzScientificNotation(f *testing.F) {
	seed := []string{
		" String: 1.23e-3 ",
		"$1.0E+10",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.ScientificNotation(input)
		for _, r := range out {
			valid := unicode.IsDigit(r) || r == '.' || r == 'e' || r == 'E' || r == '+' || r == '-'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzScripts validates that Scripts removes script and embed tags.
func FuzzScripts(f *testing.F) {
	seed := []string{
		"<script>alert('x')</script>",
		"<iframe src='t'></iframe>",
	}
	for _, tc := range seed {
		f.Add(tc)
	}

	scriptPattern := regexp.MustCompile(`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`)

	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.Scripts(input)
		require.False(t, scriptPattern.MatchString(out), "output %q still contains script tags", out)
	})
}

// FuzzSingleLine validates that SingleLine removes all newline characters.
func FuzzSingleLine(f *testing.F) {
	seed := []string{
		"First\nSecond",
		"Tab\tSeparated",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.SingleLine(input)
		require.NotContains(t, out, "\r")
		require.NotContains(t, out, "\n")
		require.NotContains(t, out, "\t")
		require.NotContains(t, out, "\v")
		require.NotContains(t, out, "\f")
	})
}

// FuzzTime validates that Time only returns digits and colons.
func FuzzTime(f *testing.F) {
	seed := []string{
		"t00:00d -EST",
		"Time 12:34:56!",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.Time(input)
		for _, r := range out {
			valid := unicode.IsDigit(r) || r == ':'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzURI validates that URI only returns valid URI characters.
func FuzzURI(f *testing.F) {
	seed := []string{
		"/This/Works/?that=123&this#page10%",
		"Test?=what! &this=that",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.URI(input)
		for _, r := range out {
			valid := unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '/' ||
				r == '?' || r == '&' || r == '=' || r == '#' || r == '%'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}

// FuzzURL validates that URL only returns valid URL characters.
func FuzzURL(f *testing.F) {
	seed := []string{
		"https://domain.com/this/test?this=value&another=123%#page",
		"https://Example.com/This/Works?^No&this",
	}
	for _, tc := range seed {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		out := sanitize.URL(input)
		for _, r := range out {
			valid := unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '/' ||
				r == ':' || r == '.' || r == ',' || r == '?' || r == '&' || r == '@' ||
				r == '=' || r == '#' || r == '%'
			require.Truef(t, valid,
				"invalid rune %q in %q (input: %q)", r, out, input)
		}
	})
}
