package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var result []Record

	for _, record := range in {
		if predicate(record) {
			result = append(result, record)
		}
	}

	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		if record.Day >= p.From && record.Day <= p.To {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		if record.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filteredRecords := Filter(in, ByDaysPeriod(p))
	var result float64

	for _, record := range filteredRecords {
		result += record.Amount
	}

	return result
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filteredByCategory := Filter(in, ByCategory(c))

	if filteredByCategory == nil || len(filteredByCategory) == 0 {
		return 0, errors.New("unknown category " + c)
	}

	filteredByCategoryAndPeriod := Filter(filteredByCategory, ByDaysPeriod(p))

	if filteredByCategoryAndPeriod == nil || len(filteredByCategoryAndPeriod) == 0 {
		return 0, nil
	}

	var result float64

	for _, record := range filteredByCategoryAndPeriod {
		result += record.Amount
	}

	return result, nil
}
