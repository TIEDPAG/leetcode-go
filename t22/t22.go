package t22

func generateParenthesis(n int) []string {
	return generate([]string{}, []byte{}, 0, 0, n)
}

func generate(rs []string, chs []byte, left, right, n int) []string {
	if len(chs) == 2*n {
		return append(rs, string(chs))
	}
	// 左括号不足
	if left < n {
		chs = append(chs, '(')
		rs = generate(rs, chs, left+1, right, n)
		// 递归完成 删除左括号
		chs = chs[:len(chs)-1]
	}
	// 左括号比右括号多 尝试加左括号  有点像爬楼梯
	if right < left {
		chs = append(chs, ')')
		rs = generate(rs, chs, left, right+1, n)
		chs = chs[:len(chs)-1]
	}
	return rs
}

var cache = make(map[int][]string, 10)

func generateParenthesis1(n int) []string {
	return generate1(n, []string{})
}

func generate1(n int, rs []string) []string {
	if len(cache[n]) != 0 {
		return cache[n]
	}

	if n == 0 {
		rs = append(rs, "")
	} else {
		for i := 0; i < n; i++ {
			for _, left := range generate1(i, rs) {
				for _, right := range generate1(n-i-1, rs) {
					rs = append(rs, left+"("+right+")")
				}
			}
		}
	}
	cache[n] = rs
	return rs
}
