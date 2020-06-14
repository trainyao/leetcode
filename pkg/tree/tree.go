package tree

import (
	"fmt"
	stack2 "github.com/trainyao/leetcode/pkg/stack"
)

type Tree struct {
	parents []int
	root    *Node
}

type Node struct {
	Child []*Node
	Value interface{}
	Tree  *Tree
}

func NewTree(root Node) (*Tree, error) {
	parents := []int{}

	tree := &Tree{
		root: &root,
	}

	var node *Node
	node = &root
	stack := stack2.Stack{}

	var ok bool
	for node != nil {
		if len(node.Child) != 0 {
			for _, child := range node.Child {
				stack.Push(child)
			}

			interf, err := stack.Pop()
			if err != nil {
				node = nil
			}
			if node, ok = interf.(*Node); !ok {
				return nil, fmt.Errorf("parsing node pointer failed")
			}

			continue
		}
	}

	tree.parents = parents

	return tree
}
