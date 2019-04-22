package main

import (
	"log"
)

func main() {
	log.Printf("result is %d", maxArea([]int{1,2,3,4,5}))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	var left int = 0
	var right int = len(height) - 1
	var max int = 0

	for left < right {
		var minLength = 0
		var distance = right - left
		if height[left] < height[right] {
			minLength = height[left]
			left += 1
		} else {
			minLength = height[right]
			right -= 1
		}
		var area int = minLength * distance
		if area > max {
			max = area
		}
	}

	return max
}
