package t77

func combine(n int, k int) [][]int {
	rs := make([][]int, 0)

	t := make([]int, 0)
	var dfs func(int)
	dfs = func(cur int) {
		if k-len(t) > n-cur+1 {
			return
		}
		if len(t) == k {
			arr := make([]int, k)
			copy(arr, t)
			rs = append(rs, arr)
			return
		}
		t = append(t, cur)
		dfs(cur + 1)
		t = t[:len(t)-1]
		dfs(cur + 1)
	}

	dfs(1)
	return rs
}
