// Package weather the current weather conditions
// of various cities in Goblinocus.
package weather

// CurrentCondition stores a value of type string used
// to communicate current weather conditions.
var CurrentCondition string

// CurrentLocation stores a value of type string user
// to communicate current location for weatherforecast.
var CurrentLocation string

// Forecast returns a value of type string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
