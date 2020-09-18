package t25

import "testing"

func Test1(t *testing.T) {
	h := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	h = reverseKGroup(h, 4)
	//h, _ = reverse(h, nil, 3)
	for h != nil {
		t.Log(h.Val)
		h = h.Next
	}
}
