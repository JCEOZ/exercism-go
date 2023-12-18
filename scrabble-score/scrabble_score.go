// Package scrabble contains the solution for Scrabble Score exercise from https://exercism.org/tracks/go/exercises/scrabble-score
package scrabble

import "strings"

// Score calculates the Scrabble score for the provided word
func Score(word string) int {
	score := 0

	for _, r := range strings.ToUpper(word) {
		score += letterValue(r)
	}

	return score
}

func letterValue(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
