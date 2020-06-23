package heap

import (
	"errors"
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

	index = len(h.data) - 1
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
	return
}

func (h *MinimumBinaryHeap) Remove(index int) {
	// move deleting index to tail of array
	// then move it
}

func (h *MinimumBinaryHeap) SetScore(index int, score int) {
	// edit score
	// do adjust logic

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
