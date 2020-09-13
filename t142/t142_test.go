package t142

import "testing"

func Test1(t *testing.T) {
	h := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	h.Next.Next.Next.Next = h.Next

	t.Log(detectCycle(h).Val)
}
