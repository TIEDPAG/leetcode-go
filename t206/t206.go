package t206

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	t := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return t
}

func reverseList1(head *ListNode) *ListNode {
	var l, r, t *ListNode = nil, head, nil
	for r != nil {
		t = r
		r = r.Next
		t.Next = l
		l = t
	}
	return t
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ps := make([]*ListNode, 0)
	for head != nil {
		ps = append(ps, head)
		head = head.Next
	}
	l := len(ps)
	for i := l - 1; i > 0; i-- {
		ps[i].Next = ps[i-1]
	}
	ps[0].Next = nil
	return ps[l-1]
}
