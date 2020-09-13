package t66

func plusOne(digits []int) []int {
	l := len(digits)
	flag := false
	digits[l-1] += 1
	for i := l - 1; i >= 0; i-- {
		v := digits[i]
		if flag {
			v++
		}
		digits[i] = v % 10
		flag = v >= 10
	}
	if flag {
		rs := make([]int, 0, l + 1)
		rs = append(rs, 1)
		rs = append(rs, digits...)
		return rs
	}
	return digits
}
