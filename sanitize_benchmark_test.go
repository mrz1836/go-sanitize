package sanitize_test

import (
	"regexp"
	"testing"

	"github.com/mrz1836/go-sanitize"
)

// BenchmarkAlpha benchmarks the Alpha method
func BenchmarkAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Alpha("This is the test string.", false)
	}
}

// BenchmarkAlphaNumeric benchmarks the AlphaNumeric method
func BenchmarkAlphaNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.AlphaNumeric("This is the test string 12345.", false)
	}
}

// BenchmarkAlphaNumeric_WithSpaces benchmarks the AlphaNumeric method
func BenchmarkAlphaNumeric_WithSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.AlphaNumeric("This is the test string 12345.", true)
	}
}

// BenchmarkAlpha_WithSpaces benchmarks the Alpha method
func BenchmarkAlpha_WithSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Alpha("This is the test string.", true)
	}
}

// BenchmarkBitcoinAddress benchmarks the BitcoinAddress method
func BenchmarkBitcoinAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.BitcoinAddress("1K6c7LGpdB8LwoGNVfG51dRV9UUEijbrWs")
	}
}

// BenchmarkBitcoinCashAddress benchmarks the BitcoinCashAddress() method
func BenchmarkBitcoinCashAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.BitcoinCashAddress("qze7yy2au5vuznvn8lzj5y0j5t066vhs75e3m0eptz")
	}
}

// BenchmarkCustom benchmarks the Custom method
func BenchmarkCustom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Custom("This is the test string 12345.", `[^a-zA-Z0-9]`)
	}
}

// BenchmarkCustomCompiled benchmarks the CustomCompiled method
func BenchmarkCustomCompiled(b *testing.B) {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	for i := 0; i < b.N; i++ {
		_ = sanitize.CustomCompiled("This is the test string 12345.", re)
	}
}

// BenchmarkDecimal benchmarks the Decimal method
func BenchmarkDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Decimal("String: -123.12345")
	}
}

// BenchmarkDomain benchmarks the Domain method
func BenchmarkDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sanitize.Domain("https://Example.COM/?param=value", false, false)
	}
}

// BenchmarkDomain_PreserveCase benchmarks the Domain method
func BenchmarkDomain_PreserveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sanitize.Domain("https://Example.COM/?param=value", true, false)
	}
}

// BenchmarkDomain_RemoveWww benchmarks the Domain method
func BenchmarkDomain_RemoveWww(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sanitize.Domain("https://Example.COM/?param=value", false, true)
	}
}

// BenchmarkEmail benchmarks the Email method
func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Email("mailto:Person@Example.COM ", false)
	}
}

// BenchmarkEmail_PreserveCase benchmarks the Email method
func BenchmarkEmail_PreserveCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Email("mailto:Person@Example.COM ", true)
	}
}

// BenchmarkFirstToUpper benchmarks the FirstToUpper method
func BenchmarkFirstToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.FirstToUpper("make this upper")
	}
}

// BenchmarkFormalName benchmarks the FormalName method
func BenchmarkFormalName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.FormalName("John McDonald Jr.")
	}
}

// BenchmarkHTML benchmarks the HTML method
func BenchmarkHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.HTML("<html><b>Test This!</b></html>")
	}
}

// BenchmarkIPAddress benchmarks the IPAddress method
func BenchmarkIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.IPAddress(" 192.168.0.1 ")
	}
}

// BenchmarkIPAddress_V6 benchmarks the IPAddress method
func BenchmarkIPAddress_IPV6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.IPAddress(" 2602:305:bceb:1bd0:44ef:fedb:4f8f:da4f ")
	}
}

// BenchmarkNumeric benchmarks the numeric method
func BenchmarkNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Numeric(" 192.168.0.1 ")
	}
}

// BenchmarkPathName benchmarks the PathName method
func BenchmarkPathName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.PathName("/This-Path-Name_Works-/")
	}
}

// BenchmarkPunctuation benchmarks the Punctuation method
func BenchmarkPunctuation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Punctuation("Does this work? They're doing it?")
	}
}

// BenchmarkScientificNotation benchmarks the ScientificNotation method
func BenchmarkScientificNotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.ScientificNotation("String: -1.096e-3")
	}
}

// BenchmarkScripts benchmarks the Scripts method
func BenchmarkScripts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Scripts("<script>$(){ var remove='me'; }</script>")
	}
}

// BenchmarkSingleLine benchmarks the SingleLine method
func BenchmarkSingleLine(b *testing.B) {
	testString := `This line
That Line
Another Line`
	for i := 0; i < b.N; i++ {
		_ = sanitize.SingleLine(testString)
	}
}

// BenchmarkTime benchmarks the Time method
func BenchmarkTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.Time("Time is 05:10:23")
	}
}

// BenchmarkURI benchmarks the URI method
func BenchmarkURI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.URI("/Test/This/Url/?param=value")
	}
}

// BenchmarkURL benchmarks the URL method
func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.URL("/Test/This/Url/?param=value")
	}
}

// BenchmarkXML benchmarks the XML method
func BenchmarkXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.XML("<xml>Test This!</xml>")
	}
}

// BenchmarkXSS benchmarks the XSS method
func BenchmarkXSS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize.XSS("<script>Test This!</script>")
	}
}
