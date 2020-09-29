package t617

import "testing"

func Test1(t *testing.T) {
	t1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	
	t2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	mergeTrees(t1, t2)
}
