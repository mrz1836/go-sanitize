/*
Package goSanitize implements a simple library of various sanitation methods for data transformation.
*/
package goSanitize

import (
	"net"
	"regexp"
	"strings"
	"unicode"
)

// Set all the regular expressions
var (
	alphaNumericRegExp   = regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	alphaRegExp          = regexp.MustCompile(`[^a-zA-Z\s]`)
	decimalRegExp        = regexp.MustCompile(`[^0-9.-]`)
	domainRegExp         = regexp.MustCompile(`[^a-zA-Z0-9-.]`)
	driversLicenseRegExp = regexp.MustCompile(`[^a-zA-Z0-9]`)
	emailRegExp        = regexp.MustCompile(`[^a-zA-Z0-9-_.@]`)
	htmlOpenRegExp     = regexp.MustCompile(`(?i)<[^>]*>`)
	nameFormalRegExp   = regexp.MustCompile(`[^a-zA-Z0-9-',.\s]`)
	numericRegExp      = regexp.MustCompile(`[^0-9]`)
	punctuationRegExp  = regexp.MustCompile(`[^a-zA-Z0-9-'"#&!?,.\s]+`)
	scriptRegExp       = regexp.MustCompile(`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`) //`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`
	seoRegExp          = regexp.MustCompile(`[^a-zA-Z0-9-_]`)
	singleLineRegExp   = regexp.MustCompile(`\r?\n`)
	socialNumberRegExp = regexp.MustCompile(`[^0-9-]`)
	timeRegExp         = regexp.MustCompile(`[^0-9:]`)
	unicodeRegExp        = regexp.MustCompile(`[[^:unicode]]`) //`[[^:unicode:]]`
	uriRegExp          = regexp.MustCompile(`[^a-zA-Z0-9-_/?&=%]`)
	urlRegExp          = regexp.MustCompile(`[^a-zA-Z0-9-_/:.?&=#%]`)
)

//Unicode returns unicode characters only
func Unicode(original string) string {
	bytes := unicodeRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//Alpha characters only
func Alpha(original string) string {
	bytes := alphaRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//AlphaNumeric alpha and numeric characters only
func AlphaNumeric(original string) string {
	bytes := alphaNumericRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//Decimal floats and decimals
func Decimal(original string) string {
	bytes := decimalRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//Domain only formatting
func Domain(original string) string {
	original = strings.Replace(strings.Replace(strings.Replace(original, "https", "", 3), "www.", "", 3), "http", "", 3)
	bytes := domainRegExp.ReplaceAll([]byte(strings.ToLower(original)), []byte(""))
	return string(bytes)
}

//DriversLicense only formatting (a-z and 0-9)
func DriversLicense(original string) string {
	bytes := driversLicenseRegExp.ReplaceAll([]byte(strings.ToUpper(original)), []byte(""))
	return string(bytes)
}

//Email only formatting
func Email(original string) string {
	original = strings.Replace(original, "mailto:", "", 3)
	bytes := emailRegExp.ReplaceAll([]byte(strings.ToLower(original)), []byte(""))
	return string(bytes)
}

//FirstToUpper overwrites the first letter as an uppercase letter
func FirstToUpper(original string) string {
	return makeFirstUpperCase(original)
}

//HTML removes all basic html tags that we accept
func HTML(original string) string {
	bytes := htmlOpenRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//IPAddress format as ip address for both ipv4 and ipv6
func IPAddress(original string) string {
	ipAddress := net.ParseIP(strings.TrimSpace(original))

	if ipAddress == nil {
		return ""
	}

	return ipAddress.String()
}

//NameFormal is (for First, Middle and Last)
func NameFormal(original string) string {
	bytes := nameFormalRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//Numeric numbers only
func Numeric(original string) string {
	bytes := numericRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//Punctuation used for generic sentences
func Punctuation(original string) string {
	bytes := punctuationRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//Scripts removes all script / iframes / embeds tags
func Scripts(original string) string {
	bytes := scriptRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//Seo returns URL Friendly
func Seo(original string) string {
	bytes := seoRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//SocialSecurityNumber formats social - xxx-xx-xxxx
func SocialSecurityNumber(original string) string {
	//Replace all unwanted characters with and empty string
	bytes := socialNumberRegExp.ReplaceAll([]byte(original), []byte(""))

	//return the clean string
	return string(bytes)
}

//SingleLine used for forcing to a single line
func SingleLine(original string) string {
	return singleLineRegExp.ReplaceAllString(original, " ")
}

//Time only (0-9 :)
func Time(original string) string {
	bytes := timeRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//URI allowed URI characters
func URI(original string) string {
	bytes := uriRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//URL format as URL
func URL(original string) string {
	bytes := urlRegExp.ReplaceAll([]byte(original), []byte(""))
	return string(bytes)
}

//XSS remove XSS attack strings
func XSS(original string) string {
	//Remove all XSS attacks
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
	original = strings.Replace(original, "&lt;", "", -1)

	//return the clean string
	return original
}

//makeFirstUpperCase upper cases the first letter of the string
func makeFirstUpperCase(s string) string {

	// Handle empty and 1 character strings
	if len(s) < 2 {
		return strings.ToUpper(s)
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}