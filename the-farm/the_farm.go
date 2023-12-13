package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numberOfCows  int
	customMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.customMessage)
}

func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	fodderAmount, err := fodderCalculator.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount / float64(numberOfCows) * fatteningFactor, nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	switch {
	case numberOfCows <= 0:
		return 0, errors.New("invalid number of cows")
	default:
		return DivideFood(fodderCalculator, numberOfCows)
	}
}

func ValidateNumberOfCows(numberOfCows int) error {
	switch {
	case numberOfCows < 0:
		return &InvalidCowsError{numberOfCows, "there are no negative cows"}
	case numberOfCows == 0:
		return &InvalidCowsError{numberOfCows, "no cows don't need food"}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
