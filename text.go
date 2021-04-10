package myutil

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/logrusorgru/aurora"
)

func TextBlack(s string) string   { return aurora.Sprintf(aurora.Black(s)) }
func TextRed(s string) string     { return aurora.Sprintf(aurora.Red(s)) }
func TextGreen(s string) string   { return aurora.Sprintf(aurora.Green(s)) }
func TextYellow(s string) string  { return aurora.Sprintf(aurora.Yellow(s)) }
func TextBlue(s string) string    { return aurora.Sprintf(aurora.Blue(s)) }
func TextMagenta(s string) string { return aurora.Sprintf(aurora.Magenta(s)) }
func TextCyan(s string) string    { return aurora.Sprintf(aurora.Cyan(s)) }
func TextWhite(s string) string   { return aurora.Sprintf(aurora.White(s)) }

// BreakLongStr adds linebreak to a long string with certain input line length
// also trim by total length at the end
func BreakLongStr(s string, lineLen, totalLen int) string {
	if len(s) <= lineLen {
		return s
	}
	var lineList []string
	for len(s) != 0 {
		var line string
		if len(s) > lineLen {
			line = s[0:lineLen]
			s = s[lineLen:]
		} else {
			line = s[0:]
			s = ""
		}
		lineList = append(lineList, line)
	}
	result := strings.Join(lineList, "\n")
	result = TrimStr(result, totalLen)
	return result
}

// BreakLongParagraph break lines in the paragraphs
func BreakLongParagraph(p string, lineLen, totalLen int) string {
	ss := strings.Split(p, "\n")
	for i, _ := range ss {
		s := strings.TrimSpace(ss[i])
		ss[i] = BreakLongStr(s, lineLen, totalLen)
	}
	return strings.Join(ss, "\n")
}

// TrimStr trims long string to next n characters
func TrimStr(s string, n int) string {
	if n == 0 {
		return s
	}
	if len(s) >= n {
		s = s[0:n] + "..."
	}
	return s
}

// HandleCamalCase helps to split a string with camal case being sticked as a single word
func HandleCamalCase(src string) string {
	var results string
	if !utf8.ValidString(src) {
		return src
	}
	var entries []string
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}
	// handle upper case -> lower case sequences, e.g.
	// "PDFL", "oader" -> "PDF", "Loader"
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, strings.TrimSpace(string(s)))
		}
	}
	results = strings.Join(entries, " ")
	results = strings.ReplaceAll(results, "  ", " ")
	return results
}
