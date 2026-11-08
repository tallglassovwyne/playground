//Package weather provides a function that predicts 
// the weather condition of a particular location.
package weather

var (
	//CurrentCondition represents the climate.
	CurrentCondition string
	//CurrentLocation represents the actual location whose forecast is to be predicted.
	CurrentLocation  string
)

//Forecast predicts the weather condition of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
