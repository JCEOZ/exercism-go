// Package collatzconjecture contains the solution for Raindrops exercise from https://exercism.org/tracks/go/exercises/collatz-conjecture
package collatzconjecture

import "errors"

// CollatzConjecture calculates number of steps required to solve Collatz Conjecture problem
func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, errors.New("input number must be positive")
	}

	var numberOfSteps int
	for ; n != 1; numberOfSteps++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	return numberOfSteps, nil
}
