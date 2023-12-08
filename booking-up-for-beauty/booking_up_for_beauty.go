package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return parseStringDateToTime(date, "1/2/2006 15:04:05")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	return now.After(parseStringDateToTime(date, "January 2, 2006 15:04:05"))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate := parseStringDateToTime(date, "Monday, January 2, 2006 15:04:05")
	return parsedDate.Hour() >= 12 && parsedDate.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedDate := parseStringDateToTime(date, "1/2/2006 15:04:05")
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", parsedDate.Weekday(), parsedDate.Month(), parsedDate.Day(), parsedDate.Year(), parsedDate.Hour(), parsedDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

func parseStringDateToTime(date, layout string) time.Time {
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate
}
