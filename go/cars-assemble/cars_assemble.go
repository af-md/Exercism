package cars

// import (
// 	"fmt"
// )

// things to learn: printing values using their right format? what are the right ways of dealing with floats and success rates?

const costOfTenCars int = 95000
const costOfOneCar int = 10000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if successRate == 100.0 {
		return float64(productionRate)
	}
	var p float64 = float64(productionRate) / 10
	var result = p * (successRate / 10)
	return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var pHCars float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
	var cars int = int(pHCars)
	var carMin int = cars / 60
	return carMin
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return uint(carsCount/10*costOfTenCars + carsCount%10*costOfOneCar)
}
