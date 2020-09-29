package t98

import "testing"

func Test1(t *testing.T) {
	t.Log(isValidBST1(&TreeNode{
		Val:   -2147483648,
		Left:  nil,
		Right: &TreeNode{
			Val:   2147483647,
			Left:  nil,
			Right: nil,
		},
	}))
}
