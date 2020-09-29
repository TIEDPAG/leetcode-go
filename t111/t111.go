package t111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := 0
	if root.Left != nil {
		min = minDepth(root.Left) + 1
	}
	if root.Right != nil {
		r := minDepth(root.Right) + 1
		if min == 0 {
			min = r
		}
		if min > r {
			min = r
		}
	}
	return min
}
