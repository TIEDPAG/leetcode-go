package t1

func twoSum(nums []int, target int) []int {
	l := len(nums)
	rs := make([]int, 0, 2)
	firstL := l - 1
	for i := 0; i < firstL; i++ {
		v := target - nums[i]
		for j := i + 1; j < l; j++ {
			if nums[j] == v {
				rs = append(rs, i, j)
				return rs
			}
		}
	}
	return rs
}

func twoSum1(nums []int, target int) []int {
	vToI := make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {
		v := target - nums[i]
		if j, ok := vToI[v]; ok {
			return []int{j, i}
		}
		vToI[nums[i]] = i
	}
	return nil
}
