package t66

func plusOne(digits []int) []int {
	l := len(digits)
	last := digits[l -1]
	if last < 9{
		digits[l - 1] = last + 1
		return digits
	}
	// 如果大于9 开始循环
	flag := true
	for i := l; i > 0; i-- {
		if !flag{
			break
		}
		last = digits[i - 1] + 1
		flag = false
		digits[i - 1] = last
		if last > 9 {
			digits[i - 1] = last % 10
			flag = true
		}
	}
	// 结束循环时不需要进位
	if !flag {
		return digits
	}

	rs := make([]int, 0, l+1)
	rs = append(rs, 1)
	rs = append(rs, digits[:]...)
	return rs
}
