package djset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinSet(t *testing.T) {
	set := NewDisJointSet()

	// same index join set
	set.JoinSet(1, 1)
	assert.Equal(t, 1, set.parents[1])
	assert.Equal(t, 1, set.setHeight[1])
	// assert 1 has no setHeight record after join
	set.JoinSet(999, 999)
	set.JoinSet(1, 999)
	_, ok := set.setHeight[1]
	assert.False(t, ok)
	// assert 999's height is 2
	assert.Equal(t, 2, set.setHeight[999])

	// test set height added
	set.JoinSet(888, 888)
	assert.Equal(t, 1, set.setHeight[888])
	set.JoinSet(887, 888)
	assert.Equal(t, 2, set.setHeight[888])

	// different index join set
	set.JoinSet(2, 3)
	assert.Equal(t, 3, set.parents[2])
	assert.Equal(t, 2, set.setHeight[3])
	// assert child has no setHeight record
	_, ok = set.setHeight[2]
	assert.False(t, ok)
	// assert parent has setHeight record
	setHeight, ok := set.setHeight[3]
	assert.True(t, ok)
	assert.Equal(t, 2, setHeight)

	set.JoinSet(4, 3)
	assert.Equal(t, 3, set.parents[4])
	// assert 3's height is still 2
	setHeight, _ = set.setHeight[3]
	assert.Equal(t, 2, setHeight)

	set.JoinSet(5, 4)
	assert.Equal(t, 3, set.parents[5])
	// assert 3's height is still 2
	setHeight, _ = set.setHeight[3]
	assert.Equal(t, 2, setHeight)
}

func TestFindSet(t *testing.T) {
	set := NewDisJointSet()

	_, err := set.FindSet(1)
	assert.Equal(t, SetNotExists, err)

	set.JoinSet(1, 1)
	setId, err := set.FindSet(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, setId)

	set.JoinSet(2, 3)
	set.JoinSet(4, 3)
	set.JoinSet(5, 4)

	setId, err = set.FindSet(2)
	assert.Equal(t, 3, setId)
	setId, err = set.FindSet(4)
	assert.Equal(t, 3, setId)
	setId, err = set.FindSet(5)
	assert.Equal(t, 3, setId)
}

func TestUnion(t *testing.T) {
	set := NewDisJointSet()

	// test two index not exist
	err := set.Union(1, 2)
	assert.Equal(t, SetNotExists, err)

	// test one index not exist
	set.JoinSet(1, 1)
	err = set.Union(1, 2)
	assert.Equal(t, SetNotExists, err)

	// test already in one set
	err = set.Union(1, 1)
	assert.Equal(t, AlreadyInOneSet, err)

	// put 1, 2 together
	// current:
	// 2
	// |
	// 1
	set.JoinSet(1, 2)

	// test height is the same
	// current:      expected after:
	// 2  4               4
	// |  |            /  |
	// 1  3            2  3
	//                 |
	//                 1
	set.JoinSet(3, 4)
	err = set.Union(1, 3)
	assert.NoError(t, err)
	// 1's ste should be 4
	setIdOf1, err := set.FindSet(1)
	assert.NoError(t, err)
	assert.Equal(t, 4, setIdOf1)
	assert.Equal(t, 3, set.setHeight[4])

	// test index1 higher than 2
	// current:      expected after:
	//     4               4
	//  /  |            /  | \
	//  2  3            2  3  5
	//  |               |
	//  1               1
	set.JoinSet(5, 5)
	err = set.Union(1, 5)
	assert.NoError(t, err)
	assert.Equal(t, 3, set.setHeight[4])
	_, ok := set.setHeight[5]
	assert.False(t, ok)

	// test index2 higher than 1
	// current:      expected after:
	//     4               4
	//  /  | \          /  | \ \
	//  2  3  5         2  3  5 6
	//  |               |
	//  1               1
	set.JoinSet(6, 6)
	err = set.Union(6, 1)
	assert.NoError(t, err)
	assert.Equal(t, 3, set.setHeight[4])
	_, ok = set.setHeight[6]
	assert.False(t, ok)

	// check data
	assert.Equal(t, 4, set.parents[2])
	assert.Equal(t, 4, set.parents[3])
	assert.Equal(t, 4, set.parents[5])
	assert.Equal(t, 4, set.parents[6])
	assert.Equal(t, 2, set.parents[1])
}
