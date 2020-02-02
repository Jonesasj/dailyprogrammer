package main

import (
	"fmt"
)

func main() {

	roll := [5]int{6, 6, 6, 6, 6}
	max := maxScore(roll)
	fmt.Println(max)
}

//returns the maximum score for this roll
func maxScore(roll [5]int) int {
	counts := map[int]int{}
	for _, value := range roll {
		counts[value]++
	}

	scores := map[int]int{}

	for key, value := range counts {
		scores[key] = key * value
	}

	max := 0
	for _, value := range scores {
		if value > max {
			max = value
		}
	}
	return max
}
