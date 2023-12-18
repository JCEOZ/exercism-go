// Package isogram contains the solution for Isogram exercise from https://exercism.org/tracks/go/exercises/isogram
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if provided word is isogram
func IsIsogram(word string) bool {
	characterToOccurrences := make(map[rune]int)

	for _, character := range strings.ToLower(word) {
		if unicode.IsLetter(character) {
			if characterToOccurrences[character] > 0 {
				return false
			}
			characterToOccurrences[character]++
		}
	}

	return true
}
