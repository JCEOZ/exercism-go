// Package luhn contains the solution for Luhn exercise from https://exercism.org/tracks/go/exercises/luhn
package luhn

import (
	"unicode"
)

// Valid determines if provided string is valid for the Luhn formula
func Valid(id string) bool {
	numbers := convertToNumbers(id)

	if len(numbers) < 2 {
		return false
	}

	var sum int
	parity := len(numbers) % 2

	for i := 0; i < len(numbers); i++ {
		if (i)%2 != parity {
			sum += numbers[i]
		} else if numbers[i] > 4 {
			sum += 2*numbers[i] - 9
		} else {
			sum += 2 * numbers[i]
		}
	}

	return sum%10 == 0
}

func convertToNumbers(id string) []int {
	numbers := make([]int, 0)

	for _, r := range id {
		if unicode.IsNumber(r) {
			numbers = append(numbers, int(r-'0'))
		} else if !unicode.IsSpace(r) {
			return make([]int, 0)
		}
	}

	return numbers
}
