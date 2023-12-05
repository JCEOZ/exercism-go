// Package weather provides the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents current weather conditions in provided location.
var CurrentCondition string

// CurrentLocation represents a location for which the weather condition should be calculated.
var CurrentLocation string

// Forecast returns a current weather condition for provided city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
