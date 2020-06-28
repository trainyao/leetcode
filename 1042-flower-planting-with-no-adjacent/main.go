package main

import "github.com/trainyao/leetcode/pkg/queue"

func main() {

}

func gardenNoAdj(N int, paths [][]int) []int {
	q := queue.NewQueue()
	pathMapped := map[int][][]int{}

	// use paths's from point to map paths
	for i := 0; i < len(paths); i++ {
		from := paths[i][0]
		if p, ok := pathMapped[from]; !ok {
			pathMapped[from] = make([][]int, 0)
		} else {
			pathMapped[from] = append(p, paths[i])
		}
	}

	res := make([]int, N)
	candidate := [][4]bool{}

	q.Enqueue(1)

	var pointInterf interface{}
	for pointInterf = q.Equeue(); pointInterf != nil; {
		point := pointInterf.(int)

		// if has paths, handle it, enqueue destinations
		if c, ok := pathMapped[point]; ok {
			

		}
	}
	return res
}
