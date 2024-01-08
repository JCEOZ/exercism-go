// Package diffsquares contains the solution for Difference of Squares exercise from https://exercism.org/tracks/go/exercises/difference-of-squares
package diffsquares

// SquareOfSum returns square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	var sum int

	for i := 1; i < n+1; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares returns sum of squares of first n natural number
func SumOfSquares(n int) int {
	var result int

	for i := 1; i < n+1; i++ {
		result += i * i
	}

	return result
}

// Difference returns difference between square of sum and sum of squares for first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
