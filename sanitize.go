/*
Package sanitize (go-sanitize) implements a simple library of various sanitation methods for data transformation.

If you have any suggestions or comments, please feel free to open an issue on this project's GitHub page.

Author: MrZ
*/
package sanitize

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

//Set all the regular expressions
var (
	alphaNumericRegExp           = regexp.MustCompile(`[^a-zA-Z0-9]`)                                                             //Alpha numeric
	alphaNumericWithSpacesRegExp = regexp.MustCompile(`[^a-zA-Z0-9\s]`)                                                           //Alpha numeric (with spaces)
	alphaRegExp                  = regexp.MustCompile(`[^a-zA-Z]`)                                                                //Alpha characters
	alphaWithSpacesRegExp        = regexp.MustCompile(`[^a-zA-Z\s]`)                                                              //Alpha characters (with spaces)
	bitcoinRegExp                = regexp.MustCompile(`[^a-km-zA-HJ-NP-Z1-9]`)                                                    //Bitcoin address accepted characters
	bitcoinCashAddrRegExp        = regexp.MustCompile(`[^ac-hj-np-zAC-HJ-NP-Z02-9]`)                                              //Bitcoin `cashaddr` address accepted characters
	decimalRegExp                = regexp.MustCompile(`[^0-9.-]`)                                                                 //Decimals (positive and negative)
	domainRegExp                 = regexp.MustCompile(`[^a-zA-Z0-9-.]`)                                                           //Domain accepted characters
	emailRegExp                  = regexp.MustCompile(`[^a-zA-Z0-9-_.@+]`)                                                        //Email address characters
	formalNameRegExp             = regexp.MustCompile(`[^a-zA-Z0-9-',.\s]`)                                                       //Characters recognized in surnames and proper names
	htmlRegExp                   = regexp.MustCompile(`(?i)<[^>]*>`)                                                              //HTML/XML tags or any alligator open/close tags
	ipAddressRegExp              = regexp.MustCompile(`[^a-zA-Z0-9:.]`)                                                           //IPV4 and IPV6 characters only
	numericRegExp                = regexp.MustCompile(`[^0-9]`)                                                                   //Numbers only
	pathNameRegExp               = regexp.MustCompile(`[^a-zA-Z0-9-_]`)                                                           //Path name (file name, seo)
	punctuationRegExp            = regexp.MustCompile(`[^a-zA-Z0-9-'"#&!?,.\s]+`)                                                 //Standard accepted punctuation characters
	scriptRegExp                 = regexp.MustCompile(`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`) //Scripts and embeds
	singleLineRegExp             = regexp.MustCompile(`(\r)|(\n)|(\t)|(\v)|(\f)`)                                                 //Carriage returns, line feeds, tabs, for single line transition
	timeRegExp                   = regexp.MustCompile(`[^0-9:]`)                                                                  //Time allowed characters
	uriRegExp                    = regexp.MustCompile(`[^a-zA-Z0-9-_/?&=#%]`)                                                     //URI allowed characters
	urlRegExp                    = regexp.MustCompile(`[^a-zA-Z0-9-_/:.?&=#%]`)                                                   //URL allowed characters
)

// Alpha returns only alpha characters. Set the parameter spaces to true if you want to allow space characters. Valid characters are a-z and A-Z.
//  View examples: sanitize_test.go
func Alpha(original string, spaces bool) string {

	// Leave white spaces?
	if spaces {
		return string(alphaWithSpacesRegExp.ReplaceAll([]byte(original), []byte("")))
	}

	// No spaces
	return string(alphaRegExp.ReplaceAll([]byte(original), []byte("")))
}

// AlphaNumeric returns only alphanumeric characters. Set the parameter spaces to true if you want to allow space characters. Valid characters are a-z, A-Z and 0-9.
//  View examples: sanitize_test.go
func AlphaNumeric(original string, spaces bool) string {

	// Leave white spaces?
	if spaces {
		return string(alphaNumericWithSpacesRegExp.ReplaceAll([]byte(original), []byte("")))
	}

	// No spaces
	return string(alphaNumericRegExp.ReplaceAll([]byte(original), []byte("")))
}

// BitcoinAddress returns sanitized value for bitcoin address
//  View examples: sanitize_test.go
func BitcoinAddress(original string) string {
	return string(bitcoinRegExp.ReplaceAll([]byte(original), []byte("")))
}

// BitcoinCashAddress returns sanitized value for bitcoin `cashaddr` address (https://www.bitcoinabc.org/2018-01-14-CashAddr/)
//  View examples: sanitize_test.go
func BitcoinCashAddress(original string) string {
	return string(bitcoinCashAddrRegExp.ReplaceAll([]byte(original), []byte("")))
}

// Custom uses a custom regex string and returns the sanitized result. This is used for any additional regex that this package does not contain.
//  View examples: sanitize_test.go
func Custom(original string, regExp string) string {

	// Try to compile (it will panic if its wrong!)
	compiledRegExp := regexp.MustCompile(regExp)

	// Return the processed string
	return string(compiledRegExp.ReplaceAll([]byte(original), []byte("")))
}

// Decimal returns sanitized decimal/float values in either positive or negative.
//  View examples: sanitize_test.go
func Decimal(original string) string {
	return string(decimalRegExp.ReplaceAll([]byte(original), []byte("")))
}

// Domain returns a proper hostname / domain name. Preserve case is to flag keeping the case versus forcing to lowercase. Use the removeWww flag to strip the www sub-domain. This method returns an error if parse critically fails.
//  View examples: sanitize_test.go
func Domain(original string, preserveCase bool, removeWww bool) (string, error) {

	// Try to see if we have a host
	if len(original) == 0 {
		return original, nil
	}

	// Missing http?
	if !strings.Contains(original, "http") {
		original = "http://" + original
	}

	// Try to parse the url
	u, err := url.Parse(original)
	if err != nil {
		return original, err
	}

	// Try to see if we have a host
	if len(u.Host) == 0 {
		return original, fmt.Errorf("unable to parse domain: %s", original)
	}

	// Remove leading www.
	if removeWww {
		u.Host = strings.Replace(u.Host, "www.", "", -1)
	}

	// Keeps the exact case of the original input string
	if preserveCase {
		return string(domainRegExp.ReplaceAll([]byte(u.Host), []byte(""))), nil
	}

	// Generally all domains should be uniform and lowercase
	return string(domainRegExp.ReplaceAll([]byte(strings.ToLower(u.Host)), []byte(""))), nil
}

// Email returns a sanitized email address string. Email addresses are forced to lowercase and removes any mail-to prefixes.
//  View examples: sanitize_test.go
func Email(original string, preserveCase bool) string {

	// Leave the email address in it's original case
	if preserveCase {
		return string(emailRegExp.ReplaceAll([]byte(strings.Replace(original, "mailto:", "", -1)), []byte("")))
	}

	// Standard is forced to lowercase
	return string(emailRegExp.ReplaceAll([]byte(strings.ToLower(strings.Replace(original, "mailto:", "", -1))), []byte("")))
}

// FirstToUpper overwrites the first letter as an uppercase letter and preserves the rest of the string.
//  View examples: sanitize_test.go
func FirstToUpper(original string) string {

	// Handle empty and 1 character strings
	if len(original) < 2 {
		return strings.ToUpper(original)
	}

	runes := []rune(original)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// FormalName returns a formal name or surname (for First, Middle and Last)
//  View examples: sanitize_test.go
func FormalName(original string) string {
	return string(formalNameRegExp.ReplaceAll([]byte(original), []byte("")))
}

// HTML returns a string without any <HTML> tags.
//  View examples: sanitize_test.go
func HTML(original string) string {
	return string(htmlRegExp.ReplaceAll([]byte(original), []byte("")))
}

// IPAddress returns an ip address for both ipv4 and ipv6 formats.
//  View examples: sanitize_test.go
func IPAddress(original string) string {
	// Parse the IP - Remove any invalid characters first
	ipAddress := net.ParseIP(string(ipAddressRegExp.ReplaceAll([]byte(original), []byte(""))))
	if ipAddress == nil {
		return ""
	}

	return ipAddress.String()
}

// Numeric returns numbers only.
//  View examples: sanitize_test.go
func Numeric(original string) string {
	return string(numericRegExp.ReplaceAll([]byte(original), []byte("")))
}

// PathName returns a formatted path compliant name.
//  View examples: sanitize_test.go
func PathName(original string) string {
	return string(pathNameRegExp.ReplaceAll([]byte(original), []byte("")))
}

// Punctuation returns a string with basic punctuation preserved.
//  View examples: sanitize_test.go
func Punctuation(original string) string {
	return string(punctuationRegExp.ReplaceAll([]byte(original), []byte("")))
}

// Scripts removes all scripts, iframes and embeds tags from string.
//  View examples: sanitize_test.go
func Scripts(original string) string {
	return string(scriptRegExp.ReplaceAll([]byte(original), []byte("")))
}

// SingleLine returns a single line string, removes all carriage returns.
//  View examples: sanitize_test.go
func SingleLine(original string) string {
	return singleLineRegExp.ReplaceAllString(original, " ")
}

// Time returns just the time part of the string.
//  View examples: sanitize_test.go
func Time(original string) string {
	return string(timeRegExp.ReplaceAll([]byte(original), []byte("")))
}

// URI returns allowed URI characters only.
//  View examples: sanitize_test.go
func URI(original string) string {
	return string(uriRegExp.ReplaceAll([]byte(original), []byte("")))
}

// URL returns a formatted url friendly string.
//  View examples: sanitize_test.go
func URL(original string) string {
	return string(urlRegExp.ReplaceAll([]byte(original), []byte("")))
}

// XML returns a string without any <XML> tags - alias of HTML.
//  View examples: sanitize_test.go
func XML(original string) string {
	return HTML(original)
}

// XSS removes known XSS attack strings or script strings.
//  View examples: sanitize_test.go
func XSS(original string) string {
	original = strings.Replace(original, "<script", "", -1)
	original = strings.Replace(original, "script>", "", -1)
	original = strings.Replace(original, "eval(", "", -1)
	original = strings.Replace(original, "eval&#40;", "", -1)
	original = strings.Replace(original, "javascript:", "", -1)
	original = strings.Replace(original, "javascript&#58;", "", -1)
	original = strings.Replace(original, "fromCharCode", "", -1)
	original = strings.Replace(original, "&#62;", "", -1)
	original = strings.Replace(original, "&#60;", "", -1)
	original = strings.Replace(original, "&lt;", "", -1)
	original = strings.Replace(original, "&rt;", "", -1)
	return original
}
