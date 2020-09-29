package t113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	rs := make([][]int, 0)
	var findPath func(*TreeNode, int, []int)
	findPath = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		sum -= node.Val
		// 叶子节点
		if node.Left == nil && node.Right == nil {
			if sum == 0 {
				rs = append(rs, path)
			}
			return
		}
		if node.Left != nil {
			lPath := make([]int, len(path), len(path)+1)
			copy(lPath, path)
			findPath(node.Left, sum, lPath)
		}
		if node.Right != nil {
			rPath := make([]int, len(path), len(path)+1)
			copy(rPath, path)
			findPath(node.Right, sum, rPath)
		}
	}
	findPath(root, sum, []int{})
	return rs
}
