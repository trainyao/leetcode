package main

import (
	"log"
)

func main() {
	log.Printf("result is %d", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	var max int = 0

	for index, _ := range height {
		if result := maxAreaForOneRound(height[index:]); result > max {
			max = result
		}
	}

	return max
}

func maxAreaForOneRound(height []int) int {
	var max int = 0

	for index, _ := range height {
		if index == 0 {
			continue
		}

		var smallerHeight int = height[0]
		if height[0] > height[index] {
			smallerHeight = height[index]
		}
		if result := smallerHeight * index; result > max {
			max = result
		}
	}
	return max
}
