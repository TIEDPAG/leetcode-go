package t297

import "testing"

func Test1(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}
	ser := Constructor()
	//deser := Constructor()
	data := ser.serialize(root)
	t.Log(data)
	ans := ser.deserialize(data)
	t.Log(ser.serialize(ans) == data)
}
