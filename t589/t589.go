package t589

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	return p(root, []int{})
}

func p(node *Node, arr []int) []int {
	if node == nil {
		return arr
	}
	arr = append(arr, node.Val)
	for _, c := range node.Children {
		arr = p(c, arr)
	}
	return arr
}
