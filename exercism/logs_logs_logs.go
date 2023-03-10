package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	value := Application("β recommended search product π") // => recommendation
	fmt.Println(value)

	log := "please replace 'π' with 'π'"
	fmt.Println(Replace(log, 'π', 'π')) // => please replace 'π' with 'π'"

	fmt.Println(WithinLimit("helloβ", 6)) // => true
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	/*	for _, v := range log {
		switch v {
		case 'β':
			return "recommendation"
		case 'π':
			return "search"
		case 'β':
			return "weather"
		default:
			return "default"
		}
	}*/
	var icons = map[rune]string{
		'β': "recommendation",
		'π': "search",
		'β': "weather",
	}

	for _, v := range log {
		if runes, x := icons[v]; x {
			return runes
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
