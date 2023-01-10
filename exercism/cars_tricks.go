package cars

import "fmt"

func main() {
	fmt.Println(NeedsLicense("car"))  // => True
	fmt.Println(NeedsLicense("bike")) // => Flase

	//vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla") // => "Toyota Corolla is clearly the better choice."
	fmt.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")) // => "Vw Beetke is clearly the better choice."
	fmt.Println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))

	fmt.Println(CalculateResellPrice(1000, 1))  // 800
	fmt.Println(CalculateResellPrice(1000, 5))  // 700
	fmt.Println(CalculateResellPrice(1000, 15)) // => 500

}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	} else {
		return option2 + " is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var valueCar float64
	if age < 3 {
		valueCar = originalPrice * 0.8
	} else if age >= 10 {
		valueCar = originalPrice * 0.5
	} else if age >= 3 {
		valueCar = originalPrice * 0.7
	}
	return valueCar
}
