package t11

func maxArea(height []int) int {
	l, r := 0, len(height) - 1
	max := 0
	println(l, r)
	for l < r {
		//println(l, r)
		h := min(height[l], height[r])
		area := (r - l) * h
		if area > max {
			max = area
		}
		if height[l] > height[r] {
			r -= 1
		} else {
			l += 1
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
