package main

import "log"

func main() {
	head := &ListNode{
		Val: 1,
		//Next: nil,
		Next: &ListNode{
			Val: 2,
			//Next: nil,
			Next: &ListNode{
				Val:  3,
				//Next: nil,
				Next: &ListNode{
					Val:  4,
					Next: nil,
					//Next: &ListNode{
					//	Val:  5,
					//	Next: nil,
					//},
				},
			},
		},
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

		// log
		log[counter] = pcurrent

		// my next
		wantIndex := 0
		if counter%2 == 1 {
			wantIndex = counter + 3
		} else {
			wantIndex = counter - 1
			if wantIndex <= 0 {
				wantIndex = 0
			}
		}
		if p, ok := log[wantIndex]; ok {
			log[counter].Next = p
		}

		// me be wanted
		wantedIndex := 0
		if counter%2 == 1 {
			wantedIndex = counter + 1
			if pnext == nil {
				wantedIndex = counter - 2
				//if wantedIndex <= 0 {
				//	wantedIndex = 0
				//}
			}
		} else {
			wantedIndex = counter - 3
			if wantedIndex <= 0 {
				wantedIndex = 0
			}
		}
		if p, ok := log[wantedIndex]; ok {
			p.Next = pcurrent
		}

		// handle next
		nextIndex := 0
		if counter%2 == 1 {
			nextIndex = counter - 2
			if nextIndex <= 0 {
				nextIndex = 0
			}
			if pnext == nil {
				nextIndex = 99999
			}
		} else {
			nextIndex = counter + 2
			if pnext == nil {
				nextIndex = counter - 1
			}
		}
		if p, ok := log[nextIndex]; ok {
			p.Next = pnext
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
