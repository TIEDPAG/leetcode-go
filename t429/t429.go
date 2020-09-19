package t429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	rs := make([][]int, 0)
	queue := []*Node{root}
	for len(queue) != 0 {
		l := len(queue)
		curLevel := make([]int, l)
		for i := 0; i < l; i++ {
			curLevel[i] = queue[0].Val
			queue = append(queue, queue[0].Children...)
			queue = queue[1:]
		}
		rs = append(rs, curLevel)
	}
	return rs
}
