package cars

import "fmt"

func main() {
	fmt.Println("Hello")

	fmt.Println(CalculateWorkingCarsPerHour(1547, 90)) // 1392.3

	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90)) // 16

	fmt.Println(CalculateCost(37)) // => 355000
	fmt.Println(CalculateCost(21)) // => 200000

}

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64((float64(productionRate) * (successRate / 100)))
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((float64(productionRate) * successRate / 100) / 60)
}

func CalculateCost(carsCount int) uint {
	return uint((carsCount/10)*95000 + (carsCount%10)*10000)
}
