package t21

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sort(l1, l2 *ListNode) (*ListNode, *ListNode) {
	if l1.Val > l2.Val {
		return l2, l1
	}
	return l1, l2
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	//var head *ListNode
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	head := &ListNode{
//		Val:  math.MinInt32,
//		Next: nil,
//	}
//	prev := head
//	for l1 != nil && l2 != nil {
//		l1, l2 = sort(l1, l2)
//		l1, l1.Next, prev, prev.Next = l1.Next, l2, l1, l1
//	}
//	return head.Next
//}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	l1, l2 = sort(l1, l2)
//	l1.Next = mergeTwoLists(l1.Next, l2)
//	return l1
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	first := &ListNode{}
	prev := first
	for l1 != nil && l2 != nil {
		l1, l2 = sort(l1, l2)
		//fmt.Printf("%d\t%d\n", l1.Val, l2.Val)
		prev.Next = l1
		prev = l1
		l1 = l1.Next
	}
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return first.Next
}
