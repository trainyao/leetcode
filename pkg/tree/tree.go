package tree

type BitTreeNode struct {
	Left  *BitTreeNode
	Right *BitTreeNode
	Value interface{}
}

type Tree struct {
	root *Node
}

type Node struct {
	Children []*Node
	Value    interface{}
}
