package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var result int
	for i := 0; i < len(birdsPerDay); i++ {
		result += birdsPerDay[i]
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sliceEndIndex := week * 7
	sliceStartIndex := sliceEndIndex - 7
	slice := birdsPerDay[sliceStartIndex:sliceEndIndex]
	return TotalBirdCount(slice)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return birdsPerDay
}
