// Package twofer contains the solution for Two Fer exercise from https://exercism.org/tracks/go/exercises/two-fer
package twofer

// ShareWith generates the phrase which you say when you share a cookie with someone.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
