package main

import (
	"github.com/trainyao/leetcode/pkg/stack"
	"github.com/trainyao/leetcode/pkg/tree"
)

func PostOrder(bt *tree.BitTreeNode) (res []interface{}) {
	if bt == nil {
		return
	}
	s := stack.Stack{}

	var visited *tree.BitTreeNode
	var tmp *tree.BitTreeNode
	for bt != nil || !s.Empty() {
		for bt != nil {
			s.Push(bt)
			bt = bt.Left
		}

		bti, _ := s.Pop()
		// 取栈顶
		tmp = bti.(*tree.BitTreeNode)

		// 如果栈顶是叶子, 或者左右已经访问过, 输出本节点
		if (tmp.Left == nil && tmp.Right == nil) || (tmp.Right == nil && visited == tmp.Left) || visited == tmp.Right {
			res = append(res, tmp.Value)
			visited = tmp
		} else {
			// 右子树不为空且没访问过过, 先保存本节点, 迭代右子树,走同样的流程
			s.Push(tmp)
			bt = tmp.Right
		}
	}

	return
}
