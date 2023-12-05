package cars

const productionCostPerCar uint = 10000
const productionCostPerGroup uint = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) / 60 * successRate / 100)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var numberOfFullGroups = uint(carsCount) / 10
	var remainingCars = uint(carsCount) % 10
	return numberOfFullGroups*productionCostPerGroup + remainingCars*productionCostPerCar
}
