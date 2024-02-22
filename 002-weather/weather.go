// Package weather is supposed to predict the weather forecast.
package weather

// CurrentCondition type.
var CurrentCondition string

// CurrentLocation type.
var CurrentLocation string

// Forecast that takes in two parameter, 1) city and 2) condition and return the forcast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
