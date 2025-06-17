package sanitize_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/mrz1836/go-sanitize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAlpha tests the alpha sanitize method
func TestAlpha_Basic(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
		typeCase bool
	}{
		{"regular string", "Test This String-!123", "TestThisString", false},
		{"various symbols", `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`, "Symbols", false},
		{"carriage returns", "\nThis\nThat", "ThisThat", false},
		{"quotes and ticks", "“This is a quote with tick`s … ” ☺ ", "Thisisaquotewithticks", false},
		{"spaces", "Test This String-!123", "Test This String", true},
		{"symbols and spaces", `~!@#$%^&*()-_Symbols=+[{]};:'"<>,./?`, "Symbols", true},
		{"quotes and spaces", "“This is a quote with tick`s … ” ☺ ", "This is a quote with ticks    ", true},
		{"carriage returns", "\nThis\nThat", `ThisThat`, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.Alpha(test.input, test.typeCase)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestAlphaEdgeCases tests the alpha sanitize method with edge cases
func TestAlpha_EdgeCases(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
		typeCase bool
	}{
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

// BenchmarkAlphaNoSpaces benchmarks the Alpha method
func BenchmarkAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Alpha("This is the test string.", false)
	}
}

// BenchmarkAlpha_WithSpaces benchmarks the Alpha method
func BenchmarkAlpha_WithSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Alpha("This is the test string.", true)
	}
}

// ExampleAlpha example using Alpha() and no spaces flag
func ExampleAlpha() {
	fmt.Println(sanitize.Alpha("Example String!", false))
	// Output: ExampleString
}

// ExampleAlpha_withSpaces example using Alpha with a space flag
func ExampleAlpha_withSpaces() {
	fmt.Println(sanitize.Alpha("Example String!", true))
	// Output: Example String
}

// TestAlphaNumeric tests the alphanumeric sanitize method
func TestAlphaNumeric_Basic(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
		typeCase bool
	}{
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.AlphaNumeric(test.input, test.typeCase)
			assert.Equal(t, test.expected, output)
		})
	}
}

// TestAlphaNumeric_EdgeCases tests AlphaNumeric with additional edge cases
func TestAlphaNumeric_EdgeCases(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
		typeCase bool
	}{
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

// BenchmarkAlphaNumeric benchmarks the AlphaNumeric method
func BenchmarkAlphaNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.AlphaNumeric("This is the test string 12345.", false)
	}
}

// BenchmarkAlphaNumeric_WithSpaces benchmarks the AlphaNumeric method
func BenchmarkAlphaNumeric_WithSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.AlphaNumeric("This is the test string 12345.", true)
	}
}

// ExampleAlphaNumeric example using AlphaNumeric() with no spaces
func ExampleAlphaNumeric() {
	fmt.Println(sanitize.AlphaNumeric("Example String 2!", false))
	// Output: ExampleString2
}

// ExampleAlphaNumeric_withSpaces example using AlphaNumeric() with spaces
func ExampleAlphaNumeric_withSpaces() {
	fmt.Println(sanitize.AlphaNumeric("Example String 2!", true))
	// Output: Example String 2
}

// TestBitcoinAddress will test all permutations
func TestBitcoinAddress_Basic(t *testing.T) {

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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.BitcoinAddress(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkBitcoinAddress benchmarks the BitcoinAddress method
func BenchmarkBitcoinAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.BitcoinAddress("1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs")
	}
}

// ExampleBitcoinAddress example using BitcoinAddress()
func ExampleBitcoinAddress() {
	fmt.Println(sanitize.BitcoinAddress(":1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!"))
	// Output: 1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs
}

// TestBitcoinCashAddress will test all permutations of using BitcoinCashAddress()
func TestBitcoinCashAddress_Basic(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"remove symbols", "$#:qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz!", "qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz"},
		{"remove spaces", " $#:qze7yy2 au5vuznvn8lzj5y0j5t066 vhs75e3m0eptz! ", "qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz"},

		// No letters o, b, i, or number 1
		{"remove ignored characters", "pqbq3728yw0y47sOqn6l2na30mcw6zm78idzq5ucqzc371", "pqq3728yw0y47sqn6l2na30mcw6zm78dzq5ucqzc37"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.BitcoinCashAddress(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkBitcoinCashAddress benchmarks the BitcoinCashAddress() method
func BenchmarkBitcoinCashAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.BitcoinCashAddress("qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz")
	}
}

// ExampleBitcoinCashAddress example using BitcoinCashAddress() `cashaddr`
func ExampleBitcoinCashAddress() {
	fmt.Println(sanitize.BitcoinAddress("qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz!"))
	// Output: qze7yy2au5vuznvn8zj5yj5t66vhs75e3meptz
}

// TestCustom tests the custom sanitize method
func TestCustom_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
		regex    string
	}{
		{"ThisWorks123!", "ThisWorks123", `[^a-zA-Z0-9]`},
		{"ThisWorks1.23!", "1.23", `[^0-9.-]`},
		{"ThisWorks1.23!", "ThisWorks123", `[^0-9a-zA-Z]`},
	}

	for _, test := range tests {
		output := sanitize.Custom(test.input, test.regex)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkCustom benchmarks the Custom method
func BenchmarkCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Custom("This is the test string 12345.", `[^a-zA-Z0-9]`)
	}
}

// ExampleCustom example using Custom() using an alpha regex
func ExampleCustom() {
	fmt.Println(sanitize.Custom("Example String 2!", `[^a-zA-Z]`))
	// Output: ExampleString
}

// ExampleCustom_numeric example using Custom() using a numeric regex
func ExampleCustom_numeric() {
	fmt.Println(sanitize.Custom("Example String 2!", `[^0-9]`))
	// Output: 2
}

// TestCustomCompiled verifies CustomCompiled using a precompiled regex
func TestCustomCompiled_Basic(t *testing.T) {
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

// BenchmarkCustomCompiled benchmarks the CustomCompiled method
func BenchmarkCustomCompiled(b *testing.B) {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	for i := 0; i < b.N; i++ {
		_ = sanitize.CustomCompiled("This is the test string 12345.", re)
	}
}

// ExampleCustomCompiled example using CustomCompiled with an alpha regex
func ExampleCustomCompiled() {
	re := regexp.MustCompile(`[^a-zA-Z]`)
	fmt.Println(sanitize.CustomCompiled("Example String 2!", re))
	// Output: ExampleString
}

// TestDecimal tests the decimal sanitize method
func TestDecimal_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{" String: 1.23 ", "1.23"},
		{" String: 001.2300 ", "001.2300"},
		{"  $-1.034234  Price", "-1.034234"},
		{"  $-1%.034234e  Price", "-1.034234"},
		{"/n<<  $-1.034234  >>/n", "-1.034234"},
	}

	for _, test := range tests {
		output := sanitize.Decimal(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkDecimal benchmarks the Decimal method
func BenchmarkDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Decimal("String: -123.12345")
	}
}

// ExampleDecimal example using Decimal() for a positive number
func ExampleDecimal() {
	fmt.Println(sanitize.Decimal("$ 99.99!"))
	// Output: 99.99
}

// ExampleDecimal_negative example using Decimal() for a negative number
func ExampleDecimal_negative() {
	fmt.Println(sanitize.Decimal("$ -99.99!"))
	// Output: -99.99
}

// TestDomain tests the domain sanitize method
func TestDomain_Basic(t *testing.T) {

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
				"http://www.IAmaDomain.com/?this=that#plusThis",
				"iamadomain.com",
				false,
				true,
			},
			{
				"full url with params, leave www",
				"http://www.IAmaDomain.com/?this=that#plusThis",
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

// BenchmarkDomain benchmarks the Domain method
func BenchmarkDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sanitize.Domain("https://Example.COM/?param=value", false, false)
	}
}

// BenchmarkDomain_PreserveCase benchmarks the Domain method
func BenchmarkDomain_PreserveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sanitize.Domain("https://Example.COM/?param=value", true, false)
	}
}

// BenchmarkDomain_RemoveWww benchmarks the Domain method
func BenchmarkDomain_RemoveWww(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sanitize.Domain("https://Example.COM/?param=value", false, true)
	}
}

// ExampleDomain example using Domain()
func ExampleDomain() {
	fmt.Println(sanitize.Domain("https://www.Example.COM/?param=value", false, false))
	// Output: www.example.com <nil>
}

// ExampleDomain_preserveCase example using Domain() and preserving the case
func ExampleDomain_preserveCase() {
	fmt.Println(sanitize.Domain("https://www.Example.COM/?param=value", true, false))
	// Output: www.Example.COM <nil>
}

// ExampleDomain_removeWww example using Domain() and removing the www subdomain
func ExampleDomain_removeWww() {
	fmt.Println(sanitize.Domain("https://www.Example.COM/?param=value", false, true))
	// Output: example.com <nil>
}

// TestEmail tests the email sanitize method
func TestEmail_Basic(t *testing.T) {

	var tests = []struct {
		input        string
		expected     string
		preserveCase bool
	}{
		{"mailto:testME@GmAil.com", "testme@gmail.com", false},
		{"test_ME@GmAil.com", "test_me@gmail.com", false},
		{"test-ME@GmAil.com", "test-me@gmail.com", false},
		{"test.ME@GmAil.com", "test.me@gmail.com", false},
		{" test_ME @GmAil.com ", "test_me@gmail.com", false},
		{" <<test_ME @GmAil.com!>> ", "test_me@gmail.com", false},
		{" test_ME+2@GmAil.com ", "test_me+2@gmail.com", false},
		{" test_ME+2@GmAil.com ", "test_ME+2@GmAil.com", true},
	}

	for _, test := range tests {
		output := sanitize.Email(test.input, test.preserveCase)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkEmail benchmarks the Email method
func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Email("mailto:Person@Example.COM ", false)
	}
}

// BenchmarkEmail_PreserveCase benchmarks the Email method
func BenchmarkEmail_PreserveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Email("mailto:Person@Example.COM ", true)
	}
}

// ExampleEmail example using Email()
func ExampleEmail() {
	fmt.Println(sanitize.Email("mailto:Person@Example.COM", false))
	// Output: person@example.com
}

// ExampleEmail_preserveCase example using Email() and preserving the case
func ExampleEmail_preserveCase() {
	fmt.Println(sanitize.Email("mailto:Person@Example.COM", true))
	// Output: Person@Example.COM
}

// TestFirstToUpper tests the first to upper method
func TestFirstToUpper_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"thisworks", "Thisworks"},
		{"Thisworks", "Thisworks"},
		{"this", "This"},
		{"t", "T"},
		{"tt", "Tt"},
		{"", ""}, // Edge case for empty string

		// Additional edge cases:
		{" ", " "},           // single space
		{"  ", "  "},         // multiple spaces
		{"\t", "\t"},         // tab character
		{"\n", "\n"},         // newline character
		{"123abc", "123abc"}, // starts with number
		{"!@#", "!@#"},       // starts with symbol
		{"ßeta", "ßeta"},     // German sharp S (will uppercase to "SS")
		{"éclair", "Éclair"}, // accented character
		{"Σigma", "Σigma"},   // Greek capital letter (should remain unchanged)
		{"ñandú", "Ñandú"},   // Spanish n-tilde
	}

	for _, test := range tests {
		output := sanitize.FirstToUpper(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkFirstToUpper benchmarks the FirstToUpper method
func BenchmarkFirstToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.FirstToUpper("make this upper")
	}
}

// ExampleFirstToUpper example using FirstToUpper()
func ExampleFirstToUpper() {
	fmt.Println(sanitize.FirstToUpper("this works"))
	// Output: This works
}

// TestFormalName tests the formal name method
func TestFormalName_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"Mark Mc'Cuban-Host", "Mark Mc'Cuban-Host"},
		{"Mark Mc'Cuban-Host the SR.", "Mark Mc'Cuban-Host the SR."},
		{"Mark Mc'Cuban-Host the Second.", "Mark Mc'Cuban-Host the Second."},
		{"Johnny Apple.Seed, Martin", "Johnny Apple.Seed, Martin"},
		{"Does #Not Work!", "Does Not Work"},
	}

	for _, test := range tests {
		output := sanitize.FormalName(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkFormalName benchmarks the FormalName method
func BenchmarkFormalName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.FormalName("John McDonald Jr.")
	}
}

// ExampleFormalName example using FormalName()
func ExampleFormalName() {
	fmt.Println(sanitize.FormalName("John McDonald Jr.!"))
	// Output: John McDonald Jr.
}

// TestHTML tests the HTML sanitize method
func TestHTML_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"<b>This works?</b>", "This works?"},
		{"<html><b>This works?</b><i></i></br></html>", "This works?"},
		{"<html><b class='test'>This works?</b><i></i></br></html>", "This works?"},
	}

	for _, test := range tests {
		output := sanitize.HTML(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkHTML benchmarks the HTML method
func BenchmarkHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.HTML("<html><b>Test This!</b></html>")
	}
}

// ExampleHTML example using HTML()
func ExampleHTML() {
	fmt.Println(sanitize.HTML("<body>This Works?</body>"))
	// Output: This Works?
}

// TestIPAddress tests the ip address sanitize method
func TestIPAddress_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"192.168.3.6", "192.168.3.6"},
		{"255.255.255.255", "255.255.255.255"},
		{"304.255.255.255", ""},
		{"fail", ""},
		{"192-123-122-123", ""},
		{"2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f", "2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f"},
		{"2602:305:bceb:1bd0:44ef:2:2:2", "2602:305:bceb:1bd0:44ef:2:2:2"},
		{"2:2:2:2:2:2:2:2", "2:2:2:2:2:2:2:2"},
		{"192.2", ""},
		{"192.2!", ""},
		{"IP: 192.168.0.1 ", ""},
		{" 192.168.0.1 ", "192.168.0.1"},
		{"  ##!192.168.0.1!##  ", "192.168.0.1"},
		{`		192.168.1.1`, "192.168.1.1"},
		{`2001:0db8:85a3:0000:0000:8a2e:0370:7334`, "2001:db8:85a3::8a2e:370:7334"}, // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{`2001:0db8::0001:0000`, "2001:db8::1:0"},                                   // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{`2001:db8:0:0:1:0:0:1`, "2001:db8::1:0:0:1"},                               // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{`2001:db8:0000:1:1:1:1:1`, "2001:db8:0:1:1:1:1:1"},                         // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{`0:0:0:0:0:0:0:1`, "::1"},                                                  // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
		{`0:0:0:0:0:0:0:0`, "::"},                                                   // Gets parsed and changes the display, see: https://en.wikipedia.org/wiki/IPv6_address
	}

	for _, test := range tests {
		output := sanitize.IPAddress(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkIPAddress benchmarks the IPAddress method
func BenchmarkIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.IPAddress(" 192.168.0.1 ")
	}
}

// BenchmarkIPAddress_V6 benchmarks the IPAddress method
func BenchmarkIPAddress_IPV6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.IPAddress(" 2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f ")
	}
}

// ExampleIPAddress example using IPAddress() for IPV4 address
func ExampleIPAddress() {
	fmt.Println(sanitize.IPAddress(" 192.168.0.1 "))
	// Output: 192.168.0.1
}

// ExampleIPAddress_ipv6 example using IPAddress() for IPV6 address
func ExampleIPAddress_ipv6() {
	fmt.Println(sanitize.IPAddress(" 2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f "))
	// Output: 2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f
}

// TestNumeric tests the numeric sanitize method
func TestNumeric_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{" > Test This String-!1234", "1234"},
		{" $1.00 Price!", "100"},
	}

	for _, test := range tests {
		output := sanitize.Numeric(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// TestNumeric_EdgeCases tests Numeric with additional edge cases
func TestNumeric_EdgeCases(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
	}{
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

// BenchmarkNumeric benchmarks the numeric method
func BenchmarkNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Numeric(" 192.168.0.1 ")
	}
}

// ExampleNumeric example using Numeric()
func ExampleNumeric() {
	fmt.Println(sanitize.Numeric("This:123 + 90!"))
	// Output: 12390
}

// TestPathName tests the path name sanitize method
func TestPathName_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"My BadPath (10)", "MyBadPath10"},
		{"My BadPath (10)[]()!$", "MyBadPath10"},
		{"My_Folder-Path-123_TEST", "My_Folder-Path-123_TEST"},
	}

	for _, test := range tests {
		output := sanitize.PathName(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkPathName benchmarks the PathName method
func BenchmarkPathName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.PathName("/This-Path-Name_Works-/")
	}
}

// ExampleNumeric example using PathName()
func ExamplePathName() {
	fmt.Println(sanitize.PathName("/This-Works_Now-123/!"))
	// Output: This-Works_Now-123
}

// TestPunctuation tests the punctuation method
func TestPunctuation_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"Mark Mc'Cuban-Host", "Mark Mc'Cuban-Host"},
		{"Johnny Apple.Seed, Martin", "Johnny Apple.Seed, Martin"},
		{"Does #Not Work!", "Does #Not Work!"},
		{"Does #Not Work!?", "Does #Not Work!?"},
		{"Does #Not Work! & this", "Does #Not Work! & this"},
		{`[@"Does" 'this' work?@]this`, `"Does" 'this' work?this`},
		{"Does, 123^* Not & Work!?", "Does, 123 Not & Work!?"},
	}

	for _, test := range tests {
		output := sanitize.Punctuation(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkPunctuation benchmarks the Punctuation method
func BenchmarkPunctuation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Punctuation("Does this work? They're doing it?")
	}
}

// ExamplePunctuation example using Punctuation()
func ExamplePunctuation() {
	fmt.Println(sanitize.Punctuation(`[@"Does" 'this' work?@] this too`))
	// Output: "Does" 'this' work? this too
}

// TestScientificNotation tests the scientific notation sanitize method
func TestScientificNotation_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{" String: 1.23 ", "1.23"},
		{" String: 1.23e-3 ", "1.23e-3"},
		{" String: -1.23e-3 ", "-1.23e-3"},
		{" String: 001.2300 ", "001.2300"},
		{"  $-1.034234  word", "-1.034234"},
		{"  $-1%.034234e  word", "-1.034234e"},
		{"/n<<  $-1.034234  >>/n", "-1.034234"},
	}

	for _, test := range tests {
		output := sanitize.ScientificNotation(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkDecimal benchmarks the ScientificNotation method
func BenchmarkScientificNotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.ScientificNotation("String: -1.096e-3")
	}
}

// ExampleScientificNotation example using ScientificNotation() for a positive number
func ExampleScientificNotation() {
	fmt.Println(sanitize.ScientificNotation("$ 1.096e-3!"))
	// Output: 1.096e-3
}

// TestScripts tests the script removal
func TestScripts_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"this <script>$('#something').hide()</script>", "this "},
		{"this <script type='text/javascript'>$('#something').hide()</script>", "this "},
		{`this <script type="text/javascript" class="something">$('#something').hide();</script>`, "this "},
		{`this <iframe width="50" class="something"></iframe>`, "this "},
		{`this <embed width="50" class="something"></embed>`, "this "},
		{`this <object width="50" class="something"></object>`, "this "},
	}

	for _, test := range tests {
		output := sanitize.Scripts(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkScripts benchmarks the Scripts method
func BenchmarkScripts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Scripts("<script>$(){ var remove='me'; }</script>")
	}
}

// ExampleScripts example using Scripts()
func ExampleScripts() {
	fmt.Println(sanitize.Scripts(`Does<script>This</script>Work?`))
	// Output: DoesWork?
}

// TestSingleLine test the single line sanitize method
func TestSingleLine_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{`Mark
Mc'Cuban-Host`, "Mark Mc'Cuban-Host"},
		{`Mark
Mc'Cuban-Host
something else`, "Mark Mc'Cuban-Host something else"},
		{`	Mark
Mc'Cuban-Host
something else`, " Mark Mc'Cuban-Host something else"},
	}

	for _, test := range tests {
		output := sanitize.SingleLine(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkSingleLine benchmarks the SingleLine method
func BenchmarkSingleLine(b *testing.B) {
	testString := `This line
That Line
Another Line`
	for i := 0; i < b.N; i++ {
		_ = sanitize.SingleLine(testString)
	}
}

// ExampleSingleLine example using SingleLine()
func ExampleSingleLine() {
	fmt.Println(sanitize.SingleLine(`Does
This
Work?`))
	// Output: Does This Work?
}

// TestTime tests the time sanitize method
func TestTime_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"t00:00d -EST", "00:00"},
		{"t00:00:00d -EST", "00:00:00"},
		{"SOMETHING t00:00:00d -EST DAY", "00:00:00"},
	}

	for _, test := range tests {
		output := sanitize.Time(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkTime benchmarks the Time method
func BenchmarkTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Time("Time is 05:10:23")
	}
}

// ExampleTime example using Time()
func ExampleTime() {
	fmt.Println(sanitize.Time(`Time 01:02:03!`))
	// Output: 01:02:03
}

// TestURI tests the URI sanitize method
func TestURI_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"Test?=what! &this=that", "Test?=what&this=that"},
		{"Test?=what! &this=/that/!()*^", "Test?=what&this=/that/"},
		{"/This/Works/?that=123&this#page10%", "/This/Works/?that=123&this#page10%"},
	}

	for _, test := range tests {
		output := sanitize.URI(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkURI benchmarks the URI method
func BenchmarkURI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.URI("/Test/This/Url/?param=value")
	}
}

// ExampleURI example using URI()
func ExampleURI() {
	fmt.Println(sanitize.URI("/This/Works?^No&this"))
	// Output: /This/Works?No&this
}

// TestURL tests the URL sanitize method
func TestURL_Basic(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"remove spaces", "Test?=what! &this=that#works", "Test?=what&this=that#works"},
		{"no dollar signs", "/this/test?param$", "/this/test?param"},
		{"using at sign", "https://medium.com/@username/some-title-that-is-a-article", "https://medium.com/@username/some-title-that-is-a-article"},
		{"removing symbols", "https://domain.com/this/test?param$!@()[]{}'<>", "https://domain.com/this/test?param@"},
		{"params and anchors", "https://domain.com/this/test?this=value&another=123%#page", "https://domain.com/this/test?this=value&another=123%#page"},
		{"allow commas", "https://domain.com/this/test,this,value", "https://domain.com/this/test,this,value"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sanitize.URL(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkURL benchmarks the URL method
func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.URL("/Test/This/Url/?param=value")
	}
}

// ExampleURL example using URL()
func ExampleURL() {
	fmt.Println(sanitize.URL("https://Example.com/This/Works?^No&this"))
	// Output: https://Example.com/This/Works?No&this
}

// TestXML tests the XML sanitize method
func TestXML_Basic(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{`<?xml version="1.0" encoding="UTF-8"?><note>Something</note>`, "Something"},
		{`<body>This works?</body><title>Something</title>`, "This works?Something"},
	}

	for _, test := range tests {
		output := sanitize.XML(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkXML benchmarks the XML method
func BenchmarkXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.XML("<xml>Test This!</xml>")
	}
}

// ExampleXML example using XML()
func ExampleXML() {
	fmt.Println(sanitize.XML("<xml>This?</xml>"))
	// Output: This?
}

// TestXSS tests the XSS sanitize method
func TestXSS_Basic(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Lines 1-11: Common script injection vectors
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

		// Lines 12-21: Inline event handlers
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

		// Line 22: Potential CSS/Style-based attacks
		{"Remove expression(", "expression(", ""},

		// Line 23: Potentially malicious protocols
		{"Remove data:", "data:", ""},

		// Lines 24-26: References to dangerous objects/functions
		{"Remove document.cookie", "document.cookie", ""},
		{"Remove document.write", "document.write", ""},
		{"Remove window.location", "window.location", ""},
	}

	for _, tt := range tests {
		tt := tt // pin variable for parallel sub-tests
		t.Run(tt.name, func(t *testing.T) {
			output := sanitize.XSS(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

// TestXSS_AdditionalCases tests additional cases for the XSS sanitize method
func TestXSS_AdditionalCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Multiple patterns
		{"Multiple patterns", "<script>eval(javascript:alert(1))</script>", ">alert(1))</"},
		// Embedded in text
		{"Pattern in text", "Hello<script>alert(1)</script>World", "Hello>alert(1)</World"},
		// Mixed case
		{"Mixed case script", "<ScRiPt>alert(1)</sCrIpT>", "<ScRiPt>alert(1)</sCrIpT>"},
		// Encoded/obfuscated
		{"HTML entity encoded", "&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;", "&#x3C;script&#x3E;alert(1)&#x3C;/script&#x3E;"},
		// Whitespace in tag
		{"Whitespace in tag", "<scr ipt>alert(1)</scr ipt>", "<scr ipt>alert(1)</scr ipt>"},
		// Inline event handler in tag
		{"Inline event handler", "<img src=x onerror=alert(1)>", "<img src=x alert(1)>"},
		// Obfuscated event handler
		{"Obfuscated event handler", "<img src=x oNclIck=alert(1)>", "<img src=x oNclIck=alert(1)>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := sanitize.XSS(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

// BenchmarkXSS benchmarks the XSS method
func BenchmarkXSS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.XSS("<script>Test This!</script>")
	}
}

// ExampleXSS example using XSS()
func ExampleXSS() {
	fmt.Println(sanitize.XSS("<script>This?</script>"))
	// Output: >This?</
}
