package t141

import "testing"

func Test1(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	t.Log(hasCycle2(h))
}

func Test2(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	h.Next.Next.Next.Next = h.Next.Next

	t.Log(hasCycle2(h))
}

func Test3(t *testing.T) {
	var h *ListNode = nil

	t.Log(hasCycle2(h))
}
