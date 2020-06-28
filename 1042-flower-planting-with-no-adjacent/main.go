package main

func main() {

}

func gardenNoAdj(N int, paths [][]int) []int {
	pointConnected := make([][]int, N)

	for i := 0; i < len(paths); i++ {
		fromIndex := paths[i][0] - 1
		toIndex := paths[i][1] - 1

		if pointConnected[fromIndex] == nil {
			pointConnected[fromIndex] = []int{}
		}

		pointConnected[fromIndex] = append(pointConnected[fromIndex], toIndex)

		if pointConnected[toIndex] == nil {
			pointConnected[toIndex] = []int{}
		}
		pointConnected[toIndex] = append(pointConnected[toIndex], fromIndex)
	}

	// res[i] == 0 representing i node can select color 1
	res := make([]int, N)
	colorBeenChosen := make([][4]bool, N)

	for i := 0; i < N; i++ {
		colorChosen := -1
		if res[i] == 0 {
			colorChosen = 0
		} else {
			for j := 0; j < 4; j++ {
				if !colorBeenChosen[i][j] {
					colorChosen = j
					break
				}
			}
		}
		res[i] = colorChosen + 1

		// set sibling node can not choose j color
		for _, toIndex := range pointConnected[i] {
			colorBeenChosen[toIndex][colorChosen] = true
			// represent i has been modified, can not choose color 1
			if res[toIndex] == 0 {
				res[toIndex] = -1
			}
		}
	}
	return res
}
