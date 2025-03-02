package main

import (
	"fmt"

	"github.com/mrz1836/go-sanitize"
)

func main() {

	testString := "1-> A simple test string!"

	// Remove spaces, numbers and symbols
	formattedString := sanitize.Alpha(testString, false)
	fmt.Printf("removed numbers, spaces & symbols: %s\n", formattedString)

	// Preserve spaces
	formattedString = sanitize.Alpha(testString, true)
	fmt.Printf("preserved spaces: %s\n", formattedString)

	// Only numbers
	formattedString = sanitize.Numeric(testString)
	fmt.Printf("only numbers: %s\n", formattedString)

	// Sanitize a URL
	urlString := "https://example.com/This/Works?^No&this"
	sanitizedURL := sanitize.URL(urlString)
	fmt.Printf("sanitized URL: %s\n", sanitizedURL)

	// Sanitize an XML string
	xmlString := "<xml>This?</xml>"
	sanitizedXML := sanitize.XML(xmlString)
	fmt.Printf("sanitized XML: %s\n", sanitizedXML)

	// Sanitize a URI
	uriString := "/This/Works?^No&this"
	sanitizedURI := sanitize.URI(uriString)
	fmt.Printf("sanitized URI: %s\n", sanitizedURI)

	// Sanitize a time string
	timeString := "t00:00:00d -EST"
	sanitizedTime := sanitize.Time(timeString)
	fmt.Printf("sanitized time: %s\n", sanitizedTime)

	// Sanitize a string for XSS
	xssString := "<script>alert('test');</script>"
	sanitizedXSS := sanitize.XSS(xssString)
	fmt.Printf("sanitized XSS: %s\n", sanitizedXSS)

	// See more in: sanitize_test.go
}
