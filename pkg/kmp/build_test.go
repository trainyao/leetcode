package kmp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_build(t *testing.T) {
	a := "abababca"
	table, err := BuildKMP(a)
	assert.NoError(t, err)

	assert.Equal(t, []int{-1, 0, 0, 1, 2, 3, 4, 0}, table.next)
}

func Test_partial_build(t *testing.T) {
	a := "abababca"
	table, err := BuildKMP(a, 5)
	assert.NoError(t, err)
	assert.Equal(t, table.next, []int{-1, 0, 0, 1, 2})

	err = table.ContinueBuild()
	assert.NoError(t, err)

	assert.Equal(t, []int{-1, 0, 0, 1, 2, 3, 4, 0}, table.next)
}
