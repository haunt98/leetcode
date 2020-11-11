package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists_1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// l1 and l2 is not nil
	if l1.Val <= l2.Val {
		nextL1 := l1.Next
		l1.Next = mergeTwoLists_1(nextL1, l2)
		return l1
	}

	nextL2 := l2.Next
	l2.Next = mergeTwoLists_1(l1, nextL2)
	return l2
}

func mergeTwoLists_2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head, cur, p, q *ListNode

	// l1 and l2 is not nil
	if l1.Val <= l2.Val {
		head = l1
		cur = l1
		p = l1.Next
		q = l2
	} else {
		head = l2
		cur = l2
		p = l1
		q = l2.Next
	}

	for {
		if p == nil && q == nil {
			cur.Next = nil
			break
		}

		if p == nil {
			cur.Next = q
			break
		}

		if q == nil {
			cur.Next = p
			break
		}

		// p and q is not nil
		if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
			cur = cur.Next
			continue
		}

		cur.Next = q
		q = q.Next
		cur = cur.Next
	}

	return head
}
