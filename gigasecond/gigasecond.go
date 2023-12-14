// Package gigasecond contains the solution for Gigasecond exercise from https://exercism.org/tracks/go/exercises/gigasecond
package gigasecond

import "time"

const gigaSecond = 1000000000

// AddGigasecond calculates a date which is one giga second after provided date.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigaSecond) * time.Second)
}
