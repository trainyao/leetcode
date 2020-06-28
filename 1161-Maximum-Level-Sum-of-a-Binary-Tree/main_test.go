package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_test(t *testing.T) {

	res := maxLevelSum(&TreeNode{})
	//Output: [1,2,3]
	assert.Equal(t, []int{1, 2, 3}, res)

	res = maxLevelSum(&TreeNode{})
	//Output: [1,2,1,2]
	assert.Equal(t, []int{1, 2, 1, 2}, res)

	res = maxLevelSum(&TreeNode{})
	//Output: [1,2,3,4]
	assert.Equal(t, []int{1, 2, 3, 4}, res)

}
