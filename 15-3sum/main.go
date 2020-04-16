package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {

	log.Printf("%+v", threeSum([]int{0, -1, 1, -1, -1, 2}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	ex := map[string]interface{}{}

	// dict
	store := map[int]int{}
	for i, v := range nums {
		store[v] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			target := -(nums[i] + nums[j])
			if index, ok := store[target]; ok && index > j {
				sign := fmt.Sprintf("%d%d%d", nums[i], nums[j], target)
				if _, ok := ex[sign]; ok {
					continue
				}
				ex[sign] = struct {}{}

				result = append(result, []int{nums[i], nums[j], target})
			}
		}
	}

	return result
}
