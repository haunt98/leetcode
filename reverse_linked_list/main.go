package main

import "fmt"

func main() {
	head := generateNodes(1, 2, 3, 4, 5)
	printNodes(head)

	newHead := reverseList_2(head)
	printNodes(newHead)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func debug(nodes ...*ListNode) {
	msg := ""

	for _, node := range nodes {
		if node == nil {
			msg += "nil "
			continue
		}
		msg += fmt.Sprintf("%d", node.Val) + " "
	}

	fmt.Println(msg)
}

func printNodes(head *ListNode) {
	msg := ""

	p := head
	for p != nil {
		msg += fmt.Sprintf("%d", p.Val) + " "
		p = p.Next
	}

	fmt.Println(msg)
}

func printNodesWithName(name string, head *ListNode) {
	fmt.Println(name)
	printNodes(head)
}

func generateNodes(arr ...int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	p := &ListNode{
		Val: arr[0],
	}
	head := p

	for i := 1; i < len(arr); i++ {
		n := &ListNode{
			Val: arr[i],
		}
		p.Next = n
		p = n
	}

	return head
}

// 1 -> 2 -> 3 -> NULL
// 3 -> 2 -> 1 -> NULL
func reverseList_1(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	cur = head

	for cur != nil {
		debug(prev, cur, next)
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// 1 -> 2 -> 3 -> NULL
// 1 -> 2 && 3 -> 2 -> NULL
// 1 -> NULL && 3 -> 2 -> 1 -> NULL
func reverseList_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := head.Next

	head.Next = nil

	newHead := reverseList_2(tail)

	tail.Next = head

	return newHead
}
