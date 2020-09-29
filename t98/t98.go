package t98

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	arr := p(root, []int{})
	l := len(arr)
	for i := 1; i < l; i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func p(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}

	arr = p(root.Left, arr)
	arr = append(arr, root.Val)
	arr = p(root.Right, arr)
	return arr
}

func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return valid(root.Left, root.Val, math.MinInt64) && valid(root.Right, math.MaxInt64, root.Val)
}

func valid(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	if !valid(root.Left, root.Val, min) {
		return false
	}
	if !valid(root.Right, max, root.Val) {
		return false
	}
	return true
}
