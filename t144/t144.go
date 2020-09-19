package t144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return preorder(root, []int{})
}

func preorder(node *TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}
	arr = append(arr, node.Val)
	arr = preorder(node.Left, arr)
	arr = preorder(node.Right, arr)
	return arr
}
