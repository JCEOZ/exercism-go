// Package hamming contains the solution for Hamming exercise from https://exercism.org/tracks/go/exercises/hamming
package hamming

import "errors"

// Distance calculates the distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("provided sequences must have equal length")
	}

	var distance int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
