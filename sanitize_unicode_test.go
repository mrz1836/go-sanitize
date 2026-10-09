package sanitize_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrz1836/go-sanitize"
)

// unicodeCase pins what each name sanitizer returns for one input
type unicodeCase struct {
	name               string
	input              string
	alpha              string // Alpha(input, false)
	alphaSpaces        string // Alpha(input, true)
	alphaNumeric       string // AlphaNumeric(input, false)
	alphaNumericSpaces string // AlphaNumeric(input, true)
	formalName         string // FormalName(input)
}

// unicodeCorpus returns names and text in many scripts, forms, and edge cases, each with the exact
// output of Alpha, AlphaNumeric, and FormalName. A precomposed row holds NFC text and a decomposed
// row the same text in NFD. Every non-ASCII character is written as an escape, with an ASCII
// romanization in a comment, so no editor can recompose a fixture. A change to any of the three
// functions shows up as the list of rows whose output changed.
func unicodeCorpus() []unicodeCase {
	return []unicodeCase{
		{
			name:               "empty input",
			input:              "",
			alpha:              "",
			alphaSpaces:        "",
			alphaNumeric:       "",
			alphaNumericSpaces: "",
			formalName:         "",
		},
		{
			name:               "ascii name with punctuation",
			input:              "Mary O'Brien-Smith, Jr.",
			alpha:              "MaryOBrienSmithJr",
			alphaSpaces:        "Mary OBrienSmith Jr",
			alphaNumeric:       "MaryOBrienSmithJr",
			alphaNumericSpaces: "Mary OBrienSmith Jr",
			formalName:         "Mary O'Brien-Smith, Jr.",
		},
		{
			name:               "ascii digits and symbols",
			input:              "Box 42 & Q: #9!",
			alpha:              "BoxQ",
			alphaSpaces:        "Box   Q ",
			alphaNumeric:       "Box42Q9",
			alphaNumericSpaces: "Box 42  Q 9",
			formalName:         "Box 42  Q 9",
		},
		{
			name:               "ascii punctuation only",
			input:              "!@#$%^&*()_+-=[]{};:'\",.<>/?\\|`~",
			alpha:              "",
			alphaSpaces:        "",
			alphaNumeric:       "",
			alphaNumericSpaces: "",
			formalName:         "-',.",
		},
		{
			name:               "html and script tags",
			input:              "<b>Bob</b><script>alert(1)</script>",
			alpha:              "bBobbscriptalertscript",
			alphaSpaces:        "bBobbscriptalertscript",
			alphaNumeric:       "bBobbscriptalert1script",
			alphaNumericSpaces: "bBobbscriptalert1script",
			formalName:         "bBobbscriptalert1script",
		},
		{
			name:               "sql injection attempt",
			input:              "Robert'); DROP TABLE Students;--",
			alpha:              "RobertDROPTABLEStudents",
			alphaSpaces:        "Robert DROP TABLE Students",
			alphaNumeric:       "RobertDROPTABLEStudents",
			alphaNumericSpaces: "Robert DROP TABLE Students",
			formalName:         "Robert' DROP TABLE Students--",
		},
		{
			name:               "tabs newlines and carriage returns",
			input:              "Cleo\tWard\nHale\r\nIII",
			alpha:              "CleoWardHaleIII",
			alphaSpaces:        "CleoWardHaleIII",
			alphaNumeric:       "CleoWardHaleIII",
			alphaNumericSpaces: "CleoWardHaleIII",
			formalName:         "Cleo\tWard\nHale\r\nIII",
		},
		{
			name:               "leading and trailing spaces",
			input:              "  Dana Fox  ",
			alpha:              "DanaFox",
			alphaSpaces:        "  Dana Fox  ",
			alphaNumeric:       "DanaFox",
			alphaNumericSpaces: "  Dana Fox  ",
			formalName:         "  Dana Fox  ",
		},
		{
			name:               "spaces only",
			input:              "   ",
			alpha:              "",
			alphaSpaces:        "   ",
			alphaNumeric:       "",
			alphaNumericSpaces: "   ",
			formalName:         "   ",
		},
		{
			name:               "digits only",
			input:              "0123456789",
			alpha:              "",
			alphaSpaces:        "",
			alphaNumeric:       "0123456789",
			alphaNumericSpaces: "0123456789",
			formalName:         "0123456789",
		},
		{
			name:               "no-break space",
			input:              "Eve\u00a0Hart",
			alpha:              "EveHart",
			alphaSpaces:        "EveHart",
			alphaNumeric:       "EveHart",
			alphaNumericSpaces: "EveHart",
			formalName:         "Eve\u00a0Hart",
		},
		{
			name:               "ideographic space",
			input:              "Finn\u3000Gray",
			alpha:              "FinnGray",
			alphaSpaces:        "FinnGray",
			alphaNumeric:       "FinnGray",
			alphaNumericSpaces: "FinnGray",
			formalName:         "Finn\u3000Gray",
		},
		{
			name:               "em space",
			input:              "Gail\u2003Hunt",
			alpha:              "GailHunt",
			alphaSpaces:        "GailHunt",
			alphaNumeric:       "GailHunt",
			alphaNumericSpaces: "GailHunt",
			formalName:         "Gail\u2003Hunt",
		},
		{
			name:               "line separator",
			input:              "Hugo\u2028Ives",
			alpha:              "HugoIves",
			alphaSpaces:        "HugoIves",
			alphaNumeric:       "HugoIves",
			alphaNumericSpaces: "HugoIves",
			formalName:         "Hugo\u2028Ives",
		},
		{
			name:               "next line control",
			input:              "Iris\u0085Jett",
			alpha:              "IrisJett",
			alphaSpaces:        "IrisJett",
			alphaNumeric:       "IrisJett",
			alphaNumericSpaces: "IrisJett",
			formalName:         "Iris\u0085Jett",
		},
		{
			name:               "right single quotation mark as apostrophe",
			input:              "D\u2019Angelo O\u2019Neil",
			alpha:              "DAngeloONeil",
			alphaSpaces:        "DAngelo ONeil",
			alphaNumeric:       "DAngeloONeil",
			alphaNumericSpaces: "DAngelo ONeil",
			formalName:         "D\u2019Angelo O\u2019Neil",
		},
		{
			name:               "single quotation marks around a word",
			input:              "\u2018Bob\u2019",
			alpha:              "Bob",
			alphaSpaces:        "Bob",
			alphaNumeric:       "Bob",
			alphaNumericSpaces: "Bob",
			formalName:         "Bob\u2019",
		},
		{
			name:               "okina and modifier letter apostrophe",
			input:              "Hawai\u02bbi Ma\u02bcal",
			alpha:              "Hawai\u02bbiMa\u02bcal",
			alphaSpaces:        "Hawai\u02bbi Ma\u02bcal",
			alphaNumeric:       "Hawai\u02bbiMa\u02bcal",
			alphaNumericSpaces: "Hawai\u02bbi Ma\u02bcal",
			formalName:         "Hawai\u02bbi Ma\u02bcal",
		},
		{
			name:               "acute and grave accent symbols as apostrophes",
			input:              "O\u00b4Hara O`Hara",
			alpha:              "OHaraOHara",
			alphaSpaces:        "OHara OHara",
			alphaNumeric:       "OHaraOHara",
			alphaNumericSpaces: "OHara OHara",
			formalName:         "OHara OHara",
		},
		{
			name:               "prime and fullwidth apostrophe",
			input:              "O\u2032Toole O\uff07Toole",
			alpha:              "OTooleOToole",
			alphaSpaces:        "OToole OToole",
			alphaNumeric:       "OTooleOToole",
			alphaNumericSpaces: "OToole OToole",
			formalName:         "OToole OToole",
		},
		{
			name:               "double quotation marks",
			input:              "\u201cKit\u201d \"Lane\"",
			alpha:              "KitLane",
			alphaSpaces:        "Kit Lane",
			alphaNumeric:       "KitLane",
			alphaNumericSpaces: "Kit Lane",
			formalName:         "Kit Lane",
		},
		{
			name:               "unicode hyphen and dashes",
			input:              "Mary\u2010Jo Lee\u2013Ann Bo\u2014Ko",
			alpha:              "MaryJoLeeAnnBoKo",
			alphaSpaces:        "MaryJo LeeAnn BoKo",
			alphaNumeric:       "MaryJoLeeAnnBoKo",
			alphaNumericSpaces: "MaryJo LeeAnn BoKo",
			formalName:         "MaryJo LeeAnn BoKo",
		},
		{
			name:               "non-breaking hyphen and minus sign",
			input:              "Jean\u2011Luc 5\u2212",
			alpha:              "JeanLuc",
			alphaSpaces:        "JeanLuc ",
			alphaNumeric:       "JeanLuc5",
			alphaNumericSpaces: "JeanLuc 5",
			formalName:         "JeanLuc 5",
		},
		{
			name:               "spanish name, precomposed", // Jose Maria Nunez
			input:              "Jos\u00e9 Mar\u00eda N\u00fa\u00f1ez",
			alpha:              "Jos\u00e9Mar\u00edaN\u00fa\u00f1ez",
			alphaSpaces:        "Jos\u00e9 Mar\u00eda N\u00fa\u00f1ez",
			alphaNumeric:       "Jos\u00e9Mar\u00edaN\u00fa\u00f1ez",
			alphaNumericSpaces: "Jos\u00e9 Mar\u00eda N\u00fa\u00f1ez",
			formalName:         "Jos\u00e9 Mar\u00eda N\u00fa\u00f1ez",
		},
		{
			name:               "spanish name, decomposed", // Jose Maria Nunez
			input:              "Jose\u0301 Mari\u0301a Nu\u0301n\u0303ez",
			alpha:              "Jose\u0301Mari\u0301aNu\u0301n\u0303ez",
			alphaSpaces:        "Jose\u0301 Mari\u0301a Nu\u0301n\u0303ez",
			alphaNumeric:       "Jose\u0301Mari\u0301aNu\u0301n\u0303ez",
			alphaNumericSpaces: "Jose\u0301 Mari\u0301a Nu\u0301n\u0303ez",
			formalName:         "Jose\u0301 Mari\u0301a Nu\u0301n\u0303ez",
		},
		{
			name:               "french name, precomposed", // Francois Lefevre Benoit
			input:              "Fran\u00e7ois Lef\u00e8vre Beno\u00eet",
			alpha:              "Fran\u00e7oisLef\u00e8vreBeno\u00eet",
			alphaSpaces:        "Fran\u00e7ois Lef\u00e8vre Beno\u00eet",
			alphaNumeric:       "Fran\u00e7oisLef\u00e8vreBeno\u00eet",
			alphaNumericSpaces: "Fran\u00e7ois Lef\u00e8vre Beno\u00eet",
			formalName:         "Fran\u00e7ois Lef\u00e8vre Beno\u00eet",
		},
		{
			name:               "french name, decomposed", // Francois Lefevre Benoit
			input:              "Franc\u0327ois Lefe\u0300vre Benoi\u0302t",
			alpha:              "Franc\u0327oisLefe\u0300vreBenoi\u0302t",
			alphaSpaces:        "Franc\u0327ois Lefe\u0300vre Benoi\u0302t",
			alphaNumeric:       "Franc\u0327oisLefe\u0300vreBenoi\u0302t",
			alphaNumericSpaces: "Franc\u0327ois Lefe\u0300vre Benoi\u0302t",
			formalName:         "Franc\u0327ois Lefe\u0300vre Benoi\u0302t",
		},
		{
			name:               "german name, precomposed", // Juergen Weiss Oesterreich
			input:              "J\u00fcrgen Wei\u00df \u00d6sterreich",
			alpha:              "J\u00fcrgenWei\u00df\u00d6sterreich",
			alphaSpaces:        "J\u00fcrgen Wei\u00df \u00d6sterreich",
			alphaNumeric:       "J\u00fcrgenWei\u00df\u00d6sterreich",
			alphaNumericSpaces: "J\u00fcrgen Wei\u00df \u00d6sterreich",
			formalName:         "J\u00fcrgen Wei\u00df \u00d6sterreich",
		},
		{
			name:               "german name, decomposed", // Juergen Weiss Oesterreich
			input:              "Ju\u0308rgen Wei\u00df O\u0308sterreich",
			alpha:              "Ju\u0308rgenWei\u00dfO\u0308sterreich",
			alphaSpaces:        "Ju\u0308rgen Wei\u00df O\u0308sterreich",
			alphaNumeric:       "Ju\u0308rgenWei\u00dfO\u0308sterreich",
			alphaNumericSpaces: "Ju\u0308rgen Wei\u00df O\u0308sterreich",
			formalName:         "Ju\u0308rgen Wei\u00df O\u0308sterreich",
		},
		{
			name:               "scandinavian name, precomposed", // Asa Soren Aero
			input:              "\u00c5sa S\u00f8ren \u00c6r\u00f8",
			alpha:              "\u00c5saS\u00f8ren\u00c6r\u00f8",
			alphaSpaces:        "\u00c5sa S\u00f8ren \u00c6r\u00f8",
			alphaNumeric:       "\u00c5saS\u00f8ren\u00c6r\u00f8",
			alphaNumericSpaces: "\u00c5sa S\u00f8ren \u00c6r\u00f8",
			formalName:         "\u00c5sa S\u00f8ren \u00c6r\u00f8",
		},
		{
			name:               "scandinavian name, decomposed", // Asa Soren Aero
			input:              "A\u030asa S\u00f8ren \u00c6r\u00f8",
			alpha:              "A\u030asaS\u00f8ren\u00c6r\u00f8",
			alphaSpaces:        "A\u030asa S\u00f8ren \u00c6r\u00f8",
			alphaNumeric:       "A\u030asaS\u00f8ren\u00c6r\u00f8",
			alphaNumericSpaces: "A\u030asa S\u00f8ren \u00c6r\u00f8",
			formalName:         "A\u030asa S\u00f8ren \u00c6r\u00f8",
		},
		{
			name:               "polish name, precomposed", // Lukasz Zolc
			input:              "\u0141ukasz \u017b\u00f3\u0142\u0107",
			alpha:              "\u0141ukasz\u017b\u00f3\u0142\u0107",
			alphaSpaces:        "\u0141ukasz \u017b\u00f3\u0142\u0107",
			alphaNumeric:       "\u0141ukasz\u017b\u00f3\u0142\u0107",
			alphaNumericSpaces: "\u0141ukasz \u017b\u00f3\u0142\u0107",
			formalName:         "\u0141ukasz \u017b\u00f3\u0142\u0107",
		},
		{
			name:               "polish name, decomposed", // Lukasz Zolc
			input:              "\u0141ukasz Z\u0307o\u0301\u0142c\u0301",
			alpha:              "\u0141ukaszZ\u0307o\u0301\u0142c\u0301",
			alphaSpaces:        "\u0141ukasz Z\u0307o\u0301\u0142c\u0301",
			alphaNumeric:       "\u0141ukaszZ\u0307o\u0301\u0142c\u0301",
			alphaNumericSpaces: "\u0141ukasz Z\u0307o\u0301\u0142c\u0301",
			formalName:         "\u0141ukasz Z\u0307o\u0301\u0142c\u0301",
		},
		{
			name:               "czech name, precomposed", // Antonin Dvorak
			input:              "Anton\u00edn Dvo\u0159\u00e1k",
			alpha:              "Anton\u00ednDvo\u0159\u00e1k",
			alphaSpaces:        "Anton\u00edn Dvo\u0159\u00e1k",
			alphaNumeric:       "Anton\u00ednDvo\u0159\u00e1k",
			alphaNumericSpaces: "Anton\u00edn Dvo\u0159\u00e1k",
			formalName:         "Anton\u00edn Dvo\u0159\u00e1k",
		},
		{
			name:               "czech name, decomposed", // Antonin Dvorak
			input:              "Antoni\u0301n Dvor\u030ca\u0301k",
			alpha:              "Antoni\u0301nDvor\u030ca\u0301k",
			alphaSpaces:        "Antoni\u0301n Dvor\u030ca\u0301k",
			alphaNumeric:       "Antoni\u0301nDvor\u030ca\u0301k",
			alphaNumericSpaces: "Antoni\u0301n Dvor\u030ca\u0301k",
			formalName:         "Antoni\u0301n Dvor\u030ca\u0301k",
		},
		{
			name:               "icelandic name, precomposed", // Thordis Gudmundsdottir
			input:              "\u00de\u00f3rd\u00eds Gu\u00f0mundsd\u00f3ttir",
			alpha:              "\u00de\u00f3rd\u00edsGu\u00f0mundsd\u00f3ttir",
			alphaSpaces:        "\u00de\u00f3rd\u00eds Gu\u00f0mundsd\u00f3ttir",
			alphaNumeric:       "\u00de\u00f3rd\u00edsGu\u00f0mundsd\u00f3ttir",
			alphaNumericSpaces: "\u00de\u00f3rd\u00eds Gu\u00f0mundsd\u00f3ttir",
			formalName:         "\u00de\u00f3rd\u00eds Gu\u00f0mundsd\u00f3ttir",
		},
		{
			name:               "icelandic name, decomposed", // Thordis Gudmundsdottir
			input:              "\u00deo\u0301rdi\u0301s Gu\u00f0mundsdo\u0301ttir",
			alpha:              "\u00deo\u0301rdi\u0301sGu\u00f0mundsdo\u0301ttir",
			alphaSpaces:        "\u00deo\u0301rdi\u0301s Gu\u00f0mundsdo\u0301ttir",
			alphaNumeric:       "\u00deo\u0301rdi\u0301sGu\u00f0mundsdo\u0301ttir",
			alphaNumericSpaces: "\u00deo\u0301rdi\u0301s Gu\u00f0mundsdo\u0301ttir",
			formalName:         "\u00deo\u0301rdi\u0301s Gu\u00f0mundsdo\u0301ttir",
		},
		{
			name:               "turkish name, precomposed", // Sukru Inonu
			input:              "\u015e\u00fckr\u00fc \u0130n\u00f6n\u00fc",
			alpha:              "\u015e\u00fckr\u00fc\u0130n\u00f6n\u00fc",
			alphaSpaces:        "\u015e\u00fckr\u00fc \u0130n\u00f6n\u00fc",
			alphaNumeric:       "\u015e\u00fckr\u00fc\u0130n\u00f6n\u00fc",
			alphaNumericSpaces: "\u015e\u00fckr\u00fc \u0130n\u00f6n\u00fc",
			formalName:         "\u015e\u00fckr\u00fc \u0130n\u00f6n\u00fc",
		},
		{
			name:               "turkish name, decomposed", // Sukru Inonu
			input:              "S\u0327u\u0308kru\u0308 I\u0307no\u0308nu\u0308",
			alpha:              "S\u0327u\u0308kru\u0308I\u0307no\u0308nu\u0308",
			alphaSpaces:        "S\u0327u\u0308kru\u0308 I\u0307no\u0308nu\u0308",
			alphaNumeric:       "S\u0327u\u0308kru\u0308I\u0307no\u0308nu\u0308",
			alphaNumericSpaces: "S\u0327u\u0308kru\u0308 I\u0307no\u0308nu\u0308",
			formalName:         "S\u0327u\u0308kru\u0308 I\u0307no\u0308nu\u0308",
		},
		{
			name:               "romanian name, precomposed", // Stefan Tepes
			input:              "\u0218tefan \u021aepe\u0219",
			alpha:              "\u0218tefan\u021aepe\u0219",
			alphaSpaces:        "\u0218tefan \u021aepe\u0219",
			alphaNumeric:       "\u0218tefan\u021aepe\u0219",
			alphaNumericSpaces: "\u0218tefan \u021aepe\u0219",
			formalName:         "\u0218tefan \u021aepe\u0219",
		},
		{
			name:               "romanian name, decomposed", // Stefan Tepes
			input:              "S\u0326tefan T\u0326epes\u0326",
			alpha:              "S\u0326tefanT\u0326epes\u0326",
			alphaSpaces:        "S\u0326tefan T\u0326epes\u0326",
			alphaNumeric:       "S\u0326tefanT\u0326epes\u0326",
			alphaNumericSpaces: "S\u0326tefan T\u0326epes\u0326",
			formalName:         "S\u0326tefan T\u0326epes\u0326",
		},
		{
			name:               "vietnamese name, precomposed", // Nguyen Van An
			input:              "Nguy\u1ec5n V\u0103n An",
			alpha:              "Nguy\u1ec5nV\u0103nAn",
			alphaSpaces:        "Nguy\u1ec5n V\u0103n An",
			alphaNumeric:       "Nguy\u1ec5nV\u0103nAn",
			alphaNumericSpaces: "Nguy\u1ec5n V\u0103n An",
			formalName:         "Nguy\u1ec5n V\u0103n An",
		},
		{
			name:               "vietnamese name, decomposed", // Nguyen Van An
			input:              "Nguye\u0302\u0303n Va\u0306n An",
			alpha:              "Nguye\u0302\u0303nVa\u0306nAn",
			alphaSpaces:        "Nguye\u0302\u0303n Va\u0306n An",
			alphaNumeric:       "Nguye\u0302\u0303nVa\u0306nAn",
			alphaNumericSpaces: "Nguye\u0302\u0303n Va\u0306n An",
			formalName:         "Nguye\u0302\u0303n Va\u0306n An",
		},
		{
			name:               "vietnamese name with dots below, precomposed", // Dang Thi Huong
			input:              "\u0110\u1eb7ng Th\u1ecb H\u01b0\u01a1ng",
			alpha:              "\u0110\u1eb7ngTh\u1ecbH\u01b0\u01a1ng",
			alphaSpaces:        "\u0110\u1eb7ng Th\u1ecb H\u01b0\u01a1ng",
			alphaNumeric:       "\u0110\u1eb7ngTh\u1ecbH\u01b0\u01a1ng",
			alphaNumericSpaces: "\u0110\u1eb7ng Th\u1ecb H\u01b0\u01a1ng",
			formalName:         "\u0110\u1eb7ng Th\u1ecb H\u01b0\u01a1ng",
		},
		{
			name:               "vietnamese name with dots below, decomposed", // Dang Thi Huong
			input:              "\u0110a\u0323\u0306ng Thi\u0323 Hu\u031bo\u031bng",
			alpha:              "\u0110a\u0323\u0306ngThi\u0323Hu\u031bo\u031bng",
			alphaSpaces:        "\u0110a\u0323\u0306ng Thi\u0323 Hu\u031bo\u031bng",
			alphaNumeric:       "\u0110a\u0323\u0306ngThi\u0323Hu\u031bo\u031bng",
			alphaNumericSpaces: "\u0110a\u0323\u0306ng Thi\u0323 Hu\u031bo\u031bng",
			formalName:         "\u0110a\u0323\u0306ng Thi\u0323 Hu\u031bo\u031bng",
		},
		{
			name:               "yoruba name, precomposed", // Olorunfemi Adeyemi
			input:              "\u1eccl\u1ecd\u0301runf\u1eb9\u0301mi Ad\u00e9y\u1eb9m\u00ed",
			alpha:              "\u1eccl\u1ecd\u0301runf\u1eb9\u0301miAd\u00e9y\u1eb9m\u00ed",
			alphaSpaces:        "\u1eccl\u1ecd\u0301runf\u1eb9\u0301mi Ad\u00e9y\u1eb9m\u00ed",
			alphaNumeric:       "\u1eccl\u1ecd\u0301runf\u1eb9\u0301miAd\u00e9y\u1eb9m\u00ed",
			alphaNumericSpaces: "\u1eccl\u1ecd\u0301runf\u1eb9\u0301mi Ad\u00e9y\u1eb9m\u00ed",
			formalName:         "\u1eccl\u1ecd\u0301runf\u1eb9\u0301mi Ad\u00e9y\u1eb9m\u00ed",
		},
		{
			name:               "yoruba name, decomposed", // Olorunfemi Adeyemi
			input:              "O\u0323lo\u0323\u0301runfe\u0323\u0301mi Ade\u0301ye\u0323mi\u0301",
			alpha:              "O\u0323lo\u0323\u0301runfe\u0323\u0301miAde\u0301ye\u0323mi\u0301",
			alphaSpaces:        "O\u0323lo\u0323\u0301runfe\u0323\u0301mi Ade\u0301ye\u0323mi\u0301",
			alphaNumeric:       "O\u0323lo\u0323\u0301runfe\u0323\u0301miAde\u0301ye\u0323mi\u0301",
			alphaNumericSpaces: "O\u0323lo\u0323\u0301runfe\u0323\u0301mi Ade\u0301ye\u0323mi\u0301",
			formalName:         "O\u0323lo\u0323\u0301runfe\u0323\u0301mi Ade\u0301ye\u0323mi\u0301",
		},
		{
			name:               "lithuanian dotted i with grave", // Jinas
			input:              "Ji\u0307\u0300nas",
			alpha:              "Ji\u0307\u0300nas",
			alphaSpaces:        "Ji\u0307\u0300nas",
			alphaNumeric:       "Ji\u0307\u0300nas",
			alphaNumericSpaces: "Ji\u0307\u0300nas",
			formalName:         "Ji\u0307\u0300nas",
		},
		{
			name:               "latin ligatures", // fine flow IJssel
			input:              "\ufb01ne \ufb02ow \u0132ssel",
			alpha:              "\ufb01ne\ufb02ow\u0132ssel",
			alphaSpaces:        "\ufb01ne \ufb02ow \u0132ssel",
			alphaNumeric:       "\ufb01ne\ufb02ow\u0132ssel",
			alphaNumericSpaces: "\ufb01ne \ufb02ow \u0132ssel",
			formalName:         "\ufb01ne \ufb02ow \u0132ssel",
		},
		{
			name:               "titlecase digraph", // Duro
			input:              "\u01c5uro",
			alpha:              "\u01c5uro",
			alphaSpaces:        "\u01c5uro",
			alphaNumeric:       "\u01c5uro",
			alphaNumericSpaces: "\u01c5uro",
			formalName:         "\u01c5uro",
		},
		{
			name:               "greek name, precomposed", // Alexandros Papadopoulos
			input:              "\u0391\u03bb\u03ad\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03cc\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alpha:              "\u0391\u03bb\u03ad\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2\u03a0\u03b1\u03c0\u03b1\u03b4\u03cc\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alphaSpaces:        "\u0391\u03bb\u03ad\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03cc\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alphaNumeric:       "\u0391\u03bb\u03ad\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2\u03a0\u03b1\u03c0\u03b1\u03b4\u03cc\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alphaNumericSpaces: "\u0391\u03bb\u03ad\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03cc\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			formalName:         "\u0391\u03bb\u03ad\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03cc\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
		},
		{
			name:               "greek name, decomposed", // Alexandros Papadopoulos
			input:              "\u0391\u03bb\u03b5\u0301\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03bf\u0301\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alpha:              "\u0391\u03bb\u03b5\u0301\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2\u03a0\u03b1\u03c0\u03b1\u03b4\u03bf\u0301\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alphaSpaces:        "\u0391\u03bb\u03b5\u0301\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03bf\u0301\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alphaNumeric:       "\u0391\u03bb\u03b5\u0301\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2\u03a0\u03b1\u03c0\u03b1\u03b4\u03bf\u0301\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			alphaNumericSpaces: "\u0391\u03bb\u03b5\u0301\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03bf\u0301\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
			formalName:         "\u0391\u03bb\u03b5\u0301\u03be\u03b1\u03bd\u03b4\u03c1\u03bf\u03c2 \u03a0\u03b1\u03c0\u03b1\u03b4\u03bf\u0301\u03c0\u03bf\u03c5\u03bb\u03bf\u03c2",
		},
		{
			name:               "polytonic greek, precomposed", // Athena Hades
			input:              "\u1f08\u03b8\u03b7\u03bd\u1fb6 \u1f85\u03b4\u03b7\u03c2",
			alpha:              "\u1f08\u03b8\u03b7\u03bd\u1fb6\u1f85\u03b4\u03b7\u03c2",
			alphaSpaces:        "\u1f08\u03b8\u03b7\u03bd\u1fb6 \u1f85\u03b4\u03b7\u03c2",
			alphaNumeric:       "\u1f08\u03b8\u03b7\u03bd\u1fb6\u1f85\u03b4\u03b7\u03c2",
			alphaNumericSpaces: "\u1f08\u03b8\u03b7\u03bd\u1fb6 \u1f85\u03b4\u03b7\u03c2",
			formalName:         "\u1f08\u03b8\u03b7\u03bd\u1fb6 \u1f85\u03b4\u03b7\u03c2",
		},
		{
			name:               "polytonic greek, decomposed", // Athena Hades
			input:              "\u0391\u0313\u03b8\u03b7\u03bd\u03b1\u0342 \u03b1\u0314\u0301\u0345\u03b4\u03b7\u03c2",
			alpha:              "\u0391\u0313\u03b8\u03b7\u03bd\u03b1\u0342\u03b1\u0314\u0301\u0345\u03b4\u03b7\u03c2",
			alphaSpaces:        "\u0391\u0313\u03b8\u03b7\u03bd\u03b1\u0342 \u03b1\u0314\u0301\u0345\u03b4\u03b7\u03c2",
			alphaNumeric:       "\u0391\u0313\u03b8\u03b7\u03bd\u03b1\u0342\u03b1\u0314\u0301\u0345\u03b4\u03b7\u03c2",
			alphaNumericSpaces: "\u0391\u0313\u03b8\u03b7\u03bd\u03b1\u0342 \u03b1\u0314\u0301\u0345\u03b4\u03b7\u03c2",
			formalName:         "\u0391\u0313\u03b8\u03b7\u03bd\u03b1\u0342 \u03b1\u0314\u0301\u0345\u03b4\u03b7\u03c2",
		},
		{
			name:               "cyrillic name, precomposed", // Vladimir Yosipovich Yolkin
			input:              "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0419\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0401\u043b\u043a\u0438\u043d",
			alpha:              "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440\u0419\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447\u0401\u043b\u043a\u0438\u043d",
			alphaSpaces:        "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0419\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0401\u043b\u043a\u0438\u043d",
			alphaNumeric:       "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440\u0419\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447\u0401\u043b\u043a\u0438\u043d",
			alphaNumericSpaces: "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0419\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0401\u043b\u043a\u0438\u043d",
			formalName:         "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0419\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0401\u043b\u043a\u0438\u043d",
		},
		{
			name:               "cyrillic name, decomposed", // Vladimir Yosipovich Yolkin
			input:              "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0418\u0306\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0415\u0308\u043b\u043a\u0438\u043d",
			alpha:              "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440\u0418\u0306\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447\u0415\u0308\u043b\u043a\u0438\u043d",
			alphaSpaces:        "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0418\u0306\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0415\u0308\u043b\u043a\u0438\u043d",
			alphaNumeric:       "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440\u0418\u0306\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447\u0415\u0308\u043b\u043a\u0438\u043d",
			alphaNumericSpaces: "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0418\u0306\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0415\u0308\u043b\u043a\u0438\u043d",
			formalName:         "\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u0418\u0306\u043e\u0441\u0438\u043f\u043e\u0432\u0438\u0447 \u0415\u0308\u043b\u043a\u0438\u043d",
		},
		{
			name:               "armenian name", // Aram Khachaturyan
			input:              "\u0531\u0580\u0561\u0574 \u053d\u0561\u0579\u0561\u057f\u0580\u0575\u0561\u0576",
			alpha:              "\u0531\u0580\u0561\u0574\u053d\u0561\u0579\u0561\u057f\u0580\u0575\u0561\u0576",
			alphaSpaces:        "\u0531\u0580\u0561\u0574 \u053d\u0561\u0579\u0561\u057f\u0580\u0575\u0561\u0576",
			alphaNumeric:       "\u0531\u0580\u0561\u0574\u053d\u0561\u0579\u0561\u057f\u0580\u0575\u0561\u0576",
			alphaNumericSpaces: "\u0531\u0580\u0561\u0574 \u053d\u0561\u0579\u0561\u057f\u0580\u0575\u0561\u0576",
			formalName:         "\u0531\u0580\u0561\u0574 \u053d\u0561\u0579\u0561\u057f\u0580\u0575\u0561\u0576",
		},
		{
			name:               "armenian apostrophe", // Aram
			input:              "\u0531\u055a\u0580\u0561\u0574",
			alpha:              "\u0531\u0580\u0561\u0574",
			alphaSpaces:        "\u0531\u0580\u0561\u0574",
			alphaNumeric:       "\u0531\u0580\u0561\u0574",
			alphaNumericSpaces: "\u0531\u0580\u0561\u0574",
			formalName:         "\u0531\u0580\u0561\u0574",
		},
		{
			name:               "georgian name", // Giorgi
			input:              "\u10d2\u10d8\u10dd\u10e0\u10d2\u10d8",
			alpha:              "\u10d2\u10d8\u10dd\u10e0\u10d2\u10d8",
			alphaSpaces:        "\u10d2\u10d8\u10dd\u10e0\u10d2\u10d8",
			alphaNumeric:       "\u10d2\u10d8\u10dd\u10e0\u10d2\u10d8",
			alphaNumericSpaces: "\u10d2\u10d8\u10dd\u10e0\u10d2\u10d8",
			formalName:         "\u10d2\u10d8\u10dd\u10e0\u10d2\u10d8",
		},
		{
			name:               "arabic name", // Muhammad Abdullah
			input:              "\u0645\u062d\u0645\u062f \u0639\u0628\u062f\u0627\u0644\u0644\u0647",
			alpha:              "\u0645\u062d\u0645\u062f\u0639\u0628\u062f\u0627\u0644\u0644\u0647",
			alphaSpaces:        "\u0645\u062d\u0645\u062f \u0639\u0628\u062f\u0627\u0644\u0644\u0647",
			alphaNumeric:       "\u0645\u062d\u0645\u062f\u0639\u0628\u062f\u0627\u0644\u0644\u0647",
			alphaNumericSpaces: "\u0645\u062d\u0645\u062f \u0639\u0628\u062f\u0627\u0644\u0644\u0647",
			formalName:         "\u0645\u062d\u0645\u062f \u0639\u0628\u062f\u0627\u0644\u0644\u0647",
		},
		{
			name:               "arabic name with vowel marks", // Muhammad
			input:              "\u0645\u064f\u062d\u064e\u0645\u064e\u0651\u062f",
			alpha:              "\u0645\u064f\u062d\u064e\u0645\u064e\u0651\u062f",
			alphaSpaces:        "\u0645\u064f\u062d\u064e\u0645\u064e\u0651\u062f",
			alphaNumeric:       "\u0645\u064f\u062d\u064e\u0645\u064e\u0651\u062f",
			alphaNumericSpaces: "\u0645\u064f\u062d\u064e\u0645\u064e\u0651\u062f",
			formalName:         "\u0645\u064f\u062d\u064e\u0645\u064e\u0651\u062f",
		},
		{
			name:               "arabic tatweel", // Muhammad
			input:              "\u0645\u062d\u0640\u0640\u0640\u0645\u062f",
			alpha:              "\u0645\u062d\u0640\u0640\u0640\u0645\u062f",
			alphaSpaces:        "\u0645\u062d\u0640\u0640\u0640\u0645\u062f",
			alphaNumeric:       "\u0645\u062d\u0640\u0640\u0640\u0645\u062f",
			alphaNumericSpaces: "\u0645\u062d\u0640\u0640\u0640\u0645\u062f",
			formalName:         "\u0645\u062d\u0640\u0640\u0640\u0645\u062f",
		},
		{
			name:               "arabic-indic digits", // raqm 0123456789
			input:              "\u0631\u0642\u0645 \u0660\u0661\u0662\u0663\u0664\u0665\u0666\u0667\u0668\u0669",
			alpha:              "\u0631\u0642\u0645",
			alphaSpaces:        "\u0631\u0642\u0645 ",
			alphaNumeric:       "\u0631\u0642\u0645\u0660\u0661\u0662\u0663\u0664\u0665\u0666\u0667\u0668\u0669",
			alphaNumericSpaces: "\u0631\u0642\u0645 \u0660\u0661\u0662\u0663\u0664\u0665\u0666\u0667\u0668\u0669",
			formalName:         "\u0631\u0642\u0645 \u0660\u0661\u0662\u0663\u0664\u0665\u0666\u0667\u0668\u0669",
		},
		{
			name:               "persian word with zero-width non-joiner", // mikhaham
			input:              "\u0645\u06cc\u200c\u062e\u0648\u0627\u0647\u0645",
			alpha:              "\u0645\u06cc\u062e\u0648\u0627\u0647\u0645",
			alphaSpaces:        "\u0645\u06cc\u062e\u0648\u0627\u0647\u0645",
			alphaNumeric:       "\u0645\u06cc\u062e\u0648\u0627\u0647\u0645",
			alphaNumericSpaces: "\u0645\u06cc\u062e\u0648\u0627\u0647\u0645",
			formalName:         "\u0645\u06cc\u062e\u0648\u0627\u0647\u0645",
		},
		{
			name:               "hebrew name", // David Cohen
			input:              "\u05d3\u05d5\u05d3 \u05db\u05d4\u05df",
			alpha:              "\u05d3\u05d5\u05d3\u05db\u05d4\u05df",
			alphaSpaces:        "\u05d3\u05d5\u05d3 \u05db\u05d4\u05df",
			alphaNumeric:       "\u05d3\u05d5\u05d3\u05db\u05d4\u05df",
			alphaNumericSpaces: "\u05d3\u05d5\u05d3 \u05db\u05d4\u05df",
			formalName:         "\u05d3\u05d5\u05d3 \u05db\u05d4\u05df",
		},
		{
			name:               "hebrew name with niqqud", // David
			input:              "\u05d3\u05b8\u05bc\u05d5\u05b4\u05d3",
			alpha:              "\u05d3\u05b8\u05bc\u05d5\u05b4\u05d3",
			alphaSpaces:        "\u05d3\u05b8\u05bc\u05d5\u05b4\u05d3",
			alphaNumeric:       "\u05d3\u05b8\u05bc\u05d5\u05b4\u05d3",
			alphaNumericSpaces: "\u05d3\u05b8\u05bc\u05d5\u05b4\u05d3",
			formalName:         "\u05d3\u05b8\u05bc\u05d5\u05b4\u05d3",
		},
		{
			name:               "hebrew word with cantillation", // bereshit
			input:              "\u05d1\u05b0\u05bc\u05e8\u05b5\u05d0\u05e9\u05b4\u05c1\u0596\u05d9\u05ea",
			alpha:              "\u05d1\u05b0\u05bc\u05e8\u05b5\u05d0\u05e9\u05b4\u05c1\u0596\u05d9\u05ea",
			alphaSpaces:        "\u05d1\u05b0\u05bc\u05e8\u05b5\u05d0\u05e9\u05b4\u05c1\u0596\u05d9\u05ea",
			alphaNumeric:       "\u05d1\u05b0\u05bc\u05e8\u05b5\u05d0\u05e9\u05b4\u05c1\u0596\u05d9\u05ea",
			alphaNumericSpaces: "\u05d1\u05b0\u05bc\u05e8\u05b5\u05d0\u05e9\u05b4\u05c1\u0596\u05d9\u05ea",
			formalName:         "\u05d1\u05b0\u05bc\u05e8\u05b5\u05d0\u05e9\u05b4\u05c1\u0596\u05d9\u05ea",
		},
		{
			name:               "devanagari name", // Rohit Varma
			input:              "\u0930\u094b\u0939\u093f\u0924 \u0935\u0930\u094d\u092e\u093e",
			alpha:              "\u0930\u094b\u0939\u093f\u0924\u0935\u0930\u094d\u092e\u093e",
			alphaSpaces:        "\u0930\u094b\u0939\u093f\u0924 \u0935\u0930\u094d\u092e\u093e",
			alphaNumeric:       "\u0930\u094b\u0939\u093f\u0924\u0935\u0930\u094d\u092e\u093e",
			alphaNumericSpaces: "\u0930\u094b\u0939\u093f\u0924 \u0935\u0930\u094d\u092e\u093e",
			formalName:         "\u0930\u094b\u0939\u093f\u0924 \u0935\u0930\u094d\u092e\u093e",
		},
		{
			name:               "devanagari name with nukta and anusvara", // Priyanka Chopra
			input:              "\u092a\u094d\u0930\u093f\u092f\u0902\u0915\u093e \u091a\u094b\u092a\u0921\u093c\u093e",
			alpha:              "\u092a\u094d\u0930\u093f\u092f\u0902\u0915\u093e\u091a\u094b\u092a\u0921\u093c\u093e",
			alphaSpaces:        "\u092a\u094d\u0930\u093f\u092f\u0902\u0915\u093e \u091a\u094b\u092a\u0921\u093c\u093e",
			alphaNumeric:       "\u092a\u094d\u0930\u093f\u092f\u0902\u0915\u093e\u091a\u094b\u092a\u0921\u093c\u093e",
			alphaNumericSpaces: "\u092a\u094d\u0930\u093f\u092f\u0902\u0915\u093e \u091a\u094b\u092a\u0921\u093c\u093e",
			formalName:         "\u092a\u094d\u0930\u093f\u092f\u0902\u0915\u093e \u091a\u094b\u092a\u0921\u093c\u093e",
		},
		{
			name:               "devanagari precomposed nukta letter", // Chopra
			input:              "\u091a\u094b\u092a\u095c\u093e",
			alpha:              "\u091a\u094b\u092a\u095c\u093e",
			alphaSpaces:        "\u091a\u094b\u092a\u095c\u093e",
			alphaNumeric:       "\u091a\u094b\u092a\u095c\u093e",
			alphaNumericSpaces: "\u091a\u094b\u092a\u095c\u093e",
			formalName:         "\u091a\u094b\u092a\u095c\u093e",
		},
		{
			name:               "devanagari visarga and chandrabindu", // dukh hans
			input:              "\u0926\u0941\u0903\u0916 \u0939\u0901\u0938",
			alpha:              "\u0926\u0941\u0903\u0916\u0939\u0901\u0938",
			alphaSpaces:        "\u0926\u0941\u0903\u0916 \u0939\u0901\u0938",
			alphaNumeric:       "\u0926\u0941\u0903\u0916\u0939\u0901\u0938",
			alphaNumericSpaces: "\u0926\u0941\u0903\u0916 \u0939\u0901\u0938",
			formalName:         "\u0926\u0941\u0903\u0916 \u0939\u0901\u0938",
		},
		{
			name:               "devanagari digits", // makan 123
			input:              "\u092e\u0915\u093e\u0928 \u0967\u0968\u0969",
			alpha:              "\u092e\u0915\u093e\u0928",
			alphaSpaces:        "\u092e\u0915\u093e\u0928 ",
			alphaNumeric:       "\u092e\u0915\u093e\u0928\u0967\u0968\u0969",
			alphaNumericSpaces: "\u092e\u0915\u093e\u0928 \u0967\u0968\u0969",
			formalName:         "\u092e\u0915\u093e\u0928 \u0967\u0968\u0969",
		},
		{
			name:               "bengali name, precomposed", // Sourav Ganguly
			input:              "\u09b8\u09cc\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alpha:              "\u09b8\u09cc\u09b0\u09ad\u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alphaSpaces:        "\u09b8\u09cc\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alphaNumeric:       "\u09b8\u09cc\u09b0\u09ad\u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alphaNumericSpaces: "\u09b8\u09cc\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			formalName:         "\u09b8\u09cc\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
		},
		{
			name:               "bengali name, decomposed", // Sourav Ganguly
			input:              "\u09b8\u09c7\u09d7\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alpha:              "\u09b8\u09c7\u09d7\u09b0\u09ad\u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alphaSpaces:        "\u09b8\u09c7\u09d7\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alphaNumeric:       "\u09b8\u09c7\u09d7\u09b0\u09ad\u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			alphaNumericSpaces: "\u09b8\u09c7\u09d7\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
			formalName:         "\u09b8\u09c7\u09d7\u09b0\u09ad \u0997\u09be\u0999\u09cd\u0997\u09c1\u09b2\u09c0",
		},
		{
			name:               "gurmukhi words", // Singh Panjab
			input:              "\u0a38\u0a3f\u0a70\u0a18 \u0a2a\u0a70\u0a1c\u0a3e\u0a2c",
			alpha:              "\u0a38\u0a3f\u0a70\u0a18\u0a2a\u0a70\u0a1c\u0a3e\u0a2c",
			alphaSpaces:        "\u0a38\u0a3f\u0a70\u0a18 \u0a2a\u0a70\u0a1c\u0a3e\u0a2c",
			alphaNumeric:       "\u0a38\u0a3f\u0a70\u0a18\u0a2a\u0a70\u0a1c\u0a3e\u0a2c",
			alphaNumericSpaces: "\u0a38\u0a3f\u0a70\u0a18 \u0a2a\u0a70\u0a1c\u0a3e\u0a2c",
			formalName:         "\u0a38\u0a3f\u0a70\u0a18 \u0a2a\u0a70\u0a1c\u0a3e\u0a2c",
		},
		{
			name:               "gujarati name", // Narendra
			input:              "\u0aa8\u0ab0\u0ac7\u0aa8\u0acd\u0aa6\u0acd\u0ab0",
			alpha:              "\u0aa8\u0ab0\u0ac7\u0aa8\u0acd\u0aa6\u0acd\u0ab0",
			alphaSpaces:        "\u0aa8\u0ab0\u0ac7\u0aa8\u0acd\u0aa6\u0acd\u0ab0",
			alphaNumeric:       "\u0aa8\u0ab0\u0ac7\u0aa8\u0acd\u0aa6\u0acd\u0ab0",
			alphaNumericSpaces: "\u0aa8\u0ab0\u0ac7\u0aa8\u0acd\u0aa6\u0acd\u0ab0",
			formalName:         "\u0aa8\u0ab0\u0ac7\u0aa8\u0acd\u0aa6\u0acd\u0ab0",
		},
		{
			name:               "oriya word", // Odia
			input:              "\u0b13\u0b21\u0b3c\u0b3f\u0b06",
			alpha:              "\u0b13\u0b21\u0b3c\u0b3f\u0b06",
			alphaSpaces:        "\u0b13\u0b21\u0b3c\u0b3f\u0b06",
			alphaNumeric:       "\u0b13\u0b21\u0b3c\u0b3f\u0b06",
			alphaNumericSpaces: "\u0b13\u0b21\u0b3c\u0b3f\u0b06",
			formalName:         "\u0b13\u0b21\u0b3c\u0b3f\u0b06",
		},
		{
			name:               "tamil name", // Muttiah Muralitharan
			input:              "\u0bae\u0bc1\u0ba4\u0bcd\u0ba4\u0bc8\u0baf\u0bbe \u0bae\u0bc1\u0bb0\u0bb3\u0bbf\u0ba4\u0bb0\u0ba9\u0bcd",
			alpha:              "\u0bae\u0bc1\u0ba4\u0bcd\u0ba4\u0bc8\u0baf\u0bbe\u0bae\u0bc1\u0bb0\u0bb3\u0bbf\u0ba4\u0bb0\u0ba9\u0bcd",
			alphaSpaces:        "\u0bae\u0bc1\u0ba4\u0bcd\u0ba4\u0bc8\u0baf\u0bbe \u0bae\u0bc1\u0bb0\u0bb3\u0bbf\u0ba4\u0bb0\u0ba9\u0bcd",
			alphaNumeric:       "\u0bae\u0bc1\u0ba4\u0bcd\u0ba4\u0bc8\u0baf\u0bbe\u0bae\u0bc1\u0bb0\u0bb3\u0bbf\u0ba4\u0bb0\u0ba9\u0bcd",
			alphaNumericSpaces: "\u0bae\u0bc1\u0ba4\u0bcd\u0ba4\u0bc8\u0baf\u0bbe \u0bae\u0bc1\u0bb0\u0bb3\u0bbf\u0ba4\u0bb0\u0ba9\u0bcd",
			formalName:         "\u0bae\u0bc1\u0ba4\u0bcd\u0ba4\u0bc8\u0baf\u0bbe \u0bae\u0bc1\u0bb0\u0bb3\u0bbf\u0ba4\u0bb0\u0ba9\u0bcd",
		},
		{
			name:               "tamil word with a two-part vowel, precomposed", // kodi
			input:              "\u0b95\u0bca\u0b9f\u0bbf",
			alpha:              "\u0b95\u0bca\u0b9f\u0bbf",
			alphaSpaces:        "\u0b95\u0bca\u0b9f\u0bbf",
			alphaNumeric:       "\u0b95\u0bca\u0b9f\u0bbf",
			alphaNumericSpaces: "\u0b95\u0bca\u0b9f\u0bbf",
			formalName:         "\u0b95\u0bca\u0b9f\u0bbf",
		},
		{
			name:               "tamil word with a two-part vowel, decomposed", // kodi
			input:              "\u0b95\u0bc6\u0bbe\u0b9f\u0bbf",
			alpha:              "\u0b95\u0bc6\u0bbe\u0b9f\u0bbf",
			alphaSpaces:        "\u0b95\u0bc6\u0bbe\u0b9f\u0bbf",
			alphaNumeric:       "\u0b95\u0bc6\u0bbe\u0b9f\u0bbf",
			alphaNumericSpaces: "\u0b95\u0bc6\u0bbe\u0b9f\u0bbf",
			formalName:         "\u0b95\u0bc6\u0bbe\u0b9f\u0bbf",
		},
		{
			name:               "telugu name", // Venkatesh
			input:              "\u0c35\u0c46\u0c02\u0c15\u0c1f\u0c47\u0c36\u0c4d",
			alpha:              "\u0c35\u0c46\u0c02\u0c15\u0c1f\u0c47\u0c36\u0c4d",
			alphaSpaces:        "\u0c35\u0c46\u0c02\u0c15\u0c1f\u0c47\u0c36\u0c4d",
			alphaNumeric:       "\u0c35\u0c46\u0c02\u0c15\u0c1f\u0c47\u0c36\u0c4d",
			alphaNumericSpaces: "\u0c35\u0c46\u0c02\u0c15\u0c1f\u0c47\u0c36\u0c4d",
			formalName:         "\u0c35\u0c46\u0c02\u0c15\u0c1f\u0c47\u0c36\u0c4d",
		},
		{
			name:               "kannada words", // Kannada Raj
			input:              "\u0c95\u0ca8\u0ccd\u0ca8\u0ca1 \u0cb0\u0cbe\u0c9c\u0ccd",
			alpha:              "\u0c95\u0ca8\u0ccd\u0ca8\u0ca1\u0cb0\u0cbe\u0c9c\u0ccd",
			alphaSpaces:        "\u0c95\u0ca8\u0ccd\u0ca8\u0ca1 \u0cb0\u0cbe\u0c9c\u0ccd",
			alphaNumeric:       "\u0c95\u0ca8\u0ccd\u0ca8\u0ca1\u0cb0\u0cbe\u0c9c\u0ccd",
			alphaNumericSpaces: "\u0c95\u0ca8\u0ccd\u0ca8\u0ca1 \u0cb0\u0cbe\u0c9c\u0ccd",
			formalName:         "\u0c95\u0ca8\u0ccd\u0ca8\u0ca1 \u0cb0\u0cbe\u0c9c\u0ccd",
		},
		{
			name:               "malayalam word", // Malayalam
			input:              "\u0d2e\u0d32\u0d2f\u0d3e\u0d33\u0d02",
			alpha:              "\u0d2e\u0d32\u0d2f\u0d3e\u0d33\u0d02",
			alphaSpaces:        "\u0d2e\u0d32\u0d2f\u0d3e\u0d33\u0d02",
			alphaNumeric:       "\u0d2e\u0d32\u0d2f\u0d3e\u0d33\u0d02",
			alphaNumericSpaces: "\u0d2e\u0d32\u0d2f\u0d3e\u0d33\u0d02",
			formalName:         "\u0d2e\u0d32\u0d2f\u0d3e\u0d33\u0d02",
		},
		{
			name:               "sinhala name with zero-width joiner", // Sri Lanka
			input:              "\u0dc1\u0dca\u200d\u0dbb\u0dd3 \u0dbd\u0d82\u0d9a\u0dcf",
			alpha:              "\u0dc1\u0dca\u0dbb\u0dd3\u0dbd\u0d82\u0d9a\u0dcf",
			alphaSpaces:        "\u0dc1\u0dca\u0dbb\u0dd3 \u0dbd\u0d82\u0d9a\u0dcf",
			alphaNumeric:       "\u0dc1\u0dca\u0dbb\u0dd3\u0dbd\u0d82\u0d9a\u0dcf",
			alphaNumericSpaces: "\u0dc1\u0dca\u0dbb\u0dd3 \u0dbd\u0d82\u0d9a\u0dcf",
			formalName:         "\u0dc1\u0dca\u0dbb\u0dd3 \u0dbd\u0d82\u0d9a\u0dcf",
		},
		{
			name:               "thai name", // Somsak Jaidee
			input:              "\u0e2a\u0e21\u0e28\u0e31\u0e01\u0e14\u0e34\u0e4c \u0e43\u0e08\u0e14\u0e35",
			alpha:              "\u0e2a\u0e21\u0e28\u0e31\u0e01\u0e14\u0e34\u0e4c\u0e43\u0e08\u0e14\u0e35",
			alphaSpaces:        "\u0e2a\u0e21\u0e28\u0e31\u0e01\u0e14\u0e34\u0e4c \u0e43\u0e08\u0e14\u0e35",
			alphaNumeric:       "\u0e2a\u0e21\u0e28\u0e31\u0e01\u0e14\u0e34\u0e4c\u0e43\u0e08\u0e14\u0e35",
			alphaNumericSpaces: "\u0e2a\u0e21\u0e28\u0e31\u0e01\u0e14\u0e34\u0e4c \u0e43\u0e08\u0e14\u0e35",
			formalName:         "\u0e2a\u0e21\u0e28\u0e31\u0e01\u0e14\u0e34\u0e4c \u0e43\u0e08\u0e14\u0e35",
		},
		{
			name:               "thai digits", // baan 123
			input:              "\u0e1a\u0e49\u0e32\u0e19 \u0e51\u0e52\u0e53",
			alpha:              "\u0e1a\u0e49\u0e32\u0e19",
			alphaSpaces:        "\u0e1a\u0e49\u0e32\u0e19 ",
			alphaNumeric:       "\u0e1a\u0e49\u0e32\u0e19\u0e51\u0e52\u0e53",
			alphaNumericSpaces: "\u0e1a\u0e49\u0e32\u0e19 \u0e51\u0e52\u0e53",
			formalName:         "\u0e1a\u0e49\u0e32\u0e19 \u0e51\u0e52\u0e53",
		},
		{
			name:               "lao name", // Somsak
			input:              "\u0eaa\u0ebb\u0ea1\u0eaa\u0eb1\u0e81",
			alpha:              "\u0eaa\u0ebb\u0ea1\u0eaa\u0eb1\u0e81",
			alphaSpaces:        "\u0eaa\u0ebb\u0ea1\u0eaa\u0eb1\u0e81",
			alphaNumeric:       "\u0eaa\u0ebb\u0ea1\u0eaa\u0eb1\u0e81",
			alphaNumericSpaces: "\u0eaa\u0ebb\u0ea1\u0eaa\u0eb1\u0e81",
			formalName:         "\u0eaa\u0ebb\u0ea1\u0eaa\u0eb1\u0e81",
		},
		{
			name:               "khmer name", // Sok Sina
			input:              "\u179f\u17bb\u1781 \u179f\u17ca\u17b8\u178e\u17b6",
			alpha:              "\u179f\u17bb\u1781\u179f\u17ca\u17b8\u178e\u17b6",
			alphaSpaces:        "\u179f\u17bb\u1781 \u179f\u17ca\u17b8\u178e\u17b6",
			alphaNumeric:       "\u179f\u17bb\u1781\u179f\u17ca\u17b8\u178e\u17b6",
			alphaNumericSpaces: "\u179f\u17bb\u1781 \u179f\u17ca\u17b8\u178e\u17b6",
			formalName:         "\u179f\u17bb\u1781 \u179f\u17ca\u17b8\u178e\u17b6",
		},
		{
			name:               "myanmar name", // Aung San Suu Kyi
			input:              "\u1021\u1031\u102c\u1004\u103a\u1006\u1014\u103a\u1038\u1005\u102f\u1000\u103c\u100a\u103a",
			alpha:              "\u1021\u1031\u102c\u1004\u103a\u1006\u1014\u103a\u1038\u1005\u102f\u1000\u103c\u100a\u103a",
			alphaSpaces:        "\u1021\u1031\u102c\u1004\u103a\u1006\u1014\u103a\u1038\u1005\u102f\u1000\u103c\u100a\u103a",
			alphaNumeric:       "\u1021\u1031\u102c\u1004\u103a\u1006\u1014\u103a\u1038\u1005\u102f\u1000\u103c\u100a\u103a",
			alphaNumericSpaces: "\u1021\u1031\u102c\u1004\u103a\u1006\u1014\u103a\u1038\u1005\u102f\u1000\u103c\u100a\u103a",
			formalName:         "\u1021\u1031\u102c\u1004\u103a\u1006\u1014\u103a\u1038\u1005\u102f\u1000\u103c\u100a\u103a",
		},
		{
			name:               "tibetan name", // Tenzin
			input:              "\u0f56\u0f66\u0f9f\u0f53\u0f0b\u0f60\u0f5b\u0f72\u0f53",
			alpha:              "\u0f56\u0f66\u0f9f\u0f53\u0f60\u0f5b\u0f72\u0f53",
			alphaSpaces:        "\u0f56\u0f66\u0f9f\u0f53\u0f60\u0f5b\u0f72\u0f53",
			alphaNumeric:       "\u0f56\u0f66\u0f9f\u0f53\u0f60\u0f5b\u0f72\u0f53",
			alphaNumericSpaces: "\u0f56\u0f66\u0f9f\u0f53\u0f60\u0f5b\u0f72\u0f53",
			formalName:         "\u0f56\u0f66\u0f9f\u0f53\u0f60\u0f5b\u0f72\u0f53",
		},
		{
			name:               "chinese name", // Wang Xiaoming
			input:              "\u738b\u5c0f\u660e",
			alpha:              "\u738b\u5c0f\u660e",
			alphaSpaces:        "\u738b\u5c0f\u660e",
			alphaNumeric:       "\u738b\u5c0f\u660e",
			alphaNumericSpaces: "\u738b\u5c0f\u660e",
			formalName:         "\u738b\u5c0f\u660e",
		},
		{
			name:               "chinese name with an ideographic variation selector", // Ge
			input:              "\u845b\U000e0100",
			alpha:              "\u845b",
			alphaSpaces:        "\u845b",
			alphaNumeric:       "\u845b",
			alphaNumericSpaces: "\u845b",
			formalName:         "\u845b",
		},
		{
			name:               "japanese name in kanji", // Yamada Taro
			input:              "\u5c71\u7530\u592a\u90ce",
			alpha:              "\u5c71\u7530\u592a\u90ce",
			alphaSpaces:        "\u5c71\u7530\u592a\u90ce",
			alphaNumeric:       "\u5c71\u7530\u592a\u90ce",
			alphaNumericSpaces: "\u5c71\u7530\u592a\u90ce",
			formalName:         "\u5c71\u7530\u592a\u90ce",
		},
		{
			name:               "japanese hiragana, precomposed", // gakkou sakura
			input:              "\u304c\u3063\u3053\u3046 \u3055\u304f\u3089",
			alpha:              "\u304c\u3063\u3053\u3046\u3055\u304f\u3089",
			alphaSpaces:        "\u304c\u3063\u3053\u3046 \u3055\u304f\u3089",
			alphaNumeric:       "\u304c\u3063\u3053\u3046\u3055\u304f\u3089",
			alphaNumericSpaces: "\u304c\u3063\u3053\u3046 \u3055\u304f\u3089",
			formalName:         "\u304c\u3063\u3053\u3046 \u3055\u304f\u3089",
		},
		{
			name:               "japanese hiragana, decomposed", // gakkou sakura
			input:              "\u304b\u3099\u3063\u3053\u3046 \u3055\u304f\u3089",
			alpha:              "\u304b\u3099\u3063\u3053\u3046\u3055\u304f\u3089",
			alphaSpaces:        "\u304b\u3099\u3063\u3053\u3046 \u3055\u304f\u3089",
			alphaNumeric:       "\u304b\u3099\u3063\u3053\u3046\u3055\u304f\u3089",
			alphaNumericSpaces: "\u304b\u3099\u3063\u3053\u3046 \u3055\u304f\u3089",
			formalName:         "\u304b\u3099\u3063\u3053\u3046 \u3055\u304f\u3089",
		},
		{
			name:               "japanese katakana with a middle dot, precomposed", // Jon Sumisu
			input:              "\u30b8\u30e7\u30f3\u30fb\u30b9\u30df\u30b9",
			alpha:              "\u30b8\u30e7\u30f3\u30b9\u30df\u30b9",
			alphaSpaces:        "\u30b8\u30e7\u30f3\u30b9\u30df\u30b9",
			alphaNumeric:       "\u30b8\u30e7\u30f3\u30b9\u30df\u30b9",
			alphaNumericSpaces: "\u30b8\u30e7\u30f3\u30b9\u30df\u30b9",
			formalName:         "\u30b8\u30e7\u30f3\u30b9\u30df\u30b9",
		},
		{
			name:               "japanese katakana with a middle dot, decomposed", // Jon Sumisu
			input:              "\u30b7\u3099\u30e7\u30f3\u30fb\u30b9\u30df\u30b9",
			alpha:              "\u30b7\u3099\u30e7\u30f3\u30b9\u30df\u30b9",
			alphaSpaces:        "\u30b7\u3099\u30e7\u30f3\u30b9\u30df\u30b9",
			alphaNumeric:       "\u30b7\u3099\u30e7\u30f3\u30b9\u30df\u30b9",
			alphaNumericSpaces: "\u30b7\u3099\u30e7\u30f3\u30b9\u30df\u30b9",
			formalName:         "\u30b7\u3099\u30e7\u30f3\u30b9\u30df\u30b9",
		},
		{
			name:               "halfwidth katakana with a voiced mark", // gakkou
			input:              "\uff76\uff9e\uff6f\uff7a\uff73",
			alpha:              "\uff76\uff9e\uff6f\uff7a\uff73",
			alphaSpaces:        "\uff76\uff9e\uff6f\uff7a\uff73",
			alphaNumeric:       "\uff76\uff9e\uff6f\uff7a\uff73",
			alphaNumericSpaces: "\uff76\uff9e\uff6f\uff7a\uff73",
			formalName:         "\uff76\uff9e\uff6f\uff7a\uff73",
		},
		{
			name:               "korean name, precomposed", // Hong Gildong
			input:              "\ud64d\uae38\ub3d9",
			alpha:              "\ud64d\uae38\ub3d9",
			alphaSpaces:        "\ud64d\uae38\ub3d9",
			alphaNumeric:       "\ud64d\uae38\ub3d9",
			alphaNumericSpaces: "\ud64d\uae38\ub3d9",
			formalName:         "\ud64d\uae38\ub3d9",
		},
		{
			name:               "korean name, decomposed", // Hong Gildong
			input:              "\u1112\u1169\u11bc\u1100\u1175\u11af\u1103\u1169\u11bc",
			alpha:              "\u1112\u1169\u11bc\u1100\u1175\u11af\u1103\u1169\u11bc",
			alphaSpaces:        "\u1112\u1169\u11bc\u1100\u1175\u11af\u1103\u1169\u11bc",
			alphaNumeric:       "\u1112\u1169\u11bc\u1100\u1175\u11af\u1103\u1169\u11bc",
			alphaNumericSpaces: "\u1112\u1169\u11bc\u1100\u1175\u11af\u1103\u1169\u11bc",
			formalName:         "\u1112\u1169\u11bc\u1100\u1175\u11af\u1103\u1169\u11bc",
		},
		{
			name:               "hangul filler",
			input:              "\u3164Kim",
			alpha:              "\u3164Kim",
			alphaSpaces:        "\u3164Kim",
			alphaNumeric:       "\u3164Kim",
			alphaNumericSpaces: "\u3164Kim",
			formalName:         "\u3164Kim",
		},
		{
			name:               "ethiopic name", // Haile Gebrselassie
			input:              "\u1283\u12ed\u120c \u1308\u1265\u1228\u1225\u120b\u1234",
			alpha:              "\u1283\u12ed\u120c\u1308\u1265\u1228\u1225\u120b\u1234",
			alphaSpaces:        "\u1283\u12ed\u120c \u1308\u1265\u1228\u1225\u120b\u1234",
			alphaNumeric:       "\u1283\u12ed\u120c\u1308\u1265\u1228\u1225\u120b\u1234",
			alphaNumericSpaces: "\u1283\u12ed\u120c \u1308\u1265\u1228\u1225\u120b\u1234",
			formalName:         "\u1283\u12ed\u120c \u1308\u1265\u1228\u1225\u120b\u1234",
		},
		{
			name:               "mongolian with a free variation selector",
			input:              "\u1820\u180b\u1821",
			alpha:              "\u1820\u1821",
			alphaSpaces:        "\u1820\u1821",
			alphaNumeric:       "\u1820\u1821",
			alphaNumericSpaces: "\u1820\u1821",
			formalName:         "\u1820\u1821",
		},
		{
			name:               "fullwidth letters and digits",
			input:              "\uff21\uff22\uff23\uff11\uff12\uff13",
			alpha:              "\uff21\uff22\uff23",
			alphaSpaces:        "\uff21\uff22\uff23",
			alphaNumeric:       "\uff21\uff22\uff23\uff11\uff12\uff13",
			alphaNumericSpaces: "\uff21\uff22\uff23\uff11\uff12\uff13",
			formalName:         "\uff21\uff22\uff23\uff11\uff12\uff13",
		},
		{
			name:               "superscripts fractions and roman numerals",
			input:              "x\u00b2 \u00bd \u2162",
			alpha:              "x",
			alphaSpaces:        "x  ",
			alphaNumeric:       "x",
			alphaNumericSpaces: "x  ",
			formalName:         "x  ",
		},
		{
			name:               "emoji with a skin tone",
			input:              "Ana \U0001f44d\U0001f3fd",
			alpha:              "Ana",
			alphaSpaces:        "Ana ",
			alphaNumeric:       "Ana",
			alphaNumericSpaces: "Ana ",
			formalName:         "Ana ",
		},
		{
			name:               "emoji presentation selector",
			input:              "Love\u2764\ufe0f",
			alpha:              "Love",
			alphaSpaces:        "Love",
			alphaNumeric:       "Love",
			alphaNumericSpaces: "Love",
			formalName:         "Love",
		},
		{
			name:               "keycap sequence",
			input:              "Room 1\ufe0f\u20e3",
			alpha:              "Room",
			alphaSpaces:        "Room ",
			alphaNumeric:       "Room1",
			alphaNumericSpaces: "Room 1",
			formalName:         "Room 1",
		},
		{
			name:               "zero-width joiner emoji family",
			input:              "Ty\U0001f468\u200d\U0001f469\u200d\U0001f467",
			alpha:              "Ty",
			alphaSpaces:        "Ty",
			alphaNumeric:       "Ty",
			alphaNumericSpaces: "Ty",
			formalName:         "Ty",
		},
		{
			name:               "flag emoji",
			input:              "Uma\U0001f1fa\U0001f1f8",
			alpha:              "Uma",
			alphaSpaces:        "Uma",
			alphaNumeric:       "Uma",
			alphaNumericSpaces: "Uma",
			formalName:         "Uma",
		},
		{
			name:               "emoji tag sequence",
			input:              "Val\U0001f3f4\U000e0067\U000e0062\U000e007f",
			alpha:              "Val",
			alphaSpaces:        "Val",
			alphaNumeric:       "Val",
			alphaNumericSpaces: "Val",
			formalName:         "Val",
		},
		{
			name:               "currency and math symbols",
			input:              "Wes $5 \u20ac6 \u00a57 \u00b18 \u221e",
			alpha:              "Wes",
			alphaSpaces:        "Wes     ",
			alphaNumeric:       "Wes5678",
			alphaNumericSpaces: "Wes 5 6 7 8 ",
			formalName:         "Wes 5 6 7 8 ",
		},
		{
			name:               "zero-width space",
			input:              "Xena\u200bYork",
			alpha:              "XenaYork",
			alphaSpaces:        "XenaYork",
			alphaNumeric:       "XenaYork",
			alphaNumericSpaces: "XenaYork",
			formalName:         "XenaYork",
		},
		{
			name:               "byte order mark",
			input:              "\ufeffYuri",
			alpha:              "Yuri",
			alphaSpaces:        "Yuri",
			alphaNumeric:       "Yuri",
			alphaNumericSpaces: "Yuri",
			formalName:         "Yuri",
		},
		{
			name:               "soft hyphen",
			input:              "Zane\u00adAbel",
			alpha:              "ZaneAbel",
			alphaSpaces:        "ZaneAbel",
			alphaNumeric:       "ZaneAbel",
			alphaNumericSpaces: "ZaneAbel",
			formalName:         "ZaneAbel",
		},
		{
			name:               "bidirectional override",
			input:              "Bree\u202eCole\u202c",
			alpha:              "BreeCole",
			alphaSpaces:        "BreeCole",
			alphaNumeric:       "BreeCole",
			alphaNumericSpaces: "BreeCole",
			formalName:         "BreeCole",
		},
		{
			name:               "left-to-right mark",
			input:              "Cara\u200eDale",
			alpha:              "CaraDale",
			alphaSpaces:        "CaraDale",
			alphaNumeric:       "CaraDale",
			alphaNumericSpaces: "CaraDale",
			formalName:         "CaraDale",
		},
		{
			name:               "combining grapheme joiner",
			input:              "Dora\u034fEads",
			alpha:              "DoraEads",
			alphaSpaces:        "DoraEads",
			alphaNumeric:       "DoraEads",
			alphaNumericSpaces: "DoraEads",
			formalName:         "DoraEads",
		},
		{
			name:               "khmer inherent vowel signs",
			input:              "Ella\u17b4\u17b5Ford",
			alpha:              "EllaFord",
			alphaSpaces:        "EllaFord",
			alphaNumeric:       "EllaFord",
			alphaNumericSpaces: "EllaFord",
			formalName:         "EllaFord",
		},
		{
			name:               "accent after a letter",
			input:              "Jose\u0301",
			alpha:              "Jose\u0301",
			alphaSpaces:        "Jose\u0301",
			alphaNumeric:       "Jose\u0301",
			alphaNumericSpaces: "Jose\u0301",
			formalName:         "Jose\u0301",
		},
		{
			name:               "two stacked accents",
			input:              "Nguye\u0302\u0303n",
			alpha:              "Nguye\u0302\u0303n",
			alphaSpaces:        "Nguye\u0302\u0303n",
			alphaNumeric:       "Nguye\u0302\u0303n",
			alphaNumericSpaces: "Nguye\u0302\u0303n",
			formalName:         "Nguye\u0302\u0303n",
		},
		{
			name:               "ten stacked marks",
			input:              "a\u0300\u0301\u0302\u0303\u0304\u0305\u0306\u0307\u0308\u0309b",
			alpha:              "a\u0300\u0301\u0302\u0303\u0304\u0305\u0306\u0307\u0308\u0309b",
			alphaSpaces:        "a\u0300\u0301\u0302\u0303\u0304\u0305\u0306\u0307\u0308\u0309b",
			alphaNumeric:       "a\u0300\u0301\u0302\u0303\u0304\u0305\u0306\u0307\u0308\u0309b",
			alphaNumericSpaces: "a\u0300\u0301\u0302\u0303\u0304\u0305\u0306\u0307\u0308\u0309b",
			formalName:         "a\u0300\u0301\u0302\u0303\u0304\u0305\u0306\u0307\u0308\u0309b",
		},
		{
			name:               "leading mark",
			input:              "\u0301Abe",
			alpha:              "Abe",
			alphaSpaces:        "Abe",
			alphaNumeric:       "Abe",
			alphaNumericSpaces: "Abe",
			formalName:         "Abe",
		},
		{
			name:               "marks only",
			input:              "\u0301\u0302\u0303",
			alpha:              "",
			alphaSpaces:        "",
			alphaNumeric:       "",
			alphaNumericSpaces: "",
			formalName:         "",
		},
		{
			name:               "mark after a digit",
			input:              "Unit 1\u0301 and 2\u0302",
			alpha:              "Unitand",
			alphaSpaces:        "Unit  and ",
			alphaNumeric:       "Unit1\u0301and2\u0302",
			alphaNumericSpaces: "Unit 1\u0301 and 2\u0302",
			formalName:         "Unit 1\u0301 and 2\u0302",
		},
		{
			name:               "mark after a space",
			input:              "Ben \u0301Cruz",
			alpha:              "BenCruz",
			alphaSpaces:        "Ben Cruz",
			alphaNumeric:       "BenCruz",
			alphaNumericSpaces: "Ben Cruz",
			formalName:         "Ben Cruz",
		},
		{
			name:               "mark after an apostrophe",
			input:              "O'\u0301Dell O\u2019\u0301Dell",
			alpha:              "ODellODell",
			alphaSpaces:        "ODell ODell",
			alphaNumeric:       "ODellODell",
			alphaNumericSpaces: "ODell ODell",
			formalName:         "O'Dell O\u2019Dell",
		},
		{
			name:               "mark after a hyphen and a comma",
			input:              "Lee-\u0301Ann,\u0301Jo",
			alpha:              "LeeAnnJo",
			alphaSpaces:        "LeeAnnJo",
			alphaNumeric:       "LeeAnnJo",
			alphaNumericSpaces: "LeeAnnJo",
			formalName:         "Lee-Ann,Jo",
		},
		{
			name:               "mark after a removed symbol",
			input:              "Ava#\u0301Bell",
			alpha:              "AvaBell",
			alphaSpaces:        "AvaBell",
			alphaNumeric:       "AvaBell",
			alphaNumericSpaces: "AvaBell",
			formalName:         "AvaBell",
		},
		{
			name:               "mark after a tab",
			input:              "Cy\t\u0301Dunn",
			alpha:              "CyDunn",
			alphaSpaces:        "CyDunn",
			alphaNumeric:       "CyDunn",
			alphaNumericSpaces: "CyDunn",
			formalName:         "Cy\tDunn",
		},
		{
			name:               "mark after a no-break space",
			input:              "Di\u00a0\u0301Eck",
			alpha:              "DiEck",
			alphaSpaces:        "DiEck",
			alphaNumeric:       "DiEck",
			alphaNumericSpaces: "DiEck",
			formalName:         "Di\u00a0Eck",
		},
		{
			name:               "variation selector between a letter and its accent",
			input:              "Ce\ufe0f\u0301dric",
			alpha:              "Cedric",
			alphaSpaces:        "Cedric",
			alphaNumeric:       "Cedric",
			alphaNumericSpaces: "Cedric",
			formalName:         "Cedric",
		},
		{
			name:               "enclosing circle after a letter",
			input:              "Fay\u20dd",
			alpha:              "Fay",
			alphaSpaces:        "Fay",
			alphaNumeric:       "Fay",
			alphaNumericSpaces: "Fay",
			formalName:         "Fay",
		},
		{
			name:               "enclosing marks after digits",
			input:              "Gus 5\u20e3 6\u20de",
			alpha:              "Gus",
			alphaSpaces:        "Gus  ",
			alphaNumeric:       "Gus56",
			alphaNumericSpaces: "Gus 5 6",
			formalName:         "Gus 5 6",
		},
		{
			name:               "enclosing cyrillic millions sign",
			input:              "Hal\u0488",
			alpha:              "Hal",
			alphaSpaces:        "Hal",
			alphaNumeric:       "Hal",
			alphaNumericSpaces: "Hal",
			formalName:         "Hal",
		},
		{
			name:               "spacing vowel sign after a space",
			input:              "Ida \u093e",
			alpha:              "Ida",
			alphaSpaces:        "Ida ",
			alphaNumeric:       "Ida",
			alphaNumericSpaces: "Ida ",
			formalName:         "Ida ",
		},
		{
			name:               "zalgo text",
			input:              "Z\u0300\u0305\u030a\u030f\u0314\u0319a\u030b\u0310\u0315\u031a\u031f\u0324l\u0316\u031b\u0320\u0325\u032a\u032fg\u0321\u0326\u032b\u0330\u0335\u033ao\u032c\u0331\u0336\u033b\u0340\u0345",
			alpha:              "Z\u0300\u0305\u030a\u030f\u0314\u0319a\u030b\u0310\u0315\u031a\u031f\u0324l\u0316\u031b\u0320\u0325\u032a\u032fg\u0321\u0326\u032b\u0330\u0335\u033ao\u032c\u0331\u0336\u033b\u0340\u0345",
			alphaSpaces:        "Z\u0300\u0305\u030a\u030f\u0314\u0319a\u030b\u0310\u0315\u031a\u031f\u0324l\u0316\u031b\u0320\u0325\u032a\u032fg\u0321\u0326\u032b\u0330\u0335\u033ao\u032c\u0331\u0336\u033b\u0340\u0345",
			alphaNumeric:       "Z\u0300\u0305\u030a\u030f\u0314\u0319a\u030b\u0310\u0315\u031a\u031f\u0324l\u0316\u031b\u0320\u0325\u032a\u032fg\u0321\u0326\u032b\u0330\u0335\u033ao\u032c\u0331\u0336\u033b\u0340\u0345",
			alphaNumericSpaces: "Z\u0300\u0305\u030a\u030f\u0314\u0319a\u030b\u0310\u0315\u031a\u031f\u0324l\u0316\u031b\u0320\u0325\u032a\u032fg\u0321\u0326\u032b\u0330\u0335\u033ao\u032c\u0331\u0336\u033b\u0340\u0345",
			formalName:         "Z\u0300\u0305\u030a\u030f\u0314\u0319a\u030b\u0310\u0315\u031a\u031f\u0324l\u0316\u031b\u0320\u0325\u032a\u032fg\u0321\u0326\u032b\u0330\u0335\u033ao\u032c\u0331\u0336\u033b\u0340\u0345",
		},
		{
			name:               "mixed scripts", // Rohit Rohit Vladimir Wang
			input:              "Rohit \u0930\u094b\u0939\u093f\u0924 \u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u738b",
			alpha:              "Rohit\u0930\u094b\u0939\u093f\u0924\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440\u738b",
			alphaSpaces:        "Rohit \u0930\u094b\u0939\u093f\u0924 \u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u738b",
			alphaNumeric:       "Rohit\u0930\u094b\u0939\u093f\u0924\u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440\u738b",
			alphaNumericSpaces: "Rohit \u0930\u094b\u0939\u093f\u0924 \u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u738b",
			formalName:         "Rohit \u0930\u094b\u0939\u093f\u0924 \u0412\u043b\u0430\u0434\u0438\u043c\u0438\u0440 \u738b",
		},
		{
			name:               "latin and cyrillic look-alikes", // paypal with a cyrillic a
			input:              "p\u0430ypal",
			alpha:              "p\u0430ypal",
			alphaSpaces:        "p\u0430ypal",
			alphaNumeric:       "p\u0430ypal",
			alphaNumericSpaces: "p\u0430ypal",
			formalName:         "p\u0430ypal",
		},
		{
			name:               "overlay mark that changes how a letter looks",
			input:              "pa\u0336ypal",
			alpha:              "pa\u0336ypal",
			alphaSpaces:        "pa\u0336ypal",
			alphaNumeric:       "pa\u0336ypal",
			alphaNumericSpaces: "pa\u0336ypal",
			formalName:         "pa\u0336ypal",
		},
		{
			name:               "dotless i with a dot above",
			input:              "adm\u0131\u0307n",
			alpha:              "adm\u0131\u0307n",
			alphaSpaces:        "adm\u0131\u0307n",
			alphaNumeric:       "adm\u0131\u0307n",
			alphaNumericSpaces: "adm\u0131\u0307n",
			formalName:         "adm\u0131\u0307n",
		},
		{
			name:               "invalid utf-8 byte",
			input:              "Jan\xffKos",
			alpha:              "JanKos",
			alphaSpaces:        "JanKos",
			alphaNumeric:       "JanKos",
			alphaNumericSpaces: "JanKos",
			formalName:         "JanKos",
		},
		{
			name:               "invalid byte before a mark",
			input:              "Kim\xff\u0301Lu",
			alpha:              "KimLu",
			alphaSpaces:        "KimLu",
			alphaNumeric:       "KimLu",
			alphaNumericSpaces: "KimLu",
			formalName:         "KimLu",
		},
		{
			name:               "truncated multibyte sequence",
			input:              "Jos\xc3",
			alpha:              "Jos",
			alphaSpaces:        "Jos",
			alphaNumeric:       "Jos",
			alphaNumericSpaces: "Jos",
			formalName:         "Jos",
		},
	}
}

// TestAlphaUnicodeCorpus tests Alpha, with and without spaces, against every corpus row
func TestAlphaUnicodeCorpus(t *testing.T) {
	for _, test := range unicodeCorpus() {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.alpha, sanitize.Alpha(test.input, false), "without spaces")
			assert.Equal(t, test.alphaSpaces, sanitize.Alpha(test.input, true), "with spaces")
		})
	}
}

// TestAlphaNumericUnicodeCorpus tests AlphaNumeric, with and without spaces, against every corpus row
func TestAlphaNumericUnicodeCorpus(t *testing.T) {
	for _, test := range unicodeCorpus() {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.alphaNumeric, sanitize.AlphaNumeric(test.input, false), "without spaces")
			assert.Equal(t, test.alphaNumericSpaces, sanitize.AlphaNumeric(test.input, true), "with spaces")
		})
	}
}

// TestFormalNameUnicodeCorpus tests FormalName against every corpus row
func TestFormalNameUnicodeCorpus(t *testing.T) {
	for _, test := range unicodeCorpus() {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.formalName, sanitize.FormalName(test.input))
		})
	}
}

// TestNameSanitizersUnicodeCorpusProperties tests the properties every name sanitizer output has, on every corpus row
func TestNameSanitizersUnicodeCorpusProperties(t *testing.T) {
	for _, test := range unicodeCorpus() {
		t.Run(test.name, func(t *testing.T) {
			requireNameSanitizerProperties(t, test.input)
		})
	}
}

// TestNameSanitizersKeepTheInputForm tests that precomposed and decomposed input each come back in their own form
func TestNameSanitizersKeepTheInputForm(t *testing.T) {
	precomposed := "Jos\u00e9 L\u00f3pez"
	decomposed := "Jose\u0301 Lo\u0301pez"
	for _, input := range []string{precomposed, decomposed} {
		assert.Equal(t, input, sanitize.Alpha(input, true))
		assert.Equal(t, input, sanitize.AlphaNumeric(input, true))
		assert.Equal(t, input, sanitize.FormalName(input))
	}
	assert.NotEqual(t, sanitize.FormalName(precomposed), sanitize.FormalName(decomposed), "nothing is normalized")
}

// TestNameSanitizersLongInput tests long input, and a letter that carries a long run of marks
func TestNameSanitizersLongInput(t *testing.T) {
	const repeat = 2000
	names := testDevanagariName + " " + testThaiName + " " + testDecomposedVietnameseName
	input := strings.Repeat(names+"! 42 ", repeat)
	assert.Equal(t, strings.Repeat(names+"  ", repeat), sanitize.Alpha(input, true))
	assert.Equal(t, strings.Repeat(names+" 42 ", repeat), sanitize.AlphaNumeric(input, true))
	assert.Equal(t, strings.Repeat(names+" 42 ", repeat), sanitize.FormalName(input))

	stacked := "a" + strings.Repeat("\u0301", 10000)
	assert.Equal(t, stacked, sanitize.Alpha(stacked, false), "stacked marks are not limited")
	assert.Equal(t, stacked, sanitize.AlphaNumeric(stacked, false), "stacked marks are not limited")
	assert.Equal(t, stacked, sanitize.FormalName(stacked), "stacked marks are not limited")
	assert.Empty(t, sanitize.Alpha("!"+strings.Repeat("\u0301", 10000), false), "a detached run of marks is removed")
}
