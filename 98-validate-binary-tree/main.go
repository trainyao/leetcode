package main

import (
	"github.com/trainyao/leetcode/pkg/stack"
)

func main() {
	//print(isSymmetric())
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	inorder := 0
	init := true
	s := stack.Stack{}

	for !s.Empty() || root != nil {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		interf, _ := s.Pop()
		root = interf.(*TreeNode)

		if init {
			init = false
		} else {
			if root.Val <= inorder {
				return false
			}
		}
		inorder = root.Val

		root = root.Right
	}
	return true
}