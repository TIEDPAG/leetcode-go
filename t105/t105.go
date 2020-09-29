package t105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	l := len(preorder)
	if l == 0 {
		return nil
	}
	node := &TreeNode{
		Val: preorder[0],
	}
	i := 0
	for ; i < l; i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return node
}
