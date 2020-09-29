package t106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(inorder)
	if l == 0 {
		return nil
	}
	node := &TreeNode{
		Val: postorder[l-1],
	}
	i := 0
	for ; i < l; i++ {
		if inorder[i] == postorder[l-1] {
			break
		}
	}
	node.Left = buildTree(inorder[:i], postorder[:i])
	node.Right = buildTree(inorder[i+1:], postorder[i:l-1])
	return node
}
