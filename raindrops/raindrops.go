// Package raindrops contains the solution for Raindrops exercise from https://exercism.org/tracks/go/exercises/raindrops
package raindrops

import "strconv"

// Convert converts a number into string containing raindrop sound
func Convert(number int) string {
	var result string

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
