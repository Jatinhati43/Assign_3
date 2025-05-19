// Description: Provides functionality to count the frequency of
//              alphabetic characters (ignoring digits and special characters)
//              from the contents of a given text file.


package charcount

import (
	"os"
	"unicode"
)

// CountCharacters reads the file at the given path and returns a map
// of character frequencies for all alphabetic characters (case-insensitive).
func CountCharacters(filePath string) (map[rune]int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	counts := make(map[rune]int)
	for _, ch := range data {
		if unicode.IsLetter(rune(ch)) {
			counts[unicode.ToLower(rune(ch))]++
		}
	}

	return counts, nil
}
