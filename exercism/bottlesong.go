/*func main() {
	//bottles := [10]string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	//fmt.Println(bottles[0])

	for i := len(bottles) - 1; i >= 0; i-- {
		fmt.Printf("Ten green bottles hanging on the wall,\nTen green bottles hanging on the wall,\nAnd if one green bottle should accidentally fall,\nThere'll be %s green bottles hanging on the wall.\n\n", bottles[i])
	}
	//fmt.Println(
	Recite(5, 4)
}

// []string
func Recite(startBottles, takeDown int) {
	var song string
	for i := startBottles; startBottles-takeDown > 0; i++ {
		fmt.Println(startBottles)
		fmt.Println(takeDown)
		song = song + fmt.Sprintf("Ten green bottles hanging on the wall,\nTen green bottles hanging on the wall,\nAnd if one green bottle should accidentally fall,\nThere'll be %d green bottles hanging on the wall.\n\n", takeDown)
	}
	fmt.Println(song)
	//return song
}*/