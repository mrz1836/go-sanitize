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

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.
*/
package sanitize

import (
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Set all the regular expressions
var (
	htmlRegExp   = regexp.MustCompile(`(?i)<[^>]*>`)                                                              // HTML/XML tags or any alligator open/close tags
	scriptRegExp = regexp.MustCompile(`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`) // Scripts and embeds
)

// emptySpace is an empty space for replacing
var emptySpace = []byte("")

// Alpha returns a string containing only Unicode alphabetic characters from the input.
// Optionally, it preserves spaces if the `spaces` parameter is set to true.
// All non-alphabetic characters (and spaces, if not preserved) are removed.
// This function supports Unicode letters (IsLetter) and is useful for sanitizing names or text fields
// where only letters (and optional spaces) are allowed.
//
// Parameters:
//   - original: The input string to be sanitized.
//   - spaces: If true, spaces are preserved in the output; otherwise, they are removed.
//
// Returns:
//   - A sanitized string containing only Unicode alphabetic characters and, optionally, spaces.
//
// Example:
//
//	input := "Hello, 世界! 123"
//	result := sanitize.Alpha(input, true)
//	fmt.Println(result) // Output: "Hello 世界"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Alpha(original string, spaces bool) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsLetter(r) || (spaces && r == ' ') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// AlphaNumeric returns a string containing only Unicode alphanumeric characters from the input.
// Optionally, it preserves spaces if the `spaces` parameter is set to true.
// All non-alphanumeric characters (and spaces, if not preserved) are removed.
// This function supports Unicode letters and digits, making it suitable for sanitizing user input,
// filenames, or any text where only letters, numbers, and optional spaces are allowed.
//
// Parameters:
//   - original: The input string to be sanitized.
//   - spaces: If true, spaces are preserved in the output; otherwise, they are removed.
//
// Returns:
//   - A sanitized string containing only Unicode alphanumeric characters and, optionally, spaces.
//
// Example:
//
//	input := "Hello, 世界! 123"
//	result := sanitize.AlphaNumeric(input, true)
//	fmt.Println(result) // Output: "Hello 世界 123"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func AlphaNumeric(original string, spaces bool) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || (spaces && r == ' ') {
			b.WriteRune(r)
		}
	}
	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func BitcoinAddress(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if (r >= 'a' && r <= 'k') ||
			(r >= 'm' && r <= 'z') ||
			(r >= 'A' && r <= 'H') ||
			(r >= 'J' && r <= 'N') ||
			(r >= 'P' && r <= 'Z') ||
			(r >= '1' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func BitcoinCashAddress(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if r == '0' ||
			(r >= '2' && r <= '9') ||
			r == 'a' ||
			(r >= 'c' && r <= 'h') ||
			(r >= 'j' && r <= 'n') ||
			(r >= 'p' && r <= 'z') ||
			r == 'A' ||
			(r >= 'C' && r <= 'H') ||
			(r >= 'J' && r <= 'N') ||
			(r >= 'P' && r <= 'Z') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// Custom uses a custom regex string and returns the sanitized result.
// This function allows for flexible sanitization based on user-defined regular expressions.
//
// This function allows for flexible sanitization based on user-defined regular
// expressions. It panics if the provided regular expression cannot be compiled
// successfully.
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Custom(original string, regExp string) string {

	// Return the processed string or panic if regex fails
	return string(regexp.MustCompile(regExp).ReplaceAll([]byte(original), emptySpace))
}

// CustomCompiled returns a sanitized string using a pre-compiled regular
// expression. This function provides better performance when the same pattern is
// reused across multiple calls. Passing a nil regular expression will cause a
// panic.
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Decimal(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsDigit(r) || r == '.' || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Domain(original string, preserveCase bool, removeWww bool) (string, error) {
	if original == "" {
		return original, nil
	}

	// Ensure URL has a scheme for parsing
	original = strings.TrimSpace(original)
	if !strings.HasPrefix(original, "http://") && !strings.HasPrefix(original, "https://") {
		original = "http://" + original
	}

	// Parse the URL to extract the hostname
	u, err := url.Parse(original)
	if err != nil {
		return original, err
	}

	// Extract the hostname from the URL
	host := u.Hostname()

	// Remove leading www.
	if removeWww && len(host) >= 4 &&
		(host[0] == 'w' || host[0] == 'W') &&
		strings.EqualFold(host[:4], "www.") {
		host = host[4:]
	}

	// Convert to lowercase if not preserving our case
	if !preserveCase {
		host = strings.ToLower(host)
	}

	// Filter to valid domain characters: a-z, A-Z, 0-9, hyphen, dot
	var b strings.Builder
	b.Grow(len(host))
	for _, r := range host {
		if r == '-' || r == '.' || (r >= '0' && r <= '9') ||
			(r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			b.WriteRune(r)
		}
	}

	// Return the sanitized domain name
	return b.String(), nil
}

// Email returns a sanitized email address string. Email addresses are forced
// to lowercase by default and remove any MailTo prefixes (case-insensitive).
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Email(original string, preserveCase bool) string {

	// Skip all work for empty string
	if original == "" {
		return original
	}

	// Trim surrounding white space
	original = strings.TrimSpace(original)

	// Remove a leading "mailto:" prefix in any case
	if len(original) >= 7 && strings.EqualFold(original[:7], "mailto:") {
		original = original[7:]
	}

	// Standard output is lowercase
	if !preserveCase {
		original = strings.ToLower(original)
	}

	// Filter to valid email characters
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		valid := r == '@' || r == '.' || r == '_' || r == '-' || r == '+' ||
			(r >= '0' && r <= '9') ||
			(r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
		if valid {
			b.WriteRune(r)
		}
	}

	return b.String()
}

// FirstToUpper returns a copy of the input string with the first Unicode letter
// converted to its uppercase form, leaving the rest of the string unchanged.
// If the input is empty, it returns an empty string. If the input is a single
// character, it returns the uppercase version of that character. This function
// supports multibyte (UTF-8) characters and is useful for capitalizing names,
// titles, or any string where only the first character should be uppercased.
//
// Parameters:
// - original: The input string to be processed.
//
// Returns:
// - A string with the first character uppercased and the remainder unchanged.
//
// Example:
//
//	input := "hello world"
//	result := sanitize.FirstToUpper(input)
//	fmt.Println(result) // Output: "Hello world"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func FirstToUpper(original string) string {

	// Avoid extra work if string is empty
	if len(original) == 0 {
		return original
	}

	// Fast-path for single character strings
	if len(original) == 1 {
		return strings.ToUpper(original)
	}

	// Decode and uppercase the first rune to support multibyte characters
	r, size := utf8.DecodeRuneInString(original)
	r = unicode.ToUpper(r)

	var b strings.Builder
	b.Grow(len(original))
	b.WriteRune(r)
	b.WriteString(original[size:])

	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func FormalName(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			unicode.IsDigit(r) ||
			r == '-' || r == '\'' || r == ',' || r == '.' ||
			unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func IPAddress(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == ':' || r == '.' {
			b.WriteRune(r)
		}
	}
	ip := net.ParseIP(b.String())
	if ip == nil {
		return ""
	}

	return ip.String()
}

// Numeric returns a string containing only numeric characters (0-9) from the input.
// All non-digit characters are removed. This function supports Unicode digit runes
// and is useful for extracting numbers from user input, phone numbers, IDs, or any
// text where only digits should be retained.
//
// Parameters:
//   - original: The input string to be sanitized.
//
// Returns:
//   - A string containing only numeric characters.
//
// Example:
//
//	input := "Phone: 123-456-7890 ext. 42"
//	result := sanitize.Numeric(input)
//	fmt.Println(result) // Output: "123456789042"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Numeric(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// PhoneNumber returns a sanitized string containing only numeric digits and the
// plus sign (+).
//
// This function is useful for normalizing phone numbers by stripping away
// characters like spaces, dashes, parentheses, and extensions while preserving
// any leading international prefix.
//
// Parameters:
//   - original: The input string representing a phone number to be sanitized.
//
// Returns:
//   - A sanitized phone number consisting solely of digits and plus signs.
//
// Example:
//
//	input := "+1 (234) 567-8900"
//	result := sanitize.PhoneNumber(input)
//	fmt.Println(result) // Output: "+12345678900"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func PhoneNumber(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsDigit(r) || r == '+' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// PathName returns a sanitized string suitable for use as a file or directory name.
// It removes any characters that are not ASCII letters (a-z, A-Z), digits (0-9),
// hyphens (-), or underscores (_), ensuring the result is safe for use as a path component
// on most filesystems. This function is useful for normalizing user input, generating
// safe filenames, or cleaning up strings for use in file paths.
//
// Parameters:
//   - original: The input string to be sanitized.
//
// Returns:
//   - A sanitized string containing only valid path name characters.
//
// Example:
//
//	input := "file:name/with*invalid|chars"
//	result := sanitize.PathName(input)
//	fmt.Println(result) // Output: "filenamewithinvalidchars"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func PathName(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		switch {
		case '0' <= r && r <= '9':
			b.WriteRune(r)
		case 'a' <= r && r <= 'z':
			b.WriteRune(r)
		case 'A' <= r && r <= 'Z':
			b.WriteRune(r)
		case r == '-' || r == '_':
			b.WriteRune(r)
		}
	}
	return b.String()
}

// Punctuation returns a sanitized string containing only alphanumeric characters and common punctuation.
// It removes any characters that are not Unicode letters, digits, or standard punctuation marks such as
// hyphens (-), apostrophes ('), double quotes ("), hash (#), ampersand (&), exclamation mark (!),
// question mark (?), comma (,), period (.), or whitespace. This function is useful for cleaning user input,
// preserving readable punctuation in sentences, or preparing text for display where only basic punctuation is allowed.
//
// Parameters:
//   - original: The input string to be sanitized.
//
// Returns:
//   - A sanitized string containing only alphanumeric characters and common punctuation.
//
// Example:
//
//	input := "Hello, World! How's it going? (Good, I hope.) @2024"
//	result := sanitize.Punctuation(input)
//	fmt.Println(result) // Output: "Hello, World! How's it going? Good, I hope."
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Punctuation(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsLetter(r) || unicode.IsDigit(r) ||
			r == '-' || r == '\'' || r == '"' || r == '#' || r == '&' ||
			r == '!' || r == '?' || r == ',' || r == '.' || unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func ScientificNotation(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsDigit(r) || r == '.' || r == 'e' || r == 'E' || r == '+' || r == '-' {
			b.WriteRune(r)
		}
	}
	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Scripts(original string) string {
	return string(scriptRegExp.ReplaceAll([]byte(original), emptySpace))
}

// SingleLine returns a sanitized version of the input string as a single line of text.
// It replaces all carriage returns (`\r`), line feeds (`\n`), tabs (`\t`), vertical tabs (`\v`),
// and form feeds (`\f`) with a single space character, effectively flattening multi-line or
// formatted input into a single line. This is useful for normalizing user input, log entries,
// or any text that should not contain line breaks or special whitespace.
//
// Parameters:
// - original: The input string to be sanitized.
//
// Returns:
// - A single-line string with all line breaks and special whitespace replaced by spaces.
//
// Example:
//
//	input := "This is a\nmulti-line\tstring."
//	result := sanitize.SingleLine(input)
//	fmt.Println(result) // Output: "This is a multi-line string."
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func SingleLine(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		switch r {
		case '\r', '\n', '\t', '\v', '\f':
			b.WriteRune(' ')
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func Time(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsDigit(r) || r == ':' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// URI returns a sanitized string containing only valid URI characters from the input.
// It removes any characters that are not allowed in URIs, including only Unicode letters, digits,
// dashes (-), underscores (_), slashes (/), question marks (?), ampersands (&), equals signs (=),
// hashes (#), and percent signs (%). This function is useful for cleaning user input, query strings,
// or any text that should conform to URI formatting rules.
//
// Parameters:
//   - original: The input string to be sanitized.
//
// Returns:
//   - A sanitized string containing only valid URI characters.
//
// Example:
//
//	input := "Test?=what! &this=that"
//	result := sanitize.URI(input)
//	fmt.Println(result) // Output: "Test?=what&this=that"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func URI(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsLetter(r) || unicode.IsDigit(r) ||
			r == '-' || r == '_' || r == '/' || r == '?' ||
			r == '&' || r == '=' || r == '#' || r == '%' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// URL returns a sanitized, URL-friendly string containing only valid URL characters.
// It removes any characters that are not allowed in URLs, preserving only Unicode letters, digits,
// dashes (-), underscores (_), slashes (/), colons (:), periods (.), commas (,), question marks (?),
// ampersands (&), at signs (@), equals signs (=), hashes (#), and percent signs (%).
// This function is useful for cleaning user input, constructing safe URLs, or normalizing
// strings for use in web addresses, query parameters, or file paths.
//
// Parameters:
//   - original: The input string to be sanitized.
//
// Returns:
//   - A sanitized string containing only valid URL characters.
//
// Example:
//
//	input := "https://Example.com/This/Works?^No&this"
//	result := sanitize.URL(input)
//	fmt.Println(result) // Output: "https://Example.com/This/Works?No&this"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func URL(original string) string {
	var b strings.Builder
	b.Grow(len(original))
	for _, r := range original {
		if unicode.IsLetter(r) || unicode.IsDigit(r) ||
			r == '-' || r == '_' || r == '/' || r == ':' ||
			r == '.' || r == ',' || r == '?' || r == '&' ||
			r == '@' || r == '=' || r == '#' || r == '%' {
			b.WriteRune(r)
		}
	}
	return b.String()
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
//	input := `<?XML version="1.0" encoding="UTF-8"?><note>Something</note>`
//	result := sanitize.XML(input)
//	fmt.Println(result) // Output: "Something"
//
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
func XML(original string) string {
	return HTML(original)
}

// XSS removes known XSS attack strings or script strings.
// This function sanitizes the input string by removing common XSS attack vectors,
// such as script tags, eval functions, and JavaScript protocol handlers.
//
// WARNING: this is NOT a comprehensive XSS prevention solution.
//
// For a more improved approach, use a library like `github.com/microcosm-cc/bluemonday`
//
// import "github.com/microcosm-cc/bluemonday"
//
//	func SafeHTML(unsafe string) string {
//		p := bluemonday.UGCPolicy() // or build your own allow-list
//		return p.Sanitize(unsafe)
//	}
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
// See more usage examples in the `sanitize_example_test.go` file.
// See the benchmarks in the `sanitize_benchmark_test.go` file.
// See the fuzz tests in the `sanitize_fuzz_test.go` file.
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
