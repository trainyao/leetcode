package main

import (
	"github.com/trainyao/leetcode/pkg/stack"
	"github.com/trainyao/leetcode/pkg/tree"
)

func PreOrder(bt *tree.BitTreeNode) (res []interface{}) {
	if bt == nil {
		return
	}
	s := stack.Stack{}

	for bt != nil || !s.Empty() {
		// left!=nil时,输出数据
		for bt != nil {
			res = append(res, bt.Value)
			s.Push(bt)
			bt = bt.Left
		}
		bti, _ := s.Pop()
		bt = bti.(*tree.BitTreeNode)
		bt = bt.Right
	}

	return
}
