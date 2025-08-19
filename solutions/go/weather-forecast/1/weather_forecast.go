//Package weather provides the current weather condition of a city.
package weather

//CurrentCondition represents the weather condition of a location.
var CurrentCondition string
//CurrentLocation represents the name of the location.
var CurrentLocation string

//Forecast returns a string value of the current weather condition of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
