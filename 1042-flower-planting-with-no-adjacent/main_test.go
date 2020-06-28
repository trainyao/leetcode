package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_test(t *testing.T) {

	res := gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}})
	//Output: [1,2,3]
	assert.Equal(t, []int{1, 2, 3}, res)

	res = gardenNoAdj(4, [][]int{{1, 2}, {3, 4}})
	//Output: [1,2,1,2]
	assert.Equal(t, []int{1, 2, 1, 2}, res)

	res = gardenNoAdj(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}})
	//Output: [1,2,3,4]
	assert.Equal(t, []int{1, 2, 3, 4}, res)

}
