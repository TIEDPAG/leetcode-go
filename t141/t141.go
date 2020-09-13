package t141

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	kv := make(map[*ListNode]bool)
	for head != nil {
		if v := kv[head]; v {
			return true
		}
		kv[head] = true
		head = head.Next
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	t := head
	i := 0
	for t != nil {
		tt := head
		for j := 0; j < i; j++ {
			if tt == t {
				return true
			}
			tt = tt.Next
		}
		t = t.Next
		i++
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	t1, t2 := head, head
	for t2 != nil && t2.Next != nil {
		t1 = t1.Next
		t2 = t2.Next.Next
		if t1 == t2 {
			return true
		}
	}
	return false
}
