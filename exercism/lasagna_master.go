package lasagna

import "fmt"

func main() {

	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 0))

	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))
	// => 100, 0.4

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)
	// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}

	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)

	quantities1 := []float64{0.6, 300, 1, 0.5, 50, 0.1, 100}
	scaledQuantities1 := ScaleRecipe(quantities1, 3)
	fmt.Println(scaledQuantities1)
	//[]float64{0.6, 300, 1, 0.5, 50, 0.1, 100}
	//[]float64{0.9, 450, 1.5, 0.75, 75, 0.15, 150}

	resultado := 3 / 2.0
	fmt.Println(resultado)
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timeLayer int) int {
	if timeLayer <= 0 {
		return len(layers) * 2
	}
	return len(layers) * timeLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, typeLayer := range layers {
		if typeLayer == "noodles" {
			noodles++
		} else if typeLayer == "sauce" {
			sauce++
		}

	}
	return noodles * 50, sauce * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	/*lastElementMyList := len(myList) - 1
	lastElementFriendsList := len(friendsList) - 1*/
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	//(*myList)[len(*myList)-1] = (*friendsList)[len(*friendsList)-1]

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, layers int) (quantitiesScaled []float64) {
	for _, value := range quantities {
		var newValue float64 = float64(layers) / 2.0
		quantitiesScaled = append(quantitiesScaled, value*newValue)

	}
	return quantitiesScaled
}
