package t47

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	rs := make([][]int, 0)
	l := len(nums)
	temp := make([]int, 0, l)

	flag := make(map[int]bool)

	var backtrack func(int)
	backtrack = func(cur int) {
		if cur == l {
			_rs := make([]int, l)
			copy(_rs, temp)
			rs = append(rs, _rs)
			return
		}

		for i, v := range nums {
			if flag[i] || (i > 0 && !flag[i-1] && v == nums[i-1]) {
				continue
			}

			temp = append(temp, v)
			flag[i] = true
			backtrack(cur + 1)
			flag[i] = false
			temp = temp[:len(temp)-1]
		}
	}

	backtrack(0)
	return rs
}
