/*
Package sanitize (go-sanitize) implements a simple library of various sanitation methods for data transformation.

This package provides a collection of functions to sanitize and transform different types of data, such as strings, URLs, email addresses, and more. It is designed to help developers clean and format input data to ensure it meets specific criteria and is safe for further processing.

Features:
- Sanitize alpha and alphanumeric characters
- Sanitize Bitcoin and Bitcoin Cash addresses
- Custom regex-based sanitization
- Sanitize decimal numbers and scientific notation
- Sanitize domain names, email addresses, and IP addresses
- Remove HTML/XML tags and scripts
- Sanitize URIs and URLs
- Handle XSS attack strings

Usage:
To use this package, import it and call the desired sanitization function with the input data. Each function is documented with examples in the `sanitize_test.go` file.

Example:

	package main

	import (
	    "fmt"
	    "github.com/mrz1836/go-sanitize"
	)

	func main() {
	    input := "<script>alert('test');</script>"
	    sanitized := sanitize.XSS(input)
	    fmt.Println(sanitized) // Output: >alert('test');</
	}

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.
*/
package sanitize

import (
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

// Set all the regular expressions
var (
	alphaNumericRegExp           = regexp.MustCompile(`[^a-zA-Z0-9]`)                                                             // Alpha numeric
	alphaNumericWithSpacesRegExp = regexp.MustCompile(`[^a-zA-Z0-9\s]`)                                                           // Alphanumeric (with spaces)
	alphaRegExp                  = regexp.MustCompile(`[^a-zA-Z]`)                                                                // Alpha characters
	alphaWithSpacesRegExp        = regexp.MustCompile(`[^a-zA-Z\s]`)                                                              // Alpha characters (with spaces)
	bitcoinCashAddrRegExp        = regexp.MustCompile(`[^ac-hj-np-zAC-HJ-NP-Z02-9]`)                                              // Bitcoin `cashaddr` address accepted characters
	bitcoinRegExp                = regexp.MustCompile(`[^a-km-zA-HJ-NP-Z1-9]`)                                                    // Bitcoin address accepted characters
	decimalRegExp                = regexp.MustCompile(`[^0-9.-]`)                                                                 // Decimals (positive and negative)
	domainRegExp                 = regexp.MustCompile(`[^a-zA-Z0-9-.]`)                                                           // Domain accepted characters
	emailRegExp                  = regexp.MustCompile(`[^a-zA-Z0-9-_.@+]`)                                                        // Email address characters
	formalNameRegExp             = regexp.MustCompile(`[^a-zA-Z0-9-',.\s]`)                                                       // Characters recognized in surnames and proper names
	htmlRegExp                   = regexp.MustCompile(`(?i)<[^>]*>`)                                                              // HTML/XML tags or any alligator open/close tags
	ipAddressRegExp              = regexp.MustCompile(`[^a-zA-Z0-9:.]`)                                                           // IPV4 and IPV6 characters only
	numericRegExp                = regexp.MustCompile(`[^0-9]`)                                                                   // Numbers only
	pathNameRegExp               = regexp.MustCompile(`[^a-zA-Z0-9-_]`)                                                           // Path name (file name, seo)
	punctuationRegExp            = regexp.MustCompile(`[^a-zA-Z0-9-'"#&!?,.\s]+`)                                                 // Standard accepted punctuation characters
	scientificNotationRegExp     = regexp.MustCompile(`[^0-9.eE+-]`)                                                              // Scientific Notation (float) (positive and negative)
	scriptRegExp                 = regexp.MustCompile(`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`) // Scripts and embeds
	singleLineRegExp             = regexp.MustCompile(`(\r)|(\n)|(\t)|(\v)|(\f)`)                                                 // Carriage returns, line feeds, tabs, for single line transition
	timeRegExp                   = regexp.MustCompile(`[^0-9:]`)                                                                  // Time allowed characters
	uriRegExp                    = regexp.MustCompile(`[^a-zA-Z0-9-_/?&=#%]`)                                                     // URI allowed characters
	urlRegExp                    = regexp.MustCompile(`[^a-zA-Z0-9-_/:.,?&@=#%]`)                                                 // URL allowed characters
	wwwRegExp                    = regexp.MustCompile(`(?i)www.`)                                                                 // For removing www
)

// emptySpace is an empty space for replacing
var emptySpace = []byte("")

// Alpha returns a string containing only alphabetic characters (a-z, A-Z).
// If the `spaces` parameter is set to true, spaces will be preserved in the output.
//
// Parameters:
// - original: The input string to be sanitized.
// - spaces: A boolean flag indicating whether spaces should be preserved.
//
// Returns:
// - A sanitized string containing only alphabetic characters and, optionally, spaces.
//
// Example:
//
//	input := "Hello, World! 123"
//	result := sanitize.Alpha(input, true)
//	fmt.Println(result) // Output: "Hello World"
//
// View more examples in the `sanitize_test.go` file.
func Alpha(original string, spaces bool) string {

	// Leave white spaces?
	if spaces {
		return string(alphaWithSpacesRegExp.ReplaceAll([]byte(original), emptySpace))
	}

	// No spaces
	return string(alphaRegExp.ReplaceAll([]byte(original), emptySpace))
}

// AlphaNumeric returns a string containing only alphanumeric characters (a-z, A-Z, 0-9).
// If the `spaces` parameter is set to true, spaces will be preserved in the output.
//
// Parameters:
// - original: The input string to be sanitized.
// - spaces: A boolean flag indicating whether spaces should be preserved.
//
// Returns:
// - A sanitized string containing only alphanumeric characters and, optionally, spaces.
//
// Example:
//
//	input := "Hello, World! 123"
//	result := sanitize.AlphaNumeric(input, true)
//	fmt.Println(result) // Output: "Hello World 123"
//
// View more examples in the `sanitize_test.go` file.
func AlphaNumeric(original string, spaces bool) string {

	// Leave white spaces?
	if spaces {
		return string(alphaNumericWithSpacesRegExp.ReplaceAll([]byte(original), emptySpace))
	}

	// No spaces
	return string(alphaNumericRegExp.ReplaceAll([]byte(original), emptySpace))
}

// BitcoinAddress returns a sanitized string containing only valid characters for a Bitcoin address.
// This function removes any characters that are not part of the accepted Bitcoin address format.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid Bitcoin address characters.
//
// Example:
//
//	input := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa!@#"
//	result := sanitize.BitcoinAddress(input)
//	fmt.Println(result) // Output: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
//
// View more examples in the `sanitize_test.go` file.
func BitcoinAddress(original string) string {
	return string(bitcoinRegExp.ReplaceAll([]byte(original), emptySpace))
}

// BitcoinCashAddress returns a sanitized string containing only valid characters for a Bitcoin Cash address (cashaddr format).
// This function removes any characters that are not part of the accepted Bitcoin Cash address format.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid Bitcoin Cash address characters.
//
// Example:
//
//	input := "bitcoincash:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a!@#"
//	result := sanitize.BitcoinCashAddress(input)
//	fmt.Println(result) // Output: "bitcoincash:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"
//
// View more examples in the `sanitize_test.go` file.
func BitcoinCashAddress(original string) string {
	return string(bitcoinCashAddrRegExp.ReplaceAll([]byte(original), emptySpace))
}

// Custom uses a custom regex string and returns the sanitized result.
// This function allows for flexible sanitization based on user-defined regular expressions.
//
// Parameters:
// - original: The input string to be sanitized.
// - regExp: A string representing the custom regular expression to be used for sanitization.
//
// Returns:
// - A sanitized string based on the provided regular expression.
//
// Example:
//
//	input := "Hello, World! 123"
//	customRegExp := `[^a-zA-Z\s]`
//	result := sanitize.Custom(input, customRegExp)
//	fmt.Println(result) // Output: "Hello World"
//
// View more examples in the `sanitize_test.go` file.
func Custom(original string, regExp string) string {

	// Return the processed string or panic if regex fails
	return string(regexp.MustCompile(regExp).ReplaceAll([]byte(original), emptySpace))
}

// CustomCompiled returns a sanitized string using a pre-compiled regular
// expression. This function provides better performance when the same pattern is
// reused across multiple calls.
//
// Parameters:
// - original: The input string to be sanitized.
// - re: A compiled regular expression used for sanitization.
//
// Returns:
// - A sanitized string based on the provided regular expression.
//
// Example:
//
//	input := "Hello, World! 123"
//	customRegExp := regexp.MustCompile(`[^a-zA-Z\s]`)
//	result := sanitize.CustomCompiled(input, customRegExp)
//	fmt.Println(result) // Output: "Hello World"
//
// View more examples in the `sanitize_test.go` file.
func CustomCompiled(original string, re *regexp.Regexp) string {
	return re.ReplaceAllString(original, "")
}

// Decimal returns a sanitized string containing only decimal/float values, including positive and negative numbers.
// This function removes any characters that are not part of the accepted decimal format.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only decimal/float values.
//
// Example:
//
//	input := "The price is -123.45 USD"
//	result := sanitize.Decimal(input)
//	fmt.Println(result) // Output: "-123.45"
//
// View more examples in the `sanitize_test.go` file.
func Decimal(original string) string {
	return string(decimalRegExp.ReplaceAll([]byte(original), emptySpace))
}

// Domain returns a properly formatted hostname or domain name.
// This function can preserve the case of the original input or convert it to lowercase,
// and optionally remove the "www" subdomain.
//
// Parameters:
// - original: The input string to be sanitized.
// - preserveCase: A boolean flag indicating whether to preserve the case of the original input.
// - removeWww: A boolean flag indicating whether to remove the "www" subdomain.
//
// Returns:
// - A sanitized string containing a valid hostname or domain name.
// - An error if the URL parsing fails.
//
// Example:
//
//	input := "www.Example.com"
//	result, err := sanitize.Domain(input, false, true)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // Output: "example.com"
//
// View more examples in the `sanitize_test.go` file.
func Domain(original string, preserveCase bool, removeWww bool) (string, error) {

	// Try to see if we have a host
	if len(original) == 0 {
		return original, nil
	}

	// Missing http?
	if !strings.Contains(original, "http") {
		original = "http://" + strings.TrimSpace(original)
	}

	// Try to parse the url
	u, err := url.Parse(original)
	if err != nil {
		return original, err
	}

	// Remove leading www.
	if removeWww {
		u.Host = wwwRegExp.ReplaceAllString(u.Host, "")
	}

	// Keeps the exact case of the original input string
	if preserveCase {
		return string(domainRegExp.ReplaceAll([]byte(u.Host), emptySpace)), nil
	}

	// Generally, all domains should be uniform and lowercase
	return string(domainRegExp.ReplaceAll([]byte(strings.ToLower(u.Host)), emptySpace)), nil
}

// Email returns a sanitized email address string. Email addresses are forced
// to lowercase and remove any mail-to prefixes.
//
// Parameters:
// - original: The input string to be sanitized.
// - preserveCase: A boolean flag indicating whether to preserve the case of the original input.
//
// Returns:
// - A sanitized string containing a valid email address.
//
// Example:
//
//	input := "MailTo:Example@DOMAIN.com"
//	result := sanitize.Email(input, false)
//	fmt.Println(result) // Output: "example@domain.com"
//
// View more examples in the `sanitize_test.go` file.
func Email(original string, preserveCase bool) string {

	// Leave the email address in its original case
	if preserveCase {
		return string(emailRegExp.ReplaceAll(
			[]byte(strings.ReplaceAll(original, "mailto:", "")), emptySpace),
		)
	}

	// Standard is forced to lowercase
	return string(emailRegExp.ReplaceAll(
		[]byte(strings.ToLower(strings.ReplaceAll(original, "mailto:", ""))), emptySpace),
	)
}

// FirstToUpper overwrites the first letter as an uppercase letter
// and preserves the rest of the string.
//
// This function is useful for formatting strings where the first character
// needs to be capitalized, such as names or titles.
//
// Parameters:
// - original: The input string to be formatted.
//
// Returns:
// - A string with the first letter converted to uppercase.
//
// Example:
//
//	input := "hello world"
//	result := sanitize.FirstToUpper(input)
//	fmt.Println(result) // Output: "Hello world"
//
// View more examples in the `sanitize_test.go` file.
func FirstToUpper(original string) string {

	// Handle empty and 1 character strings
	if len(original) < 2 {
		return strings.ToUpper(original)
	}

	runes := []rune(original)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// FormalName returns a sanitized string containing only characters recognized in formal names or surnames.
// This function removes any characters that are not part of the accepted formal name format.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid formal name characters.
//
// Example:
//
//	input := "John D'oe, Jr."
//	result := sanitize.FormalName(input)
//	fmt.Println(result) // Output: "John Doe Jr"
//
// View more examples in the `sanitize_test.go` file.
func FormalName(original string) string {
	return string(formalNameRegExp.ReplaceAll([]byte(original), emptySpace))
}

// HTML returns a string without any HTML tags.
// This function removes all HTML tags from the input string, leaving only the text content.
//
// Parameters:
// - original: The input string containing HTML tags to be sanitized.
//
// Returns:
// - A sanitized string with all HTML tags removed.
//
// Example:
//
//	input := "<div>Hello <b>World</b>!</div>"
//	result := sanitize.HTML(input)
//	fmt.Println(result) // Output: "Hello World!"
//
// View more examples in the `sanitize_test.go` file.
func HTML(original string) string {
	return string(htmlRegExp.ReplaceAll([]byte(original), emptySpace))
}

// IPAddress returns a sanitized IP address string for both IPv4 and IPv6 formats.
// This function removes any invalid characters from the input string and attempts to parse it as an IP address.
// If the input string does not contain a valid IP address, an empty string is returned.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing a valid IP address, or an empty string if the input is not a valid IP address.
//
// Example:
//
//	input := "192.168.1.1!@#"
//	result := sanitize.IPAddress(input)
//	fmt.Println(result) // Output: "192.168.1.1"
//
// View more examples in the `sanitize_test.go` file.
func IPAddress(original string) string {
	// Parse the IP - Remove any invalid characters first
	ipAddress := net.ParseIP(
		string(ipAddressRegExp.ReplaceAll([]byte(original), emptySpace)),
	)
	if ipAddress == nil {
		return ""
	}

	return ipAddress.String()
}

// Numeric returns a string containing only numeric characters (0-9).
// This function removes any characters that are not digits from the input string.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only numeric characters.
//
// Example:
//
//	input := "Phone: 123-456-7890"
//	result := sanitize.Numeric(input)
//	fmt.Println(result) // Output: "1234567890"
//
// View more examples in the `sanitize_test.go` file.
func Numeric(original string) string {
	return string(numericRegExp.ReplaceAll([]byte(original), emptySpace))
}

// PathName returns a formatted path-compliant name.
// This function removes any characters that are not valid in file or directory names,
// ensuring the resulting string is safe to use as a path component.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid path name characters.
//
// Example:
//
//	input := "file:name/with*invalid|chars"
//	result := sanitize.PathName(input)
//	fmt.Println(result) // Output: "filenamewithinvalidchars"
//
// View more examples in the `sanitize_test.go` file.
func PathName(original string) string {
	return string(pathNameRegExp.ReplaceAll([]byte(original), emptySpace))
}

// Punctuation returns a string with basic punctuation preserved.
// This function removes any characters that are not standard punctuation or alphanumeric characters,
// ensuring the resulting string contains only valid punctuation and alphanumeric characters.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid punctuation and alphanumeric characters.
//
// Example:
//
//	input := "Hello, World! How's it going? (Good, I hope.)"
//	result := sanitize.Punctuation(input)
//	fmt.Println(result) // Output: "Hello, World! How's it going? (Good, I hope.)"
//
// View more examples in the `sanitize_test.go` file.
func Punctuation(original string) string {
	return string(punctuationRegExp.ReplaceAll([]byte(original), emptySpace))
}

// ScientificNotation returns a sanitized string containing only valid characters for scientific notation.
// This function removes any characters that are not part of the accepted scientific notation format,
// including digits (0-9), decimal points, and the characters 'e', 'E', '+', and '-'.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid scientific notation characters.
//
// Example:
//
//	input := "The value is 1.23e+10 and 4.56E-7."
//	result := sanitize.ScientificNotation(input)
//	fmt.Println(result) // Output: "1.23e+104.56E-7"
//
// View more examples in the `sanitize_test.go` file.
func ScientificNotation(original string) string {
	return string(scientificNotationRegExp.ReplaceAll([]byte(original), emptySpace))
}

// Scripts removes all script, iframe, embed, and object tags from the input string.
// This function is designed to sanitize input by removing potentially harmful tags
// that can be used for cross-site scripting (XSS) attacks or other malicious purposes.
//
// Parameters:
// - original: The input string containing HTML or script tags to be sanitized.
//
// Returns:
// - A sanitized string with all script, iframe, embed, and object tags removed.
//
// Example:
//
//	input := "<script>alert('test');</script><iframe src='example.com'></iframe>"
//	result := sanitize.Scripts(input)
//	fmt.Println(result) // Output: "alert('test');"
//
// View more examples in the `sanitize_test.go` file.
func Scripts(original string) string {
	return string(scriptRegExp.ReplaceAll([]byte(original), emptySpace))
}

// SingleLine returns a single line string by removing all carriage returns, line feeds, tabs, vertical tabs, and form feeds.
// This function is useful for sanitizing input that should be represented as a single line of text, ensuring that
// any multi-line or formatted input is condensed into a single line.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string with all line breaks and whitespace characters replaced by a single space.
//
// Example:
//
//	input := "This is a\nmulti-line\tstring."
//	result := sanitize.SingleLine(input)
//	fmt.Println(result) // Output: "This is a multi-line string."
//
// View more examples in the `sanitize_test.go` file.
func SingleLine(original string) string {
	return singleLineRegExp.ReplaceAllString(original, " ")
}

// Time returns just the time part of the string.
// This function removes any characters that are not valid in a time format (HH:MM or HH:MM:SS),
// ensuring the resulting string contains only valid time characters.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid time characters.
//
// Example:
//
//	input := "t00:00d -EST"
//	result := sanitize.Time(input)
//	fmt.Println(result) // Output: "00:00"
//
// View more examples in the `sanitize_test.go` file.
func Time(original string) string {
	return string(timeRegExp.ReplaceAll([]byte(original), emptySpace))
}

// URI returns a sanitized string containing only valid URI characters.
// This function removes any characters that are not part of the accepted URI format,
// including alphanumeric characters, dashes, underscores, slashes, question marks, ampersands, equals signs, hashes, and percent signs.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid URI characters.
//
// Example:
//
//	input := "Test?=what! &this=that"
//	result := sanitize.URI(input)
//	fmt.Println(result) // Output: "Test?=what&this=that"
//
// View more examples in the `sanitize_test.go` file.
func URI(original string) string {
	return string(uriRegExp.ReplaceAll([]byte(original), emptySpace))
}

// URL returns a formatted URL-friendly string.
// This function removes any characters that are not part of the accepted URL format,
// including alphanumeric characters, dashes, underscores, slashes, colons, periods, question marks, ampersands, at signs, equals signs, hashes, and percent signs.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string containing only valid URL characters.
//
// Example:
//
//	input := "https://Example.com/This/Works?^No&this"
//	result := sanitize.URL(input)
//	fmt.Println(result) // Output: "https://Example.com/This/Works?No&this"
//
// View more examples in the `sanitize_test.go` file.
func URL(original string) string {
	return string(urlRegExp.ReplaceAll([]byte(original), emptySpace))
}

// XML returns a string without any XML tags.
// This function removes all XML tags from the input string, leaving only the text content.
// It is an alias for the HTML function, which performs the same operation.
//
// Parameters:
// - original: The input string containing XML tags to be sanitized.
//
// Returns:
// - A sanitized string with all XML tags removed.
//
// Example:
//
//	input := `<?xml version="1.0" encoding="UTF-8"?><note>Something</note>`
//	result := sanitize.XML(input)
//	fmt.Println(result) // Output: "Something"
//
// View more examples in the `sanitize_test.go` file.
func XML(original string) string {
	return HTML(original)
}

// XSS removes known XSS attack strings or script strings.
// This function sanitizes the input string by removing common XSS attack vectors,
// such as script tags, eval functions, and JavaScript protocol handlers.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A sanitized string with known XSS attack vectors removed.
//
// Example:
//
//	input := "<script>alert('test');</script>"
//	result := sanitize.XSS(input)
//	fmt.Println(result) // Output: ">alert('test');</"
//
// View more examples in the `sanitize_test.go` file.
func XSS(original string) string {
	original = strings.ReplaceAll(original, "<script", "")
	original = strings.ReplaceAll(original, "script>", "")
	original = strings.ReplaceAll(original, "eval(", "")
	original = strings.ReplaceAll(original, "eval&#40;", "")
	original = strings.ReplaceAll(original, "javascript:", "")
	original = strings.ReplaceAll(original, "javascript&#58;", "")
	original = strings.ReplaceAll(original, "fromCharCode", "")
	original = strings.ReplaceAll(original, "&#62;", "")
	original = strings.ReplaceAll(original, "&#60;", "")
	original = strings.ReplaceAll(original, "&lt;", "")
	original = strings.ReplaceAll(original, "&rt;", "")

	// Some inline event handlers
	original = strings.ReplaceAll(original, "onclick=", "")
	original = strings.ReplaceAll(original, "onerror=", "")
	original = strings.ReplaceAll(original, "onload=", "")
	original = strings.ReplaceAll(original, "onmouseover=", "")
	original = strings.ReplaceAll(original, "onfocus=", "")
	original = strings.ReplaceAll(original, "onblur=", "")
	original = strings.ReplaceAll(original, "ondblclick=", "")
	original = strings.ReplaceAll(original, "onkeydown=", "")
	original = strings.ReplaceAll(original, "onkeyup=", "")
	original = strings.ReplaceAll(original, "onkeypress=", "")

	// Potential CSS/Style-based attacks
	original = strings.ReplaceAll(original, "expression(", "")

	// Potentially malicious protocols
	original = strings.ReplaceAll(original, "data:", "")

	// Potential references to dangerous objects/functions
	original = strings.ReplaceAll(original, "document.cookie", "")
	original = strings.ReplaceAll(original, "document.write", "")
	original = strings.ReplaceAll(original, "window.location", "")

	return original
}
