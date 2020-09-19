package t590

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	return p(root, []int{})
}

func p(node *Node, arr []int) []int {
	if node == nil {
		return arr
	}
	for _, c := range node.Children {
		arr = p(c, arr)
	}
	arr = append(arr, node.Val)
	return arr
}
