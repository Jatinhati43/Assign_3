package charcount

import (
	"unicode"
)

// CountChars counts occurrences of each alphabetic character (case insensitive) in input text.
// It ignores digits, punctuation, and special characters.
func CountChars(text string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range text {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r) // Count case-insensitively
			counts[r]++
		}
	}
	return counts
}
