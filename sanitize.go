/*
Package gosanitize implements a simple library of various sanitation methods for data transformation.
*/
package gosanitize

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

// Set all the regular expressions
var (
	alphaNumericRegExp           = regexp.MustCompile(`[^a-zA-Z0-9]`)   //Alpha numeric
	alphaNumericWithSpacesRegExp = regexp.MustCompile(`[^a-zA-Z0-9\s]`) //Alpha numeric (with spaces)
	alphaRegExp                  = regexp.MustCompile(`[^a-zA-Z]`)      //Alpha characters
	alphaWithSpacesRegExp        = regexp.MustCompile(`[^a-zA-Z\s]`)    //Alpha characters (with spaces)
	decimalRegExp                = regexp.MustCompile(`[^0-9.-]`)       //Decimals (positive and negative)
	domainRegExp                 = regexp.MustCompile(`[^a-zA-Z0-9-.]`) //Domain accepted characters
	emailRegExp                  = regexp.MustCompile(`[^a-zA-Z0-9-_.@]`)
	htmlOpenRegExp               = regexp.MustCompile(`(?i)<[^>]*>`)
	nameFormalRegExp             = regexp.MustCompile(`[^a-zA-Z0-9-',.\s]`)
	numericRegExp                = regexp.MustCompile(`[^0-9]`)
	punctuationRegExp            = regexp.MustCompile(`[^a-zA-Z0-9-'"#&!?,.\s]+`)
	scriptRegExp                 = regexp.MustCompile(`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`) //`(?i)<(script|iframe|embed|object)[^>]*>.*</(script|iframe|embed|object)>`
	seoRegExp                    = regexp.MustCompile(`[^a-zA-Z0-9-_]`)
	singleLineRegExp             = regexp.MustCompile(`\r?\n`)
	socialNumberRegExp           = regexp.MustCompile(`[^0-9-]`)
	timeRegExp                   = regexp.MustCompile(`[^0-9:]`)
	unicodeRegExp                = regexp.MustCompile(`[[^:unicode]]`) //`[[^:unicode:]]`
	uriRegExp                    = regexp.MustCompile(`[^a-zA-Z0-9-_/?&=%]`)
	urlRegExp                    = regexp.MustCompile(`[^a-zA-Z0-9-_/:.?&=#%]`)
)

//Alpha returns only alpha characters (flag for spaces)
func Alpha(original string, spaces bool) string {

	//Leave white spaces?
	if spaces {
		return string(alphaWithSpacesRegExp.ReplaceAll([]byte(original), []byte("")))
	}

	//No spaces
	return string(alphaRegExp.ReplaceAll([]byte(original), []byte("")))
}

//AlphaNumeric returns alpha and numeric characters only (flag for spaces)
func AlphaNumeric(original string, spaces bool) string {

	//Leave white spaces?
	if spaces {
		return string(alphaNumericWithSpacesRegExp.ReplaceAll([]byte(original), []byte("")))
	}

	//No spaces
	return string(alphaNumericRegExp.ReplaceAll([]byte(original), []byte("")))
}

//Decimal returns decimal values (positive and negative)
func Decimal(original string) string {
	return string(decimalRegExp.ReplaceAll([]byte(original), []byte("")))
}

//Domain returns a proper domain name (example.com - lowercase)
func Domain(original string, preserveCase bool, removeWww bool) (string, error) {

	//Try to see if we have a host
	if len(original) == 0 {
		return original, fmt.Errorf("domain value is empty")
	}

	//Missing http?
	if !strings.Contains(original, "http") {
		original = "http://" + original
	}

	//Try to parse the url
	u, err := url.Parse(original)
	if err != nil {
		return original, err
	}

	//Try to see if we have a host
	if len(u.Host) == 0 {
		return original, fmt.Errorf("unable to parse domain: %s", original)
	}

	//Remove leading www.
	if removeWww {
		u.Host = strings.Replace(u.Host, "www.", "", -1)
	}

	//Keeps the exact case of the original input string
	if preserveCase {
		return string(domainRegExp.ReplaceAll([]byte(u.Host), []byte(""))), nil
	}

	//Generally all domains should be uniform and lowercase
	return string(domainRegExp.ReplaceAll([]byte(strings.ToLower(u.Host)), []byte(""))), nil
}

//Email only formatting
func Email(original string) string {
	original = strings.Replace(original, "mailto:", "", 3)
	return string(emailRegExp.ReplaceAll([]byte(strings.ToLower(original)), []byte("")))
}

//FirstToUpper overwrites the first letter as an uppercase letter
func FirstToUpper(original string) string {
	return makeFirstUpperCase(original)
}

//HTML removes all basic html tags that we accept
func HTML(original string) string {
	return string(htmlOpenRegExp.ReplaceAll([]byte(original), []byte("")))
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
	return string(nameFormalRegExp.ReplaceAll([]byte(original), []byte("")))
}

//Numeric numbers only
func Numeric(original string) string {
	return string(numericRegExp.ReplaceAll([]byte(original), []byte("")))
}

//Punctuation used for generic sentences
func Punctuation(original string) string {
	return string(punctuationRegExp.ReplaceAll([]byte(original), []byte("")))
}

//Scripts removes all script / iframes / embeds tags
func Scripts(original string) string {
	return string(scriptRegExp.ReplaceAll([]byte(original), []byte("")))
}

//Seo returns URL Friendly
func Seo(original string) string {
	return string(seoRegExp.ReplaceAll([]byte(original), []byte("")))
}

//SocialSecurityNumber formats social - xxx-xx-xxxx
func SocialSecurityNumber(original string) string {
	return string(socialNumberRegExp.ReplaceAll([]byte(original), []byte("")))
}

//SingleLine used for forcing to a single line
func SingleLine(original string) string {
	return singleLineRegExp.ReplaceAllString(original, " ")
}

//Time only (0-9 :)
func Time(original string) string {
	return string(timeRegExp.ReplaceAll([]byte(original), []byte("")))
}

//Unicode returns unicode characters only
func Unicode(original string) string {
	return string(unicodeRegExp.ReplaceAll([]byte(original), []byte("")))
}

//URI allowed URI characters
func URI(original string) string {
	return string(uriRegExp.ReplaceAll([]byte(original), []byte("")))
}

//URL format as URL
func URL(original string) string {
	return string(urlRegExp.ReplaceAll([]byte(original), []byte("")))
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
