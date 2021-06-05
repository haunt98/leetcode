package main

import "fmt"

func main() {
	fmt.Println("2. Add Two Numbers")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var p *ListNode

	memory := 0
	for {
		if l1 == nil && l2 == nil && memory == 0 {
			break
		}
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		val := val1 + val2 + memory
		memory = val / 10
		temp := &ListNode{
			Val: val % 10,
		}

		if p == nil {
			p = temp
			head = p
		} else {
			p.Next = temp
			p = p.Next
		}
	}

	return head
}
