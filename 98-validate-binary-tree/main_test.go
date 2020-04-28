package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	if !assert.Truef(t, isValidBST(nil), "nil") {
		t.FailNow()
	}

	r := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	if !assert.Truef(t, isValidBST(r), "one") {
		t.FailNow()
	}

	r = &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	if !assert.Truef(t, isValidBST(r), "many") {
		t.FailNow()
	}

	r = &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right:  &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	if !assert.Falsef(t, isValidBST(r), "false") {
		t.FailNow()
	}

	r = &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right:  &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}
	if !assert.Falsef(t, isValidBST(r), "many false") {
		t.FailNow()
	}
}

