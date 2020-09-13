package t25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var prev *ListNode
	i := 0
	first := head
	for cur != nil {
		i++
		if i%k == 0 {

			continue
		}
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return head.Next
}
