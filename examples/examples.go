// Package main demonstrates all functions of the sanitize package.
package main

import (
	"log"
	"regexp"

	"github.com/mrz1836/go-sanitize"
)

func main() {
	// Alpha removes all non-letter characters.
	alphaIn := "Hello, World! 123"
	log.Printf("Alpha(%q) => %q\n", alphaIn, sanitize.Alpha(alphaIn, false))

	// AlphaNumeric removes symbols but preserves letters and numbers.
	alphaNumericIn := "Hello 2nd World!"
	log.Printf("AlphaNumeric(%q) => %q\n", alphaNumericIn, sanitize.AlphaNumeric(alphaNumericIn, false))

	// BitcoinAddress strips invalid bitcoin address characters.
	btcIn := " :1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs!"
	log.Printf("BitcoinAddress(%q) => %q\n", btcIn, sanitize.BitcoinAddress(btcIn))

	// BitcoinCashAddress sanitizes cashaddr formatted addresses.
	bchIn := " qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz!"
	log.Printf("BitcoinCashAddress(%q) => %q\n", bchIn, sanitize.BitcoinCashAddress(bchIn))

	// Custom allows any regular expression for sanitization.
	customIn := "Sample #1 String"
	log.Printf("Custom(%q, `[^a-zA-Z]`) => %q\n", customIn, sanitize.Custom(customIn, `[^a-zA-Z]`))

	// CustomCompiled uses a precompiled regular expression for speed.
	re := regexp.MustCompile(`[^a-zA-Z]`)
	log.Printf("CustomCompiled(%q) => %q\n", customIn, sanitize.CustomCompiled(customIn, re))

	// Decimal keeps digits and decimal separators.
	decIn := "$ -99.99!"
	log.Printf("Decimal(%q) => %q\n", decIn, sanitize.Decimal(decIn))

	// Domain extracts a normalized host name.
	domainIn := "https://www.Example.COM/?param=value"
	cleanedDomain, err := sanitize.Domain(domainIn, false, true)
	if err != nil {
		log.Fatalf("domain error: %v", err)
	}
	log.Printf("Domain(%q) => %q\n", domainIn, cleanedDomain)

	// Email cleans an email address and forces lowercase.
	emailIn := "mailto:Person@Example.COM"
	log.Printf("Email(%q) => %q\n", emailIn, sanitize.Email(emailIn, false))

	// FirstToUpper function capitalizes the first character.
	firstIn := "hello world"
	log.Printf("FirstToUpper(%q) => %q\n", firstIn, sanitize.FirstToUpper(firstIn))

	// FormalName keeps letters, numbers, dashes and common punctuation.
	nameIn := "John D'oe, Jr."
	log.Printf("FormalName(%q) => %q\n", nameIn, sanitize.FormalName(nameIn))

	// HTML strips all HTML tags.
	htmlIn := "<div>Hello <b>World</b></div>"
	log.Printf("HTML(%q) => %q\n", htmlIn, sanitize.HTML(htmlIn))

	// IPAddress validates IPv4 or IPv6 strings.
	ipIn := "192.168.1.1!"
	log.Printf("IPAddress(%q) => %q\n", ipIn, sanitize.IPAddress(ipIn))

	// Numeric keeps only the numeric characters.
	numIn := "Phone: 123-456-7890"
	log.Printf("Numeric(%q) => %q\n", numIn, sanitize.Numeric(numIn))

	// PathName removes characters not safe for file names.
	pathIn := "My File@2025!.txt"
	log.Printf("PathName(%q) => %q\n", pathIn, sanitize.PathName(pathIn))

	// Punctuation retains standard punctuation characters.
	punctuationIn := `[@"Does" 'this' work?@] this too`
	log.Printf("Punctuation(%q) => %q\n", punctuationIn, sanitize.Punctuation(punctuationIn))

	// ScientificNotation keeps floats with exponent notation.
	sciIn := "$ 1.096e-3!"
	log.Printf("ScientificNotation(%q) => %q\n", sciIn, sanitize.ScientificNotation(sciIn))

	// Scripts removes script, iframe, embed, and object tags.
	scriptIn := `Does<script>This</script>Work?`
	log.Printf("Scripts(%q) => %q\n", scriptIn, sanitize.Scripts(scriptIn))

	// SingleLine collapses all whitespace into a single line.
	singleIn := "Does\nThis\tWork?"
	log.Printf("SingleLine(%q) => %q\n", singleIn, sanitize.SingleLine(singleIn))

	// Time strips invalid time characters.
	timeIn := "Time 01:02:03!"
	log.Printf("Time(%q) => %q\n", timeIn, sanitize.Time(timeIn))

	// URI removes characters not valid in URIs.
	uriIn := "/This/Works?^No&this"
	log.Printf("URI(%q) => %q\n", uriIn, sanitize.URI(uriIn))

	// URL removes characters not valid in URLs.
	urlIn := "https://Example.com/This/Works?^No&this"
	log.Printf("URL(%q) => %q\n", urlIn, sanitize.URL(urlIn))

	// XML strips XML tags (alias of HTML).
	xmlIn := "<note>Something</note>"
	log.Printf("XML(%q) => %q\n", xmlIn, sanitize.XML(xmlIn))

	// XSS removes common XSS attack patterns.
	xssIn := "<script>This?</script>"
	log.Printf("XSS(%q) => %q\n", xssIn, sanitize.XSS(xssIn))
}
