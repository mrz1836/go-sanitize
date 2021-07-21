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

	// See more in: sanitize_test.go
}
