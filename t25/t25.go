package t25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	h := &ListNode{
		Next: head,
	}
	cur := h
	prev := cur
	next := cur
loopLink:
	for cur.Next != nil {
		prev = cur
		first := cur.Next
		for i := 0; i < k; i++ {
			if cur.Next == nil {
				break loopLink
			}
			cur = cur.Next
		}
		// 切断链表 保存上下指针
		next = cur.Next
		cur.Next = nil
		prev.Next = reverse(first)
		first.Next = next
		cur = first
	}
	return h.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		t := head.Next
		head.Next = prev
		prev = head
		head = t
	}
	return prev
}
