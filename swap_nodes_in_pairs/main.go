package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// p -> q -> ...
	// q -> p -> ...
	p := head
	q := p.Next
	p.Next = q.Next
	q.Next = p

	p.Next = swapPairs_1(p.Next)

	head = q
	return head
}
