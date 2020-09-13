package t206

import "testing"

func Test1(t *testing.T) {
	h := &ListNode{
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
	}

	rs := reverseList2(h)
	for rs != nil {
		t.Log(rs.Val)
		rs = rs.Next
	}
}
