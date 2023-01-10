package cards

import "fmt"

func main() {
	cards := FavoriteCards()
	fmt.Println(cards)

	card := GetItem([]int{1, 2, 4, 1}, 0)  // card == 4
	card1 := GetItem([]int{1, 2, 4, 1}, 4) // card == -1

	fmt.Println(card)
	fmt.Println(card1)
	index := 2
	newCard := 6
	cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println("_______________\nOutput:\t[1 2 6 1]\nREAL:\t", cards)

	index1 := -1
	newCard1 := 6
	cards1 := SetItem([]int{1, 2, 4, 1}, index1, newCard1)
	cards1 := []int{1, 2, 4, 1}

	cards1 = append(cards1, cards1[0:], 88)
	fmt.Println("_______________\nOutput:\t[1 2 4 1 6]\nREAL:\t", cards1)

	slice1 := []int{5, 2, 10, 6, 8, 7, 0, 9}
	cards1 := PrependItems(slice1, 0, 6)
	fmt.Println("_______________\nOutput:\t[0, 6, 5, 2, 10, 6, 8, 7, 0, 9]\nREAL:\t", cards1)

	cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println("_______________\nOutput:\t[3 2 4 8]\nREAL:\t", cards)

	cards1 := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
	fmt.Println("_______________\nOutput:\t[3 2 6 4 8]\nREAL:\t", cards1)

}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice) > index && index >= 0 {
		return slice[index]
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= 0 && len(slice) > index {
		slice[index] = value
		return slice
	}
	return append(slice, value)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if values != nil {
		var newSlice []int
		newSlice = append(newSlice, values...)
		return append(newSlice, slice...)
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice) > index && index >= 0 {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}
