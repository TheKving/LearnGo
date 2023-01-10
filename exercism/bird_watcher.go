package birdwatcher

import "fmt"

func main() {

	birdsPerDay := []int{4, 7, 3, 2, 1, 1, 2, 0, 2, 3, 2, 7, 1, 3, 0, 6, 5, 3, 7, 2, 3}
	//fmt.Println(len(birdsPerDay))
	//fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println(BirdsInWeek(birdsPerDay, 2))

	birdsPerDay1 := []int{2, 5, 0, 7, 4, 1}
	fmt.Println(FixBirdCountLog(birdsPerDay1))
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sumBirds int = 0
	for _, numBirds := range birdsPerDay {
		sumBirds += numBirds
	}
	return sumBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	/*offset := 7 * (week - 1)
	return TotalBirdCount(birdsPerDay[offset : offset+7])*/
	var sumBirds int = 0

	for index, numBirds := range birdsPerDay {
		if week == 1 {
			if index <= 7 {
				sumBirds += numBirds
			}
		} else if week == 2 {
			if index >= 7 && index <= 14 {
				sumBirds += numBirds
			}
		}
	}

	return sumBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, _ := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index]++
		}
	}
	return birdsPerDay
}
