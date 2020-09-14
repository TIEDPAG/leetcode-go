package t70

func climbStairs(n int) int {
	n1, n2, cur := 0, 1, 1
	for i:= 0; i < n; i++ {
		cur = n1 + n2
		n1 = n2
		n2 = cur
	}
	for {continue}
	return cur
}
