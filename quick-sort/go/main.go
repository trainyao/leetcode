package main

import (
	"log"
)

func main() {
	res, _ := quickSort([]int{2, 1, 1, 1})
	log.Printf("%a", res)
	res, _ = quickSort([]int{1, 2})
	log.Printf("%a", res)
	res, _ = quickSort([]int{2, 1})
	log.Printf("%a", res)
	res, _ = quickSort([]int{1, 2, 1, 1})
	log.Printf("%a", res)
	res, _ = quickSort([]int{2, 1, 7, 3})
	log.Printf("%a", res)
	res, _ = quickSort([]int{1, 1, 7, 3})
	log.Printf("%a", res)
	res, _ = quickSort([]int{1, 2, 3, 4})
	log.Printf("%a", res)
	res, _ = quickSort([]int{89, 34, 32, 2, 433, 44, 22, 99, 4324, 43, 23, 23, 43333, 6, 1})
	log.Printf("%a", res)
	res, _ = quickSort([]int{5, 1, 8, 9, 10, 4, 3, 2, 2, 1})
	log.Printf("%a", res)
	res, _ = quickSort([]int{1, 1, 1})
	log.Printf("%a", res)
	res, _ = quickSort([]int{1})
	log.Printf("%a", res)
	res, _ = quickSort([]int{})
	log.Printf("%a", res)
}

func quickSort(a []int) ([]int, error) {
	if len(a) <= 1 {
		return a, nil
	}
	if len(a) == 2 && a[0] > a[1] {
		a[0], a[1] = a[1], a[0]
		return a, nil
	}

	benchmarkValueIndex := 0
	benchmarkValue := a[benchmarkValueIndex]
	for index, value := range a {
		// 跨过第一个元素
		if index == 0 {
			continue
		}

		if value < benchmarkValue {
			// 如果后面的元素小于第一个元素 交换位置
			a[index], a[benchmarkValueIndex] = a[benchmarkValueIndex], a[index]

			// 保持第一个元素前面的都是小于它的元素 后面都是大于它的元素
			// 如果发现交换后放到了较大元素的后面,于是和较大元素的第一个交换一下位置
			// 如: 基准元素=5, 1[5]987[2]3 => 12[9]87[5]3 => 12[5]87[9]3
			if benchmarkValueIndex+1 < len(a) && benchmarkValueIndex+1 != index {
				a[benchmarkValueIndex+1], a[index] = a[index], a[benchmarkValueIndex+1]
				benchmarkValueIndex = benchmarkValueIndex + 1
			} else {
				benchmarkValueIndex = index
			}
		}
	}

	// 递归前面部分
	_, err := quickSort(a[:benchmarkValueIndex])
	if err != nil {
		return a, err
	}
	// 递归后面部分
	if benchmarkValueIndex+1 < len(a) {
		_, err := quickSort(a[benchmarkValueIndex+1:])
		if err != nil {
			return a, err
		}
	}

	return a, nil
}
