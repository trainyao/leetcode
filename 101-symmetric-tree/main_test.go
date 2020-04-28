package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	if !assert.Truef(t, isSymmetric(nil), "nil") {
		t.FailNow()
	}

	r := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	if !assert.Truef(t, isSymmetric(r), "one") {
		t.FailNow()
	}

	r = &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right:  &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	if !assert.Truef(t, isSymmetric(r), "many") {
		t.FailNow()
	}

	r = &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right:  &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	if !assert.Falsef(t, isSymmetric(r), "many false") {
		t.FailNow()
	}

	r = &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right:  &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	if !assert.Truef(t, isSymmetric(r), "nil true") {
		t.FailNow()
	}

	r = &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right:  &TreeNode{
			Val: 2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	if !assert.Falsef(t, isSymmetric(r), "nil false") {
		t.FailNow()
	}
}

