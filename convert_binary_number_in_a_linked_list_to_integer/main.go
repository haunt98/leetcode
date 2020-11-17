package main

import "fmt"

func main() {
	head := build([]int{1, 0, 0, 1, 0})
	debug(head)
	p := reverse(head)
	debug(p)
	head = build([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})
	fmt.Println(getDecimalValue_1(head))
	head = build([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})
	fmt.Println(getDecimalValue_2(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func build(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val: arr[0],
	}

	p := head

	for i := 1; i < len(arr); i++ {
		q := &ListNode{
			Val: arr[i],
		}
		p.Next = q
		p = q
	}

	return head
}

func debug(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}

func getDecimalValue_1(head *ListNode) int {
	head = reverse(head)

	p := head
	result := 0
	base2 := 1
	for p != nil {
		result += p.Val * base2
		base2 *= 2
		p = p.Next
	}

	return result
}

// p -> q
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	q := head.Next
	p.Next = nil

	for {
		if q == nil {
			break
		}

		temp := q.Next
		q.Next = p
		p = q
		q = temp
	}

	return p
}

func getDecimalValue_2(head *ListNode) int {
	p := head
	result := 0
	for p != nil {
		result = result*2 + p.Val
		p = p.Next
	}
	return result
}
