package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	runes := []rune(log)

	for _, logRune := range runes {
		switch logRune {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for index, logRune := range runes {
		if logRune == oldRune {
			runes[index] = newRune
		}
	}

	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runes := []rune(log)
	return len(runes) <= limit
}
