package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) / 100) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var secondsInMinute, percentage int = 60, 100
	return int((float64(productionRate) / float64(percentage) * successRate) / float64(secondsInMinute))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var numberForTen int = (carsCount/10) * 95000
	var remainder int = (carsCount%10) * 10000

	return uint(numberForTen) + uint(remainder)
}