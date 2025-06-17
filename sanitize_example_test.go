package sanitize_test

import (
	"fmt"
	"regexp"

	"github.com/mrz1836/go-sanitize"
)

// ExampleAlpha example using Alpha() and no spaces flag
func ExampleAlpha() {
	fmt.Println(sanitize.Alpha("Example String!", false))
	// Output: ExampleString
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

// ExampleAlpha_withSpaces example using Alpha with a space flag
func ExampleAlpha_withSpaces() {
	fmt.Println(sanitize.Alpha("Example String!", true))
	// Output: Example String
}

// ExampleBitcoinAddress example using BitcoinAddress()
func ExampleBitcoinAddress() {
	fmt.Println(sanitize.BitcoinAddress(":1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!"))
	// Output: 1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs
}

// ExampleBitcoinCashAddress example using BitcoinCashAddress() `cashaddr`
func ExampleBitcoinCashAddress() {
	fmt.Println(sanitize.BitcoinAddress("qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz!"))
	// Output: qze7yy2au5vuznvn8zj5yj5t66vhs75e3meptz
}

// ExampleCustom example using Custom() using an alpha regex
func ExampleCustom() {
	fmt.Println(sanitize.Custom("Example String 2!", `[^a-zA-Z]`))
	// Output: ExampleString
}

// ExampleCustomCompiled example using CustomCompiled with an alpha regex
func ExampleCustomCompiled() {
	re := regexp.MustCompile(`[^a-zA-Z]`)
	fmt.Println(sanitize.CustomCompiled("Example String 2!", re))
	// Output: ExampleString
}

// ExampleCustom_numeric example using Custom() using a numeric regex
func ExampleCustom_numeric() {
	fmt.Println(sanitize.Custom("Example String 2!", `[^0-9]`))
	// Output: 2
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

// ExampleFirstToUpper example using FirstToUpper()
func ExampleFirstToUpper() {
	fmt.Println(sanitize.FirstToUpper("this works"))
	// Output: This works
}

// ExampleFormalName example using FormalName()
func ExampleFormalName() {
	fmt.Println(sanitize.FormalName("John McDonald Jr.!"))
	// Output: John McDonald Jr.
}

// ExampleHTML example using HTML()
func ExampleHTML() {
	fmt.Println(sanitize.HTML("<body>This Works?</body>"))
	// Output: This Works?
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

// ExampleNumeric example using Numeric()
func ExampleNumeric() {
	fmt.Println(sanitize.Numeric("This:123 + 90!"))
	// Output: 12390
}

// ExampleNumeric example using PathName()
func ExamplePathName() {
	fmt.Println(sanitize.PathName("/This-Works_Now-123/!"))
	// Output: This-Works_Now-123
}

// ExamplePunctuation example using Punctuation()
func ExamplePunctuation() {
	fmt.Println(sanitize.Punctuation(`[@"Does" 'this' work?@] this too`))
	// Output: "Does" 'this' work? this too
}

// ExampleScientificNotation example using ScientificNotation() for a positive number
func ExampleScientificNotation() {
	fmt.Println(sanitize.ScientificNotation("$ 1.096e-3!"))
	// Output: 1.096e-3
}

// ExampleScripts example using Scripts()
func ExampleScripts() {
	fmt.Println(sanitize.Scripts(`Does<script>This</script>Work?`))
	// Output: DoesWork?
}

// ExampleSingleLine example using SingleLine()
func ExampleSingleLine() {
	fmt.Println(sanitize.SingleLine(`Does
This
Work?`))
	// Output: Does This Work?
}

// ExampleTime example using Time()
func ExampleTime() {
	fmt.Println(sanitize.Time(`Time 01:02:03!`))
	// Output: 01:02:03
}

// ExampleURI example using URI()
func ExampleURI() {
	fmt.Println(sanitize.URI("/This/Works?^No&this"))
	// Output: /This/Works?No&this
}

// ExampleURL example using URL()
func ExampleURL() {
	fmt.Println(sanitize.URL("https://Example.com/This/Works?^No&this"))
	// Output: https://Example.com/This/Works?No&this
}

// ExampleXML example using XML()
func ExampleXML() {
	fmt.Println(sanitize.XML("<xml>This?</xml>"))
	// Output: This?
}

// ExampleXSS example using XSS()
func ExampleXSS() {
	fmt.Println(sanitize.XSS("<script>This?</script>"))
	// Output: >This?</
}
