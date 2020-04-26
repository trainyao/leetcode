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
					//Next: nil,
					Next: &ListNode{
						Val:  5,
						Next: nil,
						//Next: &ListNode{
						//	Val:  6,
						//	Next: nil,
						//},
					},
				},
			},
		},
	}

	head = reverseKGroup(head, 1)

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

func reverseKGroup(head *ListNode, k int) (ret *ListNode) {
	if k == 1 {
		return head
	}

	var counter int
	var ptr *ListNode

	ptr = head
	n := 0
	counter = 1
	start := head
	end := head
	var tmp *ListNode
	for {
		if ptr.Next == nil || counter == k {
			if ptr.Next == nil && counter != k {
				end.Next = start
				break
			}

			// store 4
			tmp = ptr.Next

			ptr.Next = nil

			h, t := revertList(start, counter)
			if n == 0 {
				ret = h
			} else {
				end.Next = h
			}
			end = t

			// if last link
			if tmp == nil {
				break
			}

			// set ptr to 4, begin to count again
			ptr = tmp
			start = tmp
			// reset counter
			counter = 1
			n++
			continue
		}

		counter++
		ptr = ptr.Next
	}

	return
}

func revertList(l *ListNode, len int) (head *ListNode, tail *ListNode) {
	if l.Next == nil {
		return l, l
	}

	tail = l

	a := l
	b := l.Next
	l.Next = nil

	for i := 1; i < len; i++ {
		tmp := b.Next
		b.Next = a
		a = b
		if tmp != nil {
			b = tmp
		}
	}

	head = b
	return
}
