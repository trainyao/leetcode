package heap

import (
	"errors"
	"math"
	"sync"
)

var InvalidIndex = errors.New("invalid index")

type MinimumBinaryHeap struct {
	score []int
	data  []interface{}
	mutex sync.Mutex
}

// Add add data to heap and adjust score to fit minimum binary heap
// and return the new data node's index
func (h *MinimumBinaryHeap) Add(data interface{}, score int) (index int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// append data to data aaray
	// adjust index
	h.data = append(h.data, data)
	h.score = append(h.score, score)

	//index = len(h.data) - 1
	//parentIndex, err := parent(index)
	//// when err != nil, calculating root's parent, stop the loop
	//for err == nil && h.score[index] < h.score[parentIndex] {
	//	// swap data and index
	//	h.data[index], h.data[parentIndex] = h.data[parentIndex], h.data[index]
	//	h.score[index], h.score[parentIndex] = h.score[parentIndex], h.score[index]
	//
	//	// parent being index
	//	index = parentIndex
	//	// find next parent index
	//	parentIndex, err = parent(index)
	//}
	return h.goUp(len(h.data) - 1)
}

// ExtractMin remote the root node, and adjust the tree nodes
func (h *MinimumBinaryHeap) ExtractMin() (data interface{}, score int) {
	// move deleting index to tail of array
	data = h.data[0]
	score = h.score[0]
	// then move it
	index := 0
	totalCount := len(h.data)
	// operation below is comparing which is smaller, so set current score to super big
	h.score[index] = math.MaxInt64

	// if current index's left&right child is not exist
	// reach leaf node, stop the loop
	for leftChild(index) > totalCount-1 && rightChild(index) > totalCount-1 {
		chosenIndex := -1
		// for each node to adjust, choose smaller child to be current node
		if lci := leftChild(index); lci <= totalCount-1 && h.score[lci] <= h.score[index] {
			h.score[index] = h.score[lci]
			chosenIndex = lci
		}
		if rci := rightChild(index); rci <= totalCount-1 && h.score[rci] <= h.score[index] {
			h.score[index] = h.score[rci]
			chosenIndex = rci
		}
		// after that, set that node to be current node
		if chosenIndex != -1 {
			h.data[index] = h.data[chosenIndex]
			index = chosenIndex
		}
	}
	// index is leaf node index currently
	// for it has been move to parent, data and score should be deleted by update slice len
	//
	// if current is left child and has sibling right child node,
	// move it to left child
	if isLeftChild(index) && index+1 <= totalCount-1 {
		h.data[index] = h.data[index+1]
		h.score[index] = h.score[index+1]
	}

	h.data = h.data[:totalCount-1]
	h.score = h.score[:totalCount-1]

	return
}

func (h *MinimumBinaryHeap) UpdateScore(index int, score int) (indexAfter int) {
	total := len(h.data)
	if index < 0 || index >= total-1 {
		return -1
	}

	// edit score
	h.score[index] = score

	// do adjust logic
	/*
		         -2
		     1         2
		   3   4     -99      6
		7   8 9 10  11 12  13 14
	*/
	if parentIndex, err := parent(index); err == nil && h.score[parentIndex] >= h.score[index] {
		return h.goUp(index)
	}

	return h.goDown(index)

	//for leftChild(index) > total-1 && rightChild(index) > total-1 {
	//	chosenIndex := -1
	//	if lci := leftChild(index); lci <= total-1 && h.score[lci] <= h.score[index] {
	//		h.score[lci], h.score[index] = h.score[index], h.score[lci]
	//		h.data[lci], h.data[index] = h.data[index], h.data[lci]
	//		chosenIndex = lci
	//	}
	//	if rci := leftChild(index); rci <= total-1 && h.score[rci] <= h.score[index] {
	//		h.score[rci], h.score[index] = h.score[index], h.score[rci]
	//		h.data[rci], h.data[index] = h.data[index], h.data[rci]
	//		chosenIndex = rci
	//	}
	//	if chosenIndex != -1 {
	//		index = chosenIndex
	//	}
	//}
	//return index
}

func (h *MinimumBinaryHeap) goDown(index int) (indexAfter int) {
	total := len(h.data)
	for leftChild(index) > total-1 && rightChild(index) > total-1 {
		chosenIndex := -1
		if lci := leftChild(index); lci <= total-1 && h.score[lci] <= h.score[index] {
			h.score[lci], h.score[index] = h.score[index], h.score[lci]
			h.data[lci], h.data[index] = h.data[index], h.data[lci]
			chosenIndex = lci
		}
		if rci := leftChild(index); rci <= total-1 && h.score[rci] <= h.score[index] {
			h.score[rci], h.score[index] = h.score[index], h.score[rci]
			h.data[rci], h.data[index] = h.data[index], h.data[rci]
			chosenIndex = rci
		}
		if chosenIndex != -1 {
			index = chosenIndex
		}
	}
	return index
}

func (h *MinimumBinaryHeap) goUp(index int) (indexAfter int) {
	parentIndex, err := parent(index)
	// when err != nil, calculating root's parent, stop the loop
	for err == nil && h.score[index] < h.score[parentIndex] {
		// swap data and index
		h.data[index], h.data[parentIndex] = h.data[parentIndex], h.data[index]
		h.score[index], h.score[parentIndex] = h.score[parentIndex], h.score[index]

		// parent being index
		index = parentIndex
		// find next parent index
		parentIndex, err = parent(index)
	}
	return index
}

func parent(index int) (parentIndex int, err error) {
	if index <= 0 {
		return 0, InvalidIndex
	}
	return (index - 1) / 2, nil
}

func leftChild(index int) (leftChildIndex int) {
	return 2*index + 1
}

func rightChild(index int) (rightChildIndex int) {
	return 2*index + 2
}
func isLeftChild(index int) bool {
	return index%2 == 1
}

func isRightChild(index int) bool {
	return index%2 == 0
}
