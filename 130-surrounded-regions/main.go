package main

import (
	"fmt"
	"github.com/trainyao/leetcode/pkg/djset"
	"github.com/trainyao/leetcode/pkg/queue"
)

//    y -----------------
//  x
//  |   {'X', 'X', 'X', 'X'},
//  |   {'X', 'O', 'O', 'X'},
//  |   {'X', 'X', 'O', 'X'},
//  |   {'X', 'O', 'X', 'X'},

func main() {
	a := [][]byte{
		//{'X', 'X', 'X', 'X'},
		//{'X', 'O', 'O', 'X'},
		//{'X', 'X', 'O', 'X'},
		//{'X', 'O', 'X', 'X'},
		{'O','O','O','O','O','O','O','O','X','O','O','O','O','O','X','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','X','O','O','O','O','O','O','O','O','O','O','O','O'},
		{'X','O','O','X','O','X','O','O','O','O','X','O','O','X','O','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','X','X','O'},
		{'O','X','X','O','O','O','O','O','O','X','O','O','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','X','O','O','O','O','O','O','X','O','O','O','O','O','X','X','O'},
		{'O','O','O','O','O','O','O','X','O','O','O','O','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','O','O','O','O','O','O','X','O','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','X','O'},
		{'O','O','O','O','O','X','O','O','O','O','O','O','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','O','X','O','O','O','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','X','O','O','O','O','X','O','O','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','O','X','O','O','O','O','O','O','O','O','O','O','O'},
		{'X','O','O','O','O','O','O','O','O','X','X','O','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','O','O','O','O','X','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','X','O','O','O','O','O','O','O','O','X','O','O','O','O','O','X'},
		{'O','O','O','O','O','X','O','O','O','O','O','O','O','O','O','X','O','X','O','O'},
		{'O','X','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O'},
		{'O','O','O','O','O','O','O','O','X','X','O','O','O','X','O','O','X','O','O','X'},
		{'O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O','O'},
	}
	solve(a)

	fmt.Printf("%+v", a)
}

type point [2]int

func solve(board [][]byte) {
	xLen := len(board)
	if xLen == 0 {
		return
	}

	yLen := len(board[0])
	if yLen == 0 {
		return
	}

	// init solved matrix
	solved := make([][]bool, xLen)
	for i := 0; i < xLen; i++ {
		solved[i] = make([]bool, yLen)
	}

	set := djset.NewDisJointSet()
	q := queue.Queue{}
	shouldFlip := map[int]bool{}

	for i := 0; i < xLen; i++ {
		for j := 0; j < yLen; j++ {
			//fmt.Println(fmt.Sprintf("%d %d \n", i, j))
			// if current is already checked, skip
			if solved[i][j] {
				continue
			}
			if board[i][j] == 'X' {
				solved[i][j] = true
				continue
			}

			// target := current point, use target to join into same set
			target := i*yLen + j

			// push current to q
			targetPoint := &point{}
			(*targetPoint)[0] = i
			(*targetPoint)[1] = j
			q.Enqueue(targetPoint)

			// index as djset index
			for p := q.Equeue(); p != nil; p = q.Equeue() {
				pp := p.(*point)
				x := (*pp)[0]
				y := (*pp)[1]

				// if left too and left not checked, push to q
				if p := getLeft(x, y); p != nil {
					px := (*p)[0]
					py := (*p)[1]
					if board[px][py] == 'O' && !solved[px][py] {
						q.Enqueue(p)
						solved[px][py] = true
					}
				}

				// if right too and left not checked, push to q
				if p := getRight(x, y, yLen); p != nil {
					px := (*p)[0]
					py := (*p)[1]
					if board[px][py] == 'O' && !solved[px][py] {
						q.Enqueue(p)
						solved[px][py] = true
					}
				}
				// if up too and left not checked, push to q
				if p := getUp(x, y); p != nil {
					px := (*p)[0]
					py := (*p)[1]
					if board[px][py] == 'O' && !solved[px][py] {
						q.Enqueue(p)
						solved[px][py] = true
					}
				}
				// if down too and left not checked, push to q
				if p := getDown(x, y, xLen); p != nil {
					px := (*p)[0]
					py := (*p)[1]
					if board[px][py] == 'O' && !solved[px][py] {
						q.Enqueue(p)
						solved[px][py] = true
					}
				}

				if board[x][y] == 'O' {// } && !solved[x][y] {
					// currentIndex = i * yLen + j
					currentIndex := x*yLen + y
					setId := set.JoinSet(currentIndex, target)

					// check and update set should flip
					if _, ok := shouldFlip[setId]; !ok {
						shouldFlip[setId] = true
					}

					// if set should flip and set has one point on border
					if shouldFlip[setId] && pointOnBorder(x, y, xLen, yLen) {
						shouldFlip[setId] = false
					}
				}

				solved[x][y] = true
			}

			solved[i][j] = true
		}
	}

	for setId, members := range set.Sets() {
		if should, ok := shouldFlip[setId]; ok && should {
			for _, index := range members {
				x := index / yLen
				y := index % yLen
				board[x][y] = 'X'
			}
		}
	}
	return
}

func getLeft(x, y int) *point {
	if y-1 >= 0 {
		return &point{x, y - 1}
	}
	return nil
}

func getRight(x, y, yCount int) *point {
	if y+1 <= yCount-1 {
		return &point{x, y + 1}
	}
	return nil
}

func getUp(x, y int) *point {
	if x-1 >= 0 {
		return &point{x - 1, y}
	}
	return nil
}

func getDown(x, y, xCount int) *point {
	if x+1 <= xCount-1 {
		return &point{x + 1, y}
	}
	return nil
}

func pointOnBorder(x, y, xCount, yCount int) bool {
	return !(x > 0 && y > 0 && x < xCount-1 && y < yCount-1)
}
