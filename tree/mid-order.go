package main

import (
	"github.com/trainyao/leetcode/pkg/stack"
	"github.com/trainyao/leetcode/pkg/tree"
)

func MidOrder(bt *tree.BitTreeNode) (res []interface{}) {
	if bt == nil {
		return
	}
	s := stack.Stack{}

	for bt != nil || !s.Empty() {
		// 找到left == nil 的节点
		for bt != nil {
			s.Push(bt)
			bt = bt.Left
		}

		// 和先序类似, 不过找到left==nil的节点后输出自己
		bti, _ := s.Pop()
		bt = bti.(*tree.BitTreeNode)
		res = append(res, bt.Value)

		bt = bt.Right
	}

	return
}
