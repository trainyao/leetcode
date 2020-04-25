package main

func main() {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: nil,
	}
	t2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: nil,
	}
	print(isSameTree(t1, t2))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	s1 := getS(p)
	s2 := getS(q)
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	for i, c := range s1 {
		if s2[i] != c {
			return false
		}
	}
	return true
}

func getS(t *TreeNode) (r []int) {
	if t == nil {
		return r
	}

	stack := stack{}
	var interf interface{}
	stack.push(t)

	e := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	for !stack.empty() {
		interf, _ = stack.pop()
		t, _ = interf.(*TreeNode)
		r = append(r, t.Val)

		if t.Right != nil {
			if t.Left == nil {
				t.Left = e
			}
			stack.push(t.Right)
		}
		if t.Left != nil {
			if t.Right == nil {
				stack.push(e)
			}
			stack.push(t.Left)
		}
	}

	return
}
