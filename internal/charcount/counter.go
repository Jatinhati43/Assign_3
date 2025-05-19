package charcount

import (
	"os"
	"unicode"
)

func CountCharactersFromFile(filePath string) (map[rune]int, error) {
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
