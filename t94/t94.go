package t94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inorder(root, []int{})
}

func inorder(node *TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}
	arr = inorder(node.Left, arr)
	arr = append(arr, node.Val)
	arr = inorder(node.Right, arr)
	return arr
}
