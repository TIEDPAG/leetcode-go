package t46

func permute(nums []int) [][]int {
	rs := make([][]int, 0)
	l := len(nums)
	temp := make([]int, l)
	copy(temp, nums)
	var backtrack func(int)
	backtrack = func(cur int) {
		if cur == l {
			_rs := make([]int, l)
			copy(_rs, temp)
			rs = append(rs, _rs)
			return
		}

		for i := cur; i < l; i++ {
			temp[cur], temp[i] = temp[i], temp[cur]
			backtrack(cur+1)
			temp[cur], temp[i] = temp[i], temp[cur]
		}
	}

	backtrack(0)
	return rs
}
