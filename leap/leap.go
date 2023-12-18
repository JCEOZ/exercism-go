// Package leap contains the solution for Leap exercise from https://exercism.org/tracks/go/exercises/leap
package leap

// IsLeapYear check if provided year is leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
