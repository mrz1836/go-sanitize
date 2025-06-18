package sanitize_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/mrz1836/go-sanitize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAlpha tests the alpha sanitize method
func TestAlpha(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		typeCase bool
	}{
		// Basic cases
		{"regular string", "Test This String-!123", "TestThisString", false},
		{"various symbols", `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`, "Symbols", false},
		{"carriage returns", "\nThis\nThat", "ThisThat", false},
		{"quotes and ticks", "“This is a quote with tick`s … ” ☺ ", "Thisisaquotewithticks", false},
		{"spaces", "Test This String-!123", "Test This String", true},
		{"symbols and spaces", `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`, "Symbols", true},
		{"quotes and spaces", "“This is a quote with tick`s … ” ☺ ", "This is a quote with ticks    ", true},
		{"carriage returns with spaces", "\nThis\nThat", `ThisThat`, true},

		// Edge cases
		{"empty string", "", "", false},
		{"only special characters", "!@#$%^&*()", "", false},
		{"very long string", strings.Repeat("a", 1000), strings.Repeat("a", 1000), false},
		{"tabs", "\tThis1\tThat2", `ThisThat`, true},
		{"carriage returns with n", "\nThis1\nThat2", `ThisThat`, true},
		{"carriage returns with r", "\rThis1\rThat2", `ThisThat`, true},
		{"accented characters", "éclair", "éclair", false},
		{"greek characters", "Σigma", "Σigma", false},
		{"sharp s", "ßeta", "ßeta", false},
		{"numbers only", "123456", "", false},
		{"spaces only", "   ", "   ", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Alpha(test.input, test.typeCase)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestAlphaNumeric tests the alphanumeric sanitize method
func TestAlphaNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		typeCase bool
	}{
		// Basic cases
		{"regular string", "Test This String-!123", "TestThisString123", false},
		{"symbols", `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`, "Symbols", false},
		{"carriage returns", "\nThis1\nThat2", "This1That2", false},
		{"quotes and ticks", "“This is a quote with tick`s … ” ☺ 342", "Thisisaquotewithticks342", false},
		{"string with spaces", "Test This String-! 123", "Test This String 123", true},
		{"symbols and spaces", `~!@#$%^&*()-_Symbols 123=+[{]};:'"<>,./?`, "Symbols 123", true},
		{"ticks and spaces", "“This is a quote with tick`s…”☺ 123", "This is a quote with ticks 123", true},
		{"carriage returns with n", "\nThis1\nThat2", `This1That2`, true},
		{"carriage returns with r", "\rThis1\rThat2", `This1That2`, true},
		{"tabs", "\tThis1\tThat2", `This1That2`, true},

		// Edge cases
		{"empty string", "", "", false},
		{"spaces only", "   ", "   ", true},
		{"accents and numbers", "éclair123", "éclair123", false},
		{"mixed unicode", "ßeta Σigma 456", "ßeta Σigma 456", true},
		{"numbers only", "987654", "987654", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.AlphaNumeric(test.input, test.typeCase)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestBitcoinAddress will test all permutations
func TestBitcoinAddress(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"remove symbol", ":1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs", "1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs"},
		{"remove spaces", "   1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs    ", "1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs"},
		{"remove spaces 2", "   1K6c7 LGpdB 8LwoGNVfG5 1dRV 9UUEijbrWs    ", "1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs"},
		{"remove symbols 2", "$#:1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!", "1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs"},
		{"remove symbols 3", "$#:1K6c_7LGpd^B8Lw_oGN=VfG+51_dRV9-UUEijbrWs!", "1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs"},

		// No uppercase letter O, uppercase letter I, lowercase letter l, and the number 0
		{"uppercase letters", "OIl01K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!", "1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs"},

		// Additional edge cases with normal, vanity, and rare addresses
		{"vanity address", "1CounterpartyXXXXXXXXXXXXXXXUWLpVr", "1CounterpartyXXXXXXXXXXXXXXXUWLpVr"},
		{"burn address", "1111111111111111111114oLvT2", "1111111111111111111114oLvT2"},
		{"remove punctuation", "1BoatSLRHtKNngkdXEeobR76b53LETtpyT!!", "1BoatSLRHtKNngkdXEeobR76b53LETtpyT"},
		{"remove spaces around", " 17SkEw2md5avVNyYgj6RiXuQKNwkXaxFyQ ", "17SkEw2md5avVNyYgj6RiXuQKNwkXaxFyQ"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.BitcoinAddress(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestBitcoinCashAddress will test all permutations of using BitcoinCashAddress()
func TestBitcoinCashAddress(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"remove symbols", "$#:qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz!", "qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz"},
		{"remove spaces", " $#:qze7yy2 au5vuznvn8lzj5y0j5t066 vhs75e3m0eptz! ", "qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz"},
		{"remove ignored characters", "pqbq3728yw0y47sOqn6l2na30mcw6zm78idzq5ucqzc371", "pqq3728yw0y47sqn6l2na30mcw6zm78dzq5ucqzc37"},
		{"basic cashaddr", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"},
		{"remove punctuation", "bitcoincash:qqq3728yw0y47sqn6l2na30mcw6zm78dzq5ucqzc37!!", "tcncashqqq3728yw0y47sqn6l2na30mcw6zm78dzq5ucqzc37"},
		{"remove spaces", " qr95tpm9f6qt8azfzd73ydyccdefhkcdv3ldk00ht0 ", "qr95tpm9f6qt8azfzd73ydyccdefhkcdv3ldk00ht0"},

		// Additional test cases
		{"empty string", "", ""},
		{"only symbols", "!@#$%^&*()", ""},
		{"mixed case address", "QPM2QSZNHKS23Z7629MMS6S4CWEF74VCWVY22GDX6A", "QPM2QSZNHKS23Z7629MMS6S4CWEF74VCWVY22GDX6A"},
		{"address with newlines", "\nqpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a\n", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"},
		{"address with tabs", "\tqpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a\t", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"},
		{"address with unicode", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a世界", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"}, //nolint:gosmopolitan // Unicode characters are not valid in Bitcoin Cash addresses
		{"address with dashes", "qpm2qszn-hks23z7629mms6s4cwef74vcwvy22gdx6a", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"},
		{"address with numbers only", "1234567890", "234567890"},
		{"address with prefix and suffix", "!!qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a!!", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"},
		{"address with internal spaces", "qpm2qszn hks23z7629mms6s4cwef74vcwvy22gdx6a", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"},
		{"address with mixed valid and invalid", "qpm2qszn!@#hks23z76*29mms6s4cwef74vcwvy22gdx6a", "qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.BitcoinCashAddress(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestCustom tests the custom sanitize method
func TestCustom(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
		regex    string
	}{
		{"", "ThisWorks123!", "ThisWorks123", `[^a-zA-Z0-9]`},
		{"", "ThisWorks1.23!", "1.23", `[^0-9.-]`},
		{"", "ThisWorks1.23!", "ThisWorks123", `[^0-9a-zA-Z]`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Custom(test.input, test.regex)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestCustomCompiled verifies CustomCompiled using a precompiled regex
func TestCustomCompiled(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		re       *regexp.Regexp
	}{
		{"alpha numeric", "Works 123!", "Works123", regexp.MustCompile(`[^a-zA-Z0-9]`)},
		{"decimal", "ThisWorks1.23!", "1.23", regexp.MustCompile(`[^0-9.-]`)},
		{"numbers and letters", "ThisWorks1.23!", "ThisWorks123", regexp.MustCompile(`[^0-9a-zA-Z]`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := sanitize.CustomCompiled(tt.input, tt.re)
			assert.Equal(t, tt.expected, output)
		})
	}
}

// TestCustomCompiled_NilRegex verifies that CustomCompiled panics when the regex is nil
func TestCustomCompiled_NilRegex(t *testing.T) {
	require.Panics(t, func() {
		sanitize.CustomCompiled("panic", nil)
	})
}

// TestCustom_InvalidRegexPanics verifies that Custom panics when given
// an invalid regular expression pattern.
func TestCustom_InvalidRegexPanics(t *testing.T) {
	require.Panics(t, func() {
		sanitize.Custom("invalid", "(")
	})
}

// TestCustom_UnicodePattern ensures Unicode characters are preserved when
// the regex allows them.
func TestCustom_UnicodePattern(t *testing.T) {
	//nolint:gosmopolitan // test includes Unicode characters
	output := sanitize.Custom("Héllo 世界!123", `[^\p{L}\s]`)
	//nolint:gosmopolitan // test includes Unicode characters
	assert.Equal(t, "Héllo 世界", output)
}

// TestCustom_OverlappingMatches validates behavior when the pattern could
// match overlapping segments.
func TestCustom_OverlappingMatches(t *testing.T) {
	output := sanitize.Custom("ababa", "aba")
	assert.Equal(t, "ba", output)
}

// TestCustomCompiled_UnicodePattern ensures Unicode patterns work with
// precompiled regular expressions.
func TestCustomCompiled_UnicodePattern(t *testing.T) {
	re := regexp.MustCompile(`[^\p{L}\s]`)
	//nolint:gosmopolitan // test includes Unicode characters
	output := sanitize.CustomCompiled("Héllo 世界!123", re)
	//nolint:gosmopolitan // test includes Unicode characters
	assert.Equal(t, "Héllo 世界", output)
}

// TestCustomCompiled_OverlappingMatches verifies that overlapping matches are
// handled as expected when using a precompiled regex.
func TestCustomCompiled_OverlappingMatches(t *testing.T) {
	re := regexp.MustCompile("aba")
	output := sanitize.CustomCompiled("ababa", re)
	assert.Equal(t, "ba", output)
}

// TestDecimal tests the decimal sanitize method
func TestDecimal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"basic 1", " String: 1.23 ", "1.23"},
		{"basic 2", " String: 001.2300 ", "001.2300"},
		{"basic 3", "  $-1.034234  Price", "-1.034234"},
		{"basic 4", "  $-1%.034234e  Price", "-1.034234"},
		{"basic 5", "/n<<  $-1.034234  >>/n", "-1.034234"},

		// Edge cases
		{"empty string", "", ""},
		{"letters only", "abc", ""},
		{"plus sign", "+100.50", "100.50"},
		{"multiple decimals", "1.2.3", "1.2.3"},
		{"embedded minus", "1-2-3", "1-2-3"},
		{"scientific notation", "1e-3", "1-3"},
		{"comma separated", "1,234.56", "1234.56"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Decimal(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestDomain tests the domain sanitize method
func TestDomain(t *testing.T) {

	t.Run("valid cases", func(t *testing.T) {

		var tests = []struct {
			name         string
			input        string
			expected     string
			preserveCase bool
			removeWww    bool
		}{
			{
				"no domain name",
				"",
				"",
				true,
				true,
			},
			{
				"remove leading http",
				"http://IAmaDomain.com",
				"IAmaDomain.com",
				true,
				false,
			},
			{
				"remove leading http and lowercase",
				"http://IAmaDomain.com",
				"iamadomain.com",
				false,
				false,
			},
			{
				"full url with params",
				"https://IAmaDomain.com/?this=that#plusThis",
				"iamadomain.com",
				false,
				false,
			},
			{
				"full url with params, remove www",
				"https://www.IAmaDomain.com/?this=that#plusThis",
				"iamadomain.com",
				false,
				true,
			},
			{
				"full url with params, leave www",
				"https://www.IAmaDomain.com/?this=that#plusThis",
				"www.iamadomain.com",
				false,
				false,
			},
			{
				"caps domain, remove www",
				"WWW.DOMAIN.COM",
				"domain.com",
				false,
				true,
			},
			{
				"mixed caps domain, remove www",
				"WwW.DOMAIN.COM",
				"domain.com",
				false,
				true,
			},
			{
				"mixed caps domain, remove www",
				"wwW.DOMAIN.COM",
				"DOMAIN.COM",
				true,
				true,
			},
			{
				"domain with tabs and spaces",
				`		domain.com`,
				"domain.com",
				false,
				true,
			},
			{
				"invalid unicode in host",
				"https://examplé.com",
				"exampl.com",
				false,
				false,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				output, err := sanitize.Domain(test.input, test.preserveCase, test.removeWww)
				require.NoError(t, err)
				assert.Equal(t, test.expected, output)
			})
		}
	})

	t.Run("invalid cases", func(t *testing.T) {

		var tests = []struct {
			name         string
			input        string
			expected     string
			preserveCase bool
			removeWww    bool
		}{
			{
				"spaces in domain",
				"http://www.I am a domain.com",
				"http://www.I am a domain.com",
				true,
				true,
			},
			{
				"symbol in domain",
				"!I_am a domain.com",
				"http://!I_am a domain.com",
				true,
				true,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				output, err := sanitize.Domain(test.input, test.preserveCase, test.removeWww)
				require.Error(t, err)
				assert.Equal(t, test.expected, output)
			})
		}
	})
}

// TestEmail tests the email sanitize method
func TestEmail(t *testing.T) {
	var tests = []struct {
		name         string
		input        string
		expected     string
		preserveCase bool
	}{
		{"basic_1", "mailto:testME@GmAil.com", "testme@gmail.com", false},
		{"basic_2", "test_ME@GmAil.com", "test_me@gmail.com", false},
		{"basic_3", "test-ME@GmAil.com", "test-me@gmail.com", false},
		{"basic_4", "test.ME@GmAil.com", "test.me@gmail.com", false},
		{"basic_5", " test_ME @GmAil.com ", "test_me@gmail.com", false},
		{"basic_6", " <<test_ME @GmAil.com!>> ", "test_me@gmail.com", false},
		{"basic_7", " test_ME+2@GmAil.com ", "test_me+2@gmail.com", false},
		{"basic_8", " test_ME+2@GmAil.com ", "test_ME+2@GmAil.com", true},

		// Additional edge cases
		{"empty string", "", "", false},
		{"spaces only", "   ", "", false},
		{"symbols only", "!@#$%^&*()", "@", false},
		{"invalid email format", "not-an-email", "not-an-email", false},
		{"multiple @ symbols", "test@@example.com", "test@@example.com", false},
		{"unicode in local part", "tést@exámple.com", "tst@exmple.com", false},
		{"unicode in domain", "test@exámple.com", "test@exmple.com", false},
		{"email with subdomain", "user@mail.example.com", "user@mail.example.com", false},
		{"email with numbers", "user123@123mail.com", "user123@123mail.com", false},
		{"email with dash and underscore", "user-name_test@domain-name.com", "user-name_test@domain-name.com", false},
		{"email with plus and dot", "user.name+tag@domain.com", "user.name+tag@domain.com", false},
		{"email with leading/trailing spaces", "  user@domain.com  ", "user@domain.com", false},
		{"email with mixed case and preserve", "Test@Domain.COM", "Test@Domain.COM", true},
		{"email with mixed case and no preserve", "Test@Domain.COM", "test@domain.com", false},
		{"email with mailto and preserve", "MailTo:Test@Domain.COM", "Test@Domain.COM", true},
		{"email with special chars in local", "user!#$%&'*+/=?^_`{|}~@domain.com", "user+_@domain.com", false},
		{"email with special chars in domain", "user@do!main.com", "user@domain.com", false},
		{"email with multiple dots", "user..name@domain.com", "user..name@domain.com", false},
		{"email with tab and newline", "user@domain.com\t\n", "user@domain.com", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Email(test.input, test.preserveCase)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestFirstToUpper tests the first to upper method
func TestFirstToUpper(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"", "thisworks", "Thisworks"},
		{"", "Thisworks", "Thisworks"},
		{"", "this", "This"},
		{"", "t", "T"},
		{"", "tt", "Tt"},
		{"", "", ""}, // Edge case for empty string

		// Additional edge cases:
		{"single space: ' '", " ", " "},
		{"multiple spaces: '  '", "  ", "  "},
		{"tab character: '\t'", "\t", "\t"},
		{"newline character: '\n'", "\n", "\n"},
		{"starts with number: '123abc'", "123abc", "123abc"},
		{"starts with symbol: '!@#'", "!@#", "!@#"},
		{"German sharp S: 'ßeta' (uppercases to 'SS')", "ßeta", "ßeta"},
		{"accented character: 'éclair' (should become 'Éclair')", "éclair", "Éclair"},
		{"Greek capital letter: 'Σigma' (should remain unchanged)", "Σigma", "Σigma"},
		{"Spanish n-tilde: 'ñandú' (should become 'Ñandú')", "ñandú", "Ñandú"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.FirstToUpper(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestFormalName tests the FormalName sanitize method
func TestFormalName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"basic 1", "Mark Mc'Cuban-Host", "Mark Mc'Cuban-Host"},
		{"basic 2", "Mark Mc'Cuban-Host the SR.", "Mark Mc'Cuban-Host the SR."},
		{"basic 3", "Mark Mc'Cuban-Host the Second.", "Mark Mc'Cuban-Host the Second."},
		{"basic 4", "Johnny Apple.Seed, Martin", "Johnny Apple.Seed, Martin"},
		{"basic 5", "Does #Not Work!", "Does Not Work"},

		// Edge cases
		{"empty string", "", ""},
		{"accented characters", "José María", "Jos Mara"},
		{"underscores", "Name_With_Underscore", "NameWithUnderscore"},
		{"digits", "John Doe 3rd", "John Doe 3rd"},
		{"newline", "John\nDoe", "John\nDoe"},
		{"leading spaces", "  John", "  John"},
		{"apostrophe and hyphen", "O'Leary-Brown", "O'Leary-Brown"},
		{"prefix d'", "d'Artagnan", "d'Artagnan"},
		{"curly apostrophe", "D’Angelo", "DAngelo"},
		{"multiple spaces", "Van  der  Meer", "Van  der  Meer"},
		{"accented surname", "Émilie du Châtelet", "milie du Chtelet"},
		{"foreign letters", "Björk Guðmundsdóttir", "Bjrk Gumundsdttir"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.FormalName(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestHTML tests the HTML sanitize method
func TestHTML(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"test_1", "<b>This works?</b>", "This works?"},
		{"test_2", "<html><b>This works?</b><i></i></br></html>", "This works?"},
		{"test_3", "<html><b class='test'>This works?</b><i></i></br></html>", "This works?"},
		{"nested tags", "Hello <div>world <span>!</span></div>", "Hello world !"},
		{"unclosed tag", "<div>test", "test"},
		{"partial closing", "<div>test</div", "test</div"},
		{"html comments", "<!-- comment -->text", "text"},
		{"script tag remains", "<script>alert('x')</script>", "alert('x')"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.HTML(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestIPAddress tests the ip address sanitize method
func TestIPAddress(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"basic_1", "192.168.3.6", "192.168.3.6"},
		{"basic_2", "255.255.255.255", "255.255.255.255"},
		{"basic_3", "304.255.255.255", ""},
		{"basic_4", "fail", ""},
		{"basic_5", "192-123-122-123", ""},
		{"basic_6", "2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f", "2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f"},
		{"basic_7", "2602:305:bceb:1bd0:44ef:2:2:2", "2602:305:bceb:1bd0:44ef:2:2:2"},
		{"basic_8", "2:2:2:2:2:2:2:2", "2:2:2:2:2:2:2:2"},
		{"basic_9", "192.2", ""},
		{"basic_10", "192.2!", ""},
		{"basic_11", "IP: 192.168.0.1 ", ""},
		{"basic_12", " 192.168.0.1 ", "192.168.0.1"},
		{"basic_13", "  ##!192.168.0.1!##  ", "192.168.0.1"},
		{"basic_14", `		192.168.1.1`, "192.168.1.1"},
		{"basic_15", `2001:0db8:85a3:0000:0000:8a2e:0370:7334`, "2001:db8:85a3::8a2e:370:7334"}, // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{"basic_16", `2001:0db8::0001:0000`, "2001:db8::1:0"},                                   // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{"basic_17", `2001:db8:0:0:1:0:0:1`, "2001:db8::1:0:0:1"},                               // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{"basic_18", `2001:db8:0000:1:1:1:1:1`, "2001:db8:0:1:1:1:1:1"},                         // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{"basic_19", `0:0:0:0:0:0:0:1`, "::1"},                                                  // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{"basic_20", `0:0:0:0:0:0:0:0`, "::"},                                                   // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address

		// Additional edge cases
		{"empty string", "", ""},
		{"spaces only", "   ", ""},
		{"symbols only", "!@#$%^&*()", ""},
		{"letters only", "abcdef", ""},
		{"ipv4 with trailing dot", "192.168.1.1.", ""},
		{"ipv4 with internal spaces", "192. 168. 1. 1", "192.168.1.1"},
		{"ipv4 with tabs and newlines", "\t192.168.1.1\n", "192.168.1.1"},
		{"ipv6 with uppercase", "2001:DB8:0:0:8:800:200C:417A", "2001:db8::8:800:200c:417a"},
		{"ipv6 with brackets", "[2001:db8::1]", "2001:db8::1"},
		{"ipv6 with zone index", "fe80::1%lo0", ""},
		{"ipv6 with internal spaces", "2001: db8:: 1", "2001:db8::1"},
		{"ipv6 with tabs and newlines", "\t2001:db8::1\n", "2001:db8::1"},
		{"ipv4-mapped ipv6", "::ffff:192.0.2.128", "192.0.2.128"},
		{"ipv6 with embedded ipv4", "::ffff:192.168.1.1", "192.168.1.1"},
		{"ipv6 with all zeros", "::", "::"},
		{"ipv6 loopback", "::1", "::1"},
		{"ipv4 loopback", "127.0.0.1", "127.0.0.1"},
		{"ipv4 broadcast", "255.255.255.255", "255.255.255.255"},
		{"ipv4 with port", "192.168.1.1:8080", ""},
		{"ipv6 with port", "[2001:db8::1]:443", "2001:db8::1:443"},
		{"ipv4 with subnet", "192.168.1.1/24", "192.168.1.124"},
		{"ipv6 with subnet", "2001:db8::1/64", "2001:db8::164"},
		{"ipv4 with prefix text", "IP:192.168.1.1", ""},
		{"ipv6 with prefix text", "IPv6:2001:db8::1", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.IPAddress(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestNumeric tests the numeric sanitize method
func TestNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"basic 1", " > Test This String-!1234", "1234"},
		{"basic 2", " $1.00 Price!", "100"},

		// Edge cases
		{"empty string", "", ""},
		{"letters only", "abcd", ""},
		{"negative decimal", "-123.45", "12345"},
		{"phone format", "(123) 456-7890", "1234567890"},
		{"hex prefix", "0xFF 55", "055"},
		{"spaces only", "   ", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Numeric(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestPathName tests the PathName sanitize method
func TestPathName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"basic 1", "My BadPath (10)", "MyBadPath10"},
		{"basic 2", "My BadPath (10)[]()!$", "MyBadPath10"},
		{"basic 3", "My_Folder-Path-123_TEST", "My_Folder-Path-123_TEST"},

		// Edge cases
		{"empty string", "", ""},
		{"file extension", "myfile.txt", "myfiletxt"},
		{"windows path", "C:\\temp\\file.txt", "Ctempfiletxt"},
		{"unicode chars", "naïve.txt", "navetxt"},
		{"spaces", "dir name/file", "dirnamefile"},
		{"valid symbols", "filename-123_ABC", "filename-123_ABC"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.PathName(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestPunctuation tests the punctuation sanitize method
func TestPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"basic 1", "Mark Mc'Cuban-Host", "Mark Mc'Cuban-Host"},
		{"basic 2", "Johnny Apple.Seed, Martin", "Johnny Apple.Seed, Martin"},
		{"basic 3", "Does #Not Work!", "Does #Not Work!"},
		{"basic 4", "Does #Not Work!?", "Does #Not Work!?"},
		{"basic 5", "Does #Not Work! & this", "Does #Not Work! & this"},
		{"basic 6", `[@"Does" 'this' work?@]this`, `"Does" 'this' work?this`},
		{"basic 7", "Does, 123^* Not & Work!?", "Does, 123 Not & Work!?"},

		// Edge cases
		{"empty string", "", ""},
		{"spaces only", "   ", "   "},
		{"tabs and newlines", "line1\nline2\tend", "line1\nline2\tend"},
		{"disallowed punctuation", "Hello; world: [test] {case}", "Hello world test case"},
		{"unicode punctuation", "¡Hola señor!", "Hola señor!"},
		{"accents kept", "Café & crème brûlée?", "Café & crème brûlée?"},
		{"underscore and plus", "foo_bar+baz", "foobarbaz"},
		{"parentheses", "Need (something); else: yes?", "Need something else yes?"},
		{"smart quotes", "She said “Hello”", "She said Hello"},
		{"dash variants", "This—is—dash", "Thisisdash"},
		{"numbers with punctuation", "Version 2.0.1, (build #1234)", "Version 2.0.1, build #1234"},
		{"emoji", "Smile 😊, please!", "Smile , please!"},
		{"mixed allowed", `He said: "Go!"`, `He said "Go!"`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Punctuation(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestScientificNotation tests the scientific notation sanitize method
func TestScientificNotation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Standard cases
		{"simple float", " String: 1.23 ", "1.23"},
		{"with exponent", " String: 1.23e-3 ", "1.23e-3"},
		{"negative exponent", " String: -1.23e-3 ", "-1.23e-3"},
		{"leading zeros", " String: 001.2300 ", "001.2300"},
		{"prefixed with dollar and word", "  $-1.034234  word", "-1.034234"},
		{"prefixed with symbols and word", "  $-1%.034234e  word", "-1.034234e"},
		{"wrapped in symbols", "/n<<  $-1.034234  >>/n", "-1.034234"},

		// Edge cases
		{"empty string", "", ""},
		{"letters only", "abcde", "e"},
		{"uppercase exponent", "1.2E+3", "1.2E+3"},
		{"trailing plus", "1.0e+3+", "1.0e+3+"},
		{"multiple exponents", "1e2e3", "1e2e3"},
		{"comma separated", "1,234.56e7", "1234.56e7"},
		{"embedded minus", "1-2.3e4", "1-2.3e4"},
		{"arabic digits", "١٢٣.٤٥e٦", "١٢٣.٤٥e٦"},
		{"whitespace and newline", "1.2e3\n4.5e6", "1.2e34.5e6"},
		{"multiple decimals", "1.2.3e4", "1.2.3e4"},
		{"signs only", "+-", "+-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.ScientificNotation(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestScripts tests the script removal
func TestScripts(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"test_1", "this <script>$('#something').hide()</script>", "this "},
		{"test_2", "this <script type='text/javascript'>$('#something').hide()</script>", "this "},
		{"test_3", `this <script type="text/javascript" class="something">$('#something').hide();</script>`, "this "},
		{"test_4", `this <iframe width="50" class="something"></iframe>`, "this "},
		{"test_5", `this <embed width="50" class="something"></embed>`, "this "},
		{"test_6", `this <object width="50" class="something"></object>`, "this "},
		{"multiple scripts", "pre<script>1</script>mid<script>2</script>post", "prepost"},
		{"mismatched tags", "<script>one</iframe>two</script>", ""},
		{"uppercase script", "this <SCRIPT>evil()</SCRIPT> works", "this  works"},
		{"unclosed script", "<script>oops", "<script>oops"},
		{"closing only", "oops</script>", "oops</script>"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Scripts(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestSingleLine tests the SingleLine method
func TestSingleLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Regular cases
		{"basic multiline", "Mark\nMc'Cuban-Host", "Mark Mc'Cuban-Host"},
		{"multiline with extra line", "Mark\nMc'Cuban-Host\nsomething else", "Mark Mc'Cuban-Host something else"},
		{"leading tab with newlines", "\tMark\nMc'Cuban-Host\nsomething else", " Mark Mc'Cuban-Host something else"},

		// Edge cases
		{"empty string", "", ""},
		{"only whitespace", "\n\r\t\v\f", "     "},
		{"mixed whitespace", "Line1\r\nLine2\tLine3\vLine4\f", "Line1  Line2 Line3 Line4 "},
		{"leading and trailing", "\nStart\t\n", " Start  "},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.SingleLine(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestTime tests the time sanitize method
func TestTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"basic t00:00d", "t00:00d -EST", "00:00"},
		{"basic t00:00:00d", "t00:00:00d -EST", "00:00:00"},
		{"embedded time", "SOMETHING t00:00:00d -EST DAY", "00:00:00"},

		// Edge cases
		{"empty string", "", ""},
		{"nonsense string", "abc", ""},
		{"time with AM/PM", "10:20PM", "10:20"},
		{"negative time prefix", "-10:20", "10:20"},
		{"subsecond time", "12:34:56.789", "12:34:56789"},
		{"whitespace in time", "10\n:20\t:30", "10:20:30"},
		{"hyphen separated", "12-34-56", "123456"},
		{"only colons", "::", "::"},
		{"trailing colon", "12:34:", "12:34:"},
		{"unicode digits", "１２：３４", "１２３４"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Time(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestURI tests the URI sanitize method
func TestURI(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"basic 1", "Test?=what! &this=that", "Test?=what&this=that"},
		{"basic 2", "Test?=what! &this=/that/!()*^", "Test?=what&this=/that/"},
		{"basic 3", "/This/Works/?that=123&this#page10%", "/This/Works/?that=123&this#page10%"},

		// Edge cases
		{"encoded spaces", "path%20with%20space", "path%20with%20space"},
		{"remove colon", "path:to/resource", "pathto/resource"},
		{"unicode characters", "/世界/привет", "/世界/привет"}, //nolint:gosmopolitan // Unicode characters are valid in URIs
		{"plus sign in query", "/query?name=foo+bar", "/query?name=foobar"},
		{"mixed invalid characters", "/path/../to/;evil?x=1^&y=2", "/path//to/evil?x=1&y=2"},
		{"trim spaces", "  /something ", "/something"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.URI(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestURL tests the URL sanitize method
func TestURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"remove spaces", "Test?=what! &this=that#works", "Test?=what&this=that#works"},
		{"no dollar signs", "/this/test?param$", "/this/test?param"},
		{"using at sign", "https://medium.com/@username/some-title-that-is-a-article", "https://medium.com/@username/some-title-that-is-a-article"},
		{"removing symbols", "https://domain.com/this/test?param$!@()[]{}'<>", "https://domain.com/this/test?param@"},
		{"params and anchors", "https://domain.com/this/test?this=value&another=123%#page", "https://domain.com/this/test?this=value&another=123%#page"},
		{"allow commas", "https://domain.com/this/test,this,value", "https://domain.com/this/test,this,value"},

		// Edge cases
		{"with port", "https://example.com:8080/path", "https://example.com:8080/path"},
		{"ipv6 address", "https://[2001:db8::1]/path", "https://2001:db8::1/path"},
		{"plus sign in query", "https://example.com?q=foo+bar", "https://example.com?q=foobar"},
		{"trim spaces", " https://example.com/test ", "https://example.com/test"},
		{"file url path", "file:///C:/Program Files/Test", "file:///C:/ProgramFiles/Test"},
		{"with user info", "https://user:pass@example.com/", "https://user:pass@example.com/"},
		{"fragment invalid characters", "https://example.com/path#frag!", "https://example.com/path#frag"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.URL(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestXML tests the XML sanitize method
func TestXML(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"test_1", `<?xml version="1.0" encoding="UTF-8"?><note>Something</note>`, "Something"},
		{"test_2", `<body>This works?</body><title>Something</title>`, "This works?Something"},
		{"nested tags", `<root><child>data</child><child2/></root>`, "data"},
		{"attributes", `text <tag attr='1'>value</tag> more`, "text value more"},
		{"unclosed tag", `<tag>unclosed`, "unclosed"},
		{"xml header and comment", `<?xml version='1.0'?><!--comment--><a>1</a>`, "1"},
		{"cdata removed", `<a><![CDATA[test]]></a>`, ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.XML(test.input)
			assert.Equal(t, test.expected, output)
		})

	}
}

// TestXSS tests the XSS sanitize method
func TestXSS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Common script injection vectors
		{"Remove <script", "<script", ""},
		{"Remove script>", "script>", ""},
		{"Remove eval(", "eval(", ""},
		{"Remove eval&#40;", "eval&#40;", ""},
		{"Remove javascript:", "javascript:", ""},
		{"Remove javascript&#58;", "javascript&#58;", ""},
		{"Remove fromCharCode", "fromCharCode", ""},
		{"Remove &#62;", "&#62;", ""},
		{"Remove &#60;", "&#60;", ""},
		{"Remove &lt;", "&lt;", ""},
		{"Remove &rt;", "&rt;", ""},

		// Inline event handlers
		{"Remove onclick=", "onclick=", ""},
		{"Remove onerror=", "onerror=", ""},
		{"Remove onload=", "onload=", ""},
		{"Remove onmouseover=", "onmouseover=", ""},
		{"Remove onfocus=", "onfocus=", ""},
		{"Remove onblur=", "onblur=", ""},
		{"Remove ondblclick=", "ondblclick=", ""},
		{"Remove onkeydown=", "onkeydown=", ""},
		{"Remove onkeyup=", "onkeyup=", ""},
		{"Remove onkeypress=", "onkeypress=", ""},

		// Potential CSS/Style-based attacks
		{"Remove expression(", "expression(", ""},

		// Potentially malicious protocols
		{"Remove data:", "data:", ""},

		// Dangerous objects/functions
		{"Remove document.cookie", "document.cookie", ""},
		{"Remove document.write", "document.write", ""},
		{"Remove window.location", "window.location", ""},

		// Additional cases
		{"Multiple patterns", "<script>eval(javascript:alert(1))</script>", ">alert(1))</"},
		{"Pattern in text", "Hello<script>alert(1)</script>World", "Hello>alert(1)</World"},
		{"Mixed case script", "<ScRiPt>alert(1)</sCrIpT>", "<ScRiPt>alert(1)</sCrIpT>"},
		{"HTML entity encoded", "&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;", "&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;"},
		{"Whitespace in tag", "<scr ipt>alert(1)</scr ipt>", "<scr ipt>alert(1)</scr ipt>"},
		{"Inline event handler", "<img src=x onerror=alert(1)>", "<img src=x alert(1)>"},
		{"Obfuscated event handler", "<img src=x oNclIck=alert(1)>", "<img src=x oNclIck=alert(1)>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := sanitize.XSS(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}
