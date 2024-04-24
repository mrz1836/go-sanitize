package sanitize

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAlpha tests the alpha sanitize method
func TestAlpha(t *testing.T) {
	t.Parallel()

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
		{"carriage returns", "\nThis\nThat", `
This
That`, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Alpha(test.input, test.typeCase)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkAlphaNoSpaces benchmarks the Alpha method
func BenchmarkAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Alpha("This is the test string.", false)
	}
}

// BenchmarkAlpha_WithSpaces benchmarks the Alpha method
func BenchmarkAlpha_WithSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Alpha("This is the test string.", true)
	}
}

// ExampleAlpha example using Alpha() and no spaces flag
func ExampleAlpha() {
	fmt.Println(Alpha("Example String!", false))
	// Output: ExampleString
}

// ExampleAlpha_withSpaces example using Alpha with spaces flag
func ExampleAlpha_withSpaces() {
	fmt.Println(Alpha("Example String!", true))
	// Output: Example String
}

// TestAlphaNumeric tests the alphanumeric sanitize method
func TestAlphaNumeric(t *testing.T) {
	t.Parallel()

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
		{"carriage return and spaces", "\nThis1\nThat2", `
This1
That2`, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := AlphaNumeric(test.input, test.typeCase)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkAlphaNumeric benchmarks the AlphaNumeric method
func BenchmarkAlphaNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AlphaNumeric("This is the test string 12345.", false)
	}
}

// BenchmarkAlphaNumeric_WithSpaces benchmarks the AlphaNumeric method
func BenchmarkAlphaNumeric_WithSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AlphaNumeric("This is the test string 12345.", true)
	}
}

// ExampleAlphaNumeric example using AlphaNumeric() with no spaces
func ExampleAlphaNumeric() {
	fmt.Println(AlphaNumeric("Example String 2!", false))
	// Output: ExampleString2
}

// ExampleAlphaNumeric_withSpaces example using AlphaNumeric() with spaces
func ExampleAlphaNumeric_withSpaces() {
	fmt.Println(AlphaNumeric("Example String 2!", true))
	// Output: Example String 2
}

// TestBitcoinAddress will test all permutations
func TestBitcoinAddress(t *testing.T) {
	t.Parallel()

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
			output := BitcoinAddress(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkBitcoinAddress benchmarks the BitcoinAddress method
func BenchmarkBitcoinAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BitcoinAddress("1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs")
	}
}

// ExampleBitcoinAddress example using BitcoinAddress()
func ExampleBitcoinAddress() {
	fmt.Println(BitcoinAddress(":1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!"))
	// Output: 1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs
}

// TestBitcoinCashAddress will test all permutations of using BitcoinCashAddress()
func TestBitcoinCashAddress(t *testing.T) {
	t.Parallel()

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
			output := BitcoinCashAddress(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkBitcoinCashAddress benchmarks the BitcoinCashAddress() method
func BenchmarkBitcoinCashAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BitcoinCashAddress("qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz")
	}
}

// ExampleBitcoinCashAddress example using BitcoinCashAddress() `cashaddr`
func ExampleBitcoinCashAddress() {
	fmt.Println(BitcoinAddress("qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz!"))
	// Output: qze7yy2au5vuznvn8zj5yj5t66vhs75e3meptz
}

// TestCustom tests the custom sanitize method
func TestCustom(t *testing.T) {
	t.Parallel()

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
		output := Custom(test.input, test.regex)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkCustom benchmarks the Custom method
func BenchmarkCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Custom("This is the test string 12345.", `[^a-zA-Z0-9]`)
	}
}

// ExampleCustom example using Custom() using an alpha regex
func ExampleCustom() {
	fmt.Println(Custom("Example String 2!", `[^a-zA-Z]`))
	// Output: ExampleString
}

// ExampleCustom_numeric example using Custom() using a numeric regex
func ExampleCustom_numeric() {
	fmt.Println(Custom("Example String 2!", `[^0-9]`))
	// Output: 2
}

// TestDecimal tests the decimal sanitize method
func TestDecimal(t *testing.T) {
	t.Parallel()

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
		output := Decimal(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkDecimal benchmarks the Decimal method
func BenchmarkDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Decimal("String: -123.12345")
	}
}

// ExampleDecimal example using Decimal() for a positive number
func ExampleDecimal() {
	fmt.Println(Decimal("$ 99.99!"))
	// Output: 99.99
}

// ExampleDecimal_negative example using Decimal() for a negative number
func ExampleDecimal_negative() {
	fmt.Println(Decimal("$ -99.99!"))
	// Output: -99.99
}

// TestDomain tests the domain sanitize method
func TestDomain(t *testing.T) {
	t.Parallel()

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
				"domain with tabs and spaces",
				`		domain.com`,
				"domain.com",
				false,
				true,
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				output, err := Domain(test.input, test.preserveCase, test.removeWww)
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
				output, err := Domain(test.input, test.preserveCase, test.removeWww)
				require.Error(t, err)
				assert.Equal(t, test.expected, output)
			})
		}
	})
}

// BenchmarkDomain benchmarks the Domain method
func BenchmarkDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Domain("https://Example.COM/?param=value", false, false)
	}
}

// BenchmarkDomain_PreserveCase benchmarks the Domain method
func BenchmarkDomain_PreserveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Domain("https://Example.COM/?param=value", true, false)
	}
}

// BenchmarkDomain_RemoveWww benchmarks the Domain method
func BenchmarkDomain_RemoveWww(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Domain("https://Example.COM/?param=value", false, true)
	}
}

// ExampleDomain example using Domain()
func ExampleDomain() {
	fmt.Println(Domain("https://www.Example.COM/?param=value", false, false))
	// Output: www.example.com <nil>
}

// ExampleDomain_preserveCase example using Domain() and preserving the case
func ExampleDomain_preserveCase() {
	fmt.Println(Domain("https://www.Example.COM/?param=value", true, false))
	// Output: www.Example.COM <nil>
}

// ExampleDomain_removeWww example using Domain() and removing the www subdomain
func ExampleDomain_removeWww() {
	fmt.Println(Domain("https://www.Example.COM/?param=value", false, true))
	// Output: example.com <nil>
}

// TestEmail tests the email sanitize method
func TestEmail(t *testing.T) {
	t.Parallel()

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
		output := Email(test.input, test.preserveCase)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkEmail benchmarks the Email method
func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Email("mailto:Person@Example.COM ", false)
	}
}

// BenchmarkEmail_PreserveCase benchmarks the Email method
func BenchmarkEmail_PreserveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Email("mailto:Person@Example.COM ", true)
	}
}

// ExampleEmail example using Email()
func ExampleEmail() {
	fmt.Println(Email("mailto:Person@Example.COM", false))
	// Output: person@example.com
}

// ExampleEmail_preserveCase example using Email() and preserving the case
func ExampleEmail_preserveCase() {
	fmt.Println(Email("mailto:Person@Example.COM", true))
	// Output: Person@Example.COM
}

// TestFirstToUpper tests the first to upper method
func TestFirstToUpper(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"thisworks", "Thisworks"},
		{"Thisworks", "Thisworks"},
		{"this", "This"},
		{"t", "T"},
		{"tt", "Tt"},
	}

	for _, test := range tests {
		output := FirstToUpper(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkFirstToUpper benchmarks the FirstToUpper method
func BenchmarkFirstToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FirstToUpper("make this upper")
	}
}

// ExampleFirstToUpper example using FirstToUpper()
func ExampleFirstToUpper() {
	fmt.Println(FirstToUpper("this works"))
	// Output: This works
}

// TestFormalName tests the formal name method
func TestFormalName(t *testing.T) {
	t.Parallel()

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
		output := FormalName(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkFormalName benchmarks the FormalName method
func BenchmarkFormalName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormalName("John McDonald Jr.")
	}
}

// ExampleFormalName example using FormalName()
func ExampleFormalName() {
	fmt.Println(FormalName("John McDonald Jr.!"))
	// Output: John McDonald Jr.
}

// TestHTML tests the HTML sanitize method
func TestHTML(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"<b>This works?</b>", "This works?"},
		{"<html><b>This works?</b><i></i></br></html>", "This works?"},
		{"<html><b class='test'>This works?</b><i></i></br></html>", "This works?"},
	}

	for _, test := range tests {
		output := HTML(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkHTML benchmarks the HTML method
func BenchmarkHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FormalName("<html><b>Test This!</b></html>")
	}
}

// ExampleHTML example using HTML()
func ExampleHTML() {
	fmt.Println(HTML("<body>This Works?</body>"))
	// Output: This Works?
}

// TestIPAddress tests the ip address sanitize method
func TestIPAddress(t *testing.T) {
	t.Parallel()

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
		output := IPAddress(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkIPAddress benchmarks the IPAddress method
func BenchmarkIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IPAddress(" 192.168.0.1 ")
	}
}

// BenchmarkIPAddress_V6 benchmarks the IPAddress method
func BenchmarkIPAddress_IPV6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IPAddress(" 2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f ")
	}
}

// ExampleIPAddress example using IPAddress() for IPV4 address
func ExampleIPAddress() {
	fmt.Println(IPAddress(" 192.168.0.1 "))
	// Output: 192.168.0.1
}

// ExampleIPAddress_ipv6 example using IPAddress() for IPV6 address
func ExampleIPAddress_ipv6() {
	fmt.Println(IPAddress(" 2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f "))
	// Output: 2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f
}

// TestNumeric tests the numeric sanitize method
func TestNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{" > Test This String-!1234", "1234"},
		{" $1.00 Price!", "100"},
	}

	for _, test := range tests {
		output := Numeric(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkNumeric benchmarks the numeric method
func BenchmarkNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Numeric(" 192.168.0.1 ")
	}
}

// ExampleNumeric example using Numeric()
func ExampleNumeric() {
	fmt.Println(Numeric("This:123 + 90!"))
	// Output: 12390
}

// TestPathName tests the path name sanitize method
func TestPathName(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"My BadPath (10)", "MyBadPath10"},
		{"My BadPath (10)[]()!$", "MyBadPath10"},
		{"My_Folder-Path-123_TEST", "My_Folder-Path-123_TEST"},
	}

	for _, test := range tests {
		output := PathName(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkPathName benchmarks the PathName method
func BenchmarkPathName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PathName("/This-Path-Name_Works-/")
	}
}

// ExampleNumeric example using PathName()
func ExamplePathName() {
	fmt.Println(PathName("/This-Works_Now-123/!"))
	// Output: This-Works_Now-123
}

// TestPunctuation tests the punctuation method
func TestPunctuation(t *testing.T) {
	t.Parallel()

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
		output := Punctuation(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkPunctuation benchmarks the Punctuation method
func BenchmarkPunctuation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Punctuation("Does this work? They're doing it?")
	}
}

// ExamplePunctuation example using Punctuation()
func ExamplePunctuation() {
	fmt.Println(Punctuation(`[@"Does" 'this' work?@] this too`))
	// Output: "Does" 'this' work? this too
}

// TestScientificNotation tests the scientific notation sanitize method
func TestScientificNotation(t *testing.T) {
	t.Parallel()

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
		output := ScientificNotation(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkDecimal benchmarks the ScientificNotation method
func BenchmarkScientificNotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ScientificNotation("String: -1.096e-3")
	}
}

// ExampleDecimal example using Decimal() for a positive number
func ExampleScientificNotation() {
	fmt.Println(ScientificNotation("$ 1.096e-3!"))
	// Output: 1.096e-3
}

// TestScripts tests the script removal
func TestScripts(t *testing.T) {
	t.Parallel()

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
		output := Scripts(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkScripts benchmarks the Scripts method
func BenchmarkScripts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Scripts("<script>$(){ var remove='me'; }</script>")
	}
}

// ExampleScripts example using Scripts()
func ExampleScripts() {
	fmt.Println(Scripts(`Does<script>This</script>Work?`))
	// Output: DoesWork?
}

// TestSingleLine test the single line sanitize method
func TestSingleLine(t *testing.T) {
	t.Parallel()

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
		output := SingleLine(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkSingleLine benchmarks the SingleLine method
func BenchmarkSingleLine(b *testing.B) {
	testString := `This line
That Line
Another Line`
	for i := 0; i < b.N; i++ {
		_ = SingleLine(testString)
	}
}

// ExampleSingleLine example using SingleLine()
func ExampleSingleLine() {
	fmt.Println(SingleLine(`Does
This
Work?`))
	// Output: Does This Work?
}

// TestTime tests the time sanitize method
func TestTime(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"t00:00d -EST", "00:00"},
		{"t00:00:00d -EST", "00:00:00"},
		{"SOMETHING t00:00:00d -EST DAY", "00:00:00"},
	}

	for _, test := range tests {
		output := Time(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkTime benchmarks the Time method
func BenchmarkTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Time("Time is 05:10:23")
	}
}

// ExampleTime example using Time()
func ExampleTime() {
	fmt.Println(Time(`Time 01:02:03!`))
	// Output: 01:02:03
}

// TestURI tests the URI sanitize method
func TestURI(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"Test?=what! &this=that", "Test?=what&this=that"},
		{"Test?=what! &this=/that/!()*^", "Test?=what&this=/that/"},
		{"/This/Works/?that=123&this#page10%", "/This/Works/?that=123&this#page10%"},
	}

	for _, test := range tests {
		output := URI(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkURI benchmarks the URI method
func BenchmarkURI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = XSS("/Test/This/Url/?param=value")
	}
}

// ExampleURI example using URI()
func ExampleURI() {
	fmt.Println(URI("/This/Works?^No&this"))
	// Output: /This/Works?No&this
}

// TestURL tests the URL sanitize method
func TestURL(t *testing.T) {
	t.Parallel()

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
			output := URL(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}

// BenchmarkURL benchmarks the URL method
func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = URL("/Test/This/Url/?param=value")
	}
}

// ExampleURL example using URL()
func ExampleURL() {
	fmt.Println(URL("https://Example.com/This/Works?^No&this"))
	// Output: https://Example.com/This/Works?No&this
}

// TestXML tests the XML sanitize method
func TestXML(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{`<?xml version="1.0" encoding="UTF-8"?><note>Something</note>`, "Something"},
		{`<body>This works?</body><title>Something</title>`, "This works?Something"},
	}

	for _, test := range tests {
		output := XML(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkXML benchmarks the XML method
func BenchmarkXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = XML("<xml>Test This!</xml>")
	}
}

// ExampleXML example using XML()
func ExampleXML() {
	fmt.Println(XML("<xml>This?</xml>"))
	// Output: This?
}

// TestXSS tests the XSS sanitize method
func TestXSS(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input    string
		expected string
	}{
		{"<script>alert('test');</script>", ">alert('test');</"},
		{"&lt;script&lt;alert('test');&lt;/script&lt;", "scriptalert('test');/script"},
		{"javascript:alert('test');", "alert('test');"},
		{"eval('test');", "'test');"},
		{"javascript&#58;('test');", "('test');"},
		{"fromCharCode('test');", "('test');"},
		{"&#60;&#62;fromCharCode('test');&#62;&#60;", "('test');"},
	}

	for _, test := range tests {
		output := XSS(test.input)
		assert.Equal(t, test.expected, output)
	}
}

// BenchmarkXSS benchmarks the XSS method
func BenchmarkXSS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = XSS("<script>Test This!</script>")
	}
}

// ExampleXSS example using XSS()
func ExampleXSS() {
	fmt.Println(XSS("<script>This?</script>"))
	// Output: >This?</
}
