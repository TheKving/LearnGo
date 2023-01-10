package chance

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	SeedWithTime()
	fmt.Println(RollADie())
}

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return 0.0 + rand.Float64()*(12.0-0.0)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	sliceAnimals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	//var newAnimals string
	rand.Shuffle(len(sliceAnimals), func(i, j int) {
		sliceAnimals[i], sliceAnimals[j] = sliceAnimals[j], sliceAnimals[i]
	})
	return sliceAnimals
}
