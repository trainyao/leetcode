package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("%+v\n", fourSum([]int{2, 2, 2, 2, 2, 2, 2}, 8))
	fmt.Printf("%+v\n", fourSum([]int{1, 2, 3, 4, 2, 2, 2}, 10))
	fmt.Printf("%+v\n", fourSum([]int{1, 2, 3, 4, 2, 10, 0, -2, 2}, 10))
}

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	tab := make(map[string]bool, 0)
	if len(nums) < 4 {
		return res
	}
	l := len(nums)
	slices.Sort(nums)
	for i := 0; i < l-3; i++ { // i < l-3,只循环4个数的组合,剪枝
		// i+1后两个数一样，剪枝
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 如果i往下数4个数相加已经 》 target，i已经不用往下算，剪枝
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// i加上最大的3个数，还《target，i不用往下算了，剪枝
		if nums[i]+nums[l-3]+nums[l-2]+nums[l-1] < target {
			continue
		}
		for j := i + 1; j < l-2; j++ { // j := i+1,不用重复算前面的i，剪枝
			// 同i，剪枝
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 同i，剪枝
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// 同i，剪枝
			if nums[i]+nums[j]+nums[l-2]+nums[l-1] < target {
				continue
			}
			o, p := 0, l-1
			tmp := nums[i] + nums[j]
			for {
				if o >= p {
					break
				}
				if i == o || i == p ||
					j == o || j == p ||
					o == p {
					o++
					continue
				}
				tmpSum := tmp + nums[o] + nums[p]
				if tmpSum == target {
					tmp2 := []int{nums[i], nums[j], nums[o], nums[p]}
					slices.Sort(tmp2)
					key := fmt.Sprintf("%d|%d|%d|%d", tmp2[0], tmp2[1], tmp2[2], tmp2[3])
					if _, ok := tab[key]; !ok {
						res = append(res, tmp2)
						tab[key] = true
					}
					o++
					continue
				}
				if tmpSum > target {
					p--
					continue
				}
				if tmpSum < target {
					o++
					continue
				}
			}
		}
	}
	return res
}
