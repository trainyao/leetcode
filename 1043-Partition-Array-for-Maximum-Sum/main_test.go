package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_test(t *testing.T) {

	res := maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3)
	assert.Equal(t, 84, res)
}
