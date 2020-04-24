package main

import "log"

func main() {
	head := &ListNode{
		Val: 1,
		Next: nil,
		//Next: &ListNode{
		//	Val: 2,
		//	Next: nil,
		//	//Next: &ListNode{
		//	//	Val:  3,
		//	//	Next: nil,
		//	//	//Next: &ListNode{
		//	//	//	Val:  4,
		//	//	//	Next: nil,
		//	//	//	//Next: &ListNode{
		//	//	//	//	Val:  5,
		//	//	//	//	Next: nil,
		//	//	//	//},
		//	//	//},
		//	//},
		//},
	}

	head = swapPairs(head)
	for true {
		log.Println("%d", head.Val)

		if head.Next == nil {
			break
		}

		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
1.next = 4
2.next = 1
0.next = 2

2.next = 1
0.next = 2
4.next = 3
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	pcurrent := head
	log := map[int]*ListNode{}
	// special node
	log[0] = &ListNode{}

	counter := 1
	for true {
		pnext := pcurrent.Next

		log[counter]=pcurrent

		if counter % 2 == 1{
			index := counter - 2
			if index <= 0 {
				index = 0
			}

			log[index].Next = pcurrent.Next

			if pnext == nil {
				log[index].Next = pcurrent
			}
		} else {
			index := counter - 1
			log[counter].Next = log[index]

			if pnext == nil {
				log[index].Next = nil
			}
		}

		if pnext == nil {
			break
		} else {
			pcurrent = pnext
			counter++
		}
	}

	return log[0].Next
}
