package charcount

import (
	"reflect"
	"testing"
)

func TestCountChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[rune]int
	}{
		{
			name:  "Empty string",
			input: "",
			expected: map[rune]int{},
		},
		{
			name:  "Letters only",
			input: "aAbB",
			expected: map[rune]int{'a': 2, 'b': 2},
		},
		{
			name:  "Ignore digits and special chars",
			input: "a1!b2@a#",
			expected: map[rune]int{'a': 2, 'b': 1},
		},
		{
			name:  "Mixed case and spaces",
			input: "Go Lang 2025!",
			expected: map[rune]int{'g': 2, 'o': 1, 'l': 1, 'a': 1, 'n': 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountChars(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("CountChars(%q) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
