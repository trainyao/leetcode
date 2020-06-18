package main

import (
	"fmt"
	"reflect"

	"github.com/trainyao/leetcode/pkg/tree"
)

func main() {
	t := &tree.BitTreeNode{
		Left: &tree.BitTreeNode{
			Left: &tree.BitTreeNode{
				Left:  nil,
				Right: nil,
				Value: 4,
			},
			Right: &tree.BitTreeNode{
				Left:  nil,
				Right: nil,
				Value: 5,
			},
			Value: 2,
		},
		Right: &tree.BitTreeNode{
			Left: &tree.BitTreeNode{
				Left:  nil,
				Right: nil,
				Value: 6,
			},
			Right: &tree.BitTreeNode{
				Left:  nil,
				Right: nil,
				Value: 7,
			},
			Value: 3,
		},
		Value: 1,
	}

	res := []int{}
	for _, i := range PreOrder(t) {
		res = append(res, i.(int))
	}
	fmt.Printf("%+v", res)
	print(reflect.DeepEqual(res, []int{1,2,4,5,3,6,7}))

	res = []int{}
	for _, i := range MidOrder(t) {
		res = append(res, i.(int))
	}
	fmt.Printf("%+v", res)
	print(reflect.DeepEqual(res, []int{4,2,5,1,6,3,7}))

	res = []int{}
	for _, i := range PostOrder(t) {
		res = append(res, i.(int))
	}
	fmt.Printf("%+v", res)
	print(reflect.DeepEqual(res, []int{4,5,2,6,7,3,1}))
}
