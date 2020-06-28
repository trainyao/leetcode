package main

import "github.com/trainyao/leetcode/pkg/queue"

func main() {

}

func gardenNoAdj(N int, paths [][]int) []int {
	q := queue.NewQueue()
	pointConnected := make([][]bool, N)

	for i := 0; i < len(paths); i++ {
		fromIndex := paths[i][0] - 1
		toIndex := paths[i][1] - 1

		if pointConnected[fromIndex] == nil {
			pointConnected[fromIndex] = make([]bool, N)
		} else {
			pointConnected[fromIndex][toIndex] = true
		}

		if pointConnected[toIndex] == nil {
			pointConnected[toIndex] = make([]bool, N)
		} else {
			pointConnected[toIndex][fromIndex] = true
		}
	}

	res := make([]int, N)

	q.Enqueue(1)
	res[0] = 1

	var pointInterf interface{}
	for pointInterf = q.Equeue(); pointInterf != nil; {
		point := pointInterf.(int) - 1
		// if 
		if res[point] != 0 {


		}
	}
	return res
}
