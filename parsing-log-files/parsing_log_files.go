package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	pattern := `\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\]`
	re := regexp.MustCompile(pattern)
	locationSlice := re.FindStringIndex(text)
	if locationSlice == nil {
		return false
	}

	return locationSlice[0] == 0
}

func SplitLogLine(text string) []string {
	pattern := `<[~*=\-]*>`
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `"[^"]*password[^"]*"`
	re := regexp.MustCompile("(?i)" + pattern)
	counter := 0

	for _, input := range lines {
		if re.MatchString(input) {
			counter++
		}
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d+`
	re := regexp.MustCompile(pattern)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+([A-Za-z0-9]+)`
	re := regexp.MustCompile(pattern)
	var result []string

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		output := line
		if len(match) > 1 {
			output = "[USR] " + match[1] + " " + line
		}

		result = append(result, output)
	}

	return result
}
