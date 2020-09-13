package t24

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	f := head.Next
	if f == nil {
		f = head
	}
	var last *ListNode
	for head != nil && head.Next != nil {
		t := head.Next
		head.Next = t.Next
		t.Next = head
		head = head.Next
		if last != nil {
			last.Next = t
		}
		last = t.Next
	}
	return f
}
