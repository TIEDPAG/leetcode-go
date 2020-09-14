package t350

func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	rs := make([]int, 0)
	for _, v := range nums1 {
		map1[v]++
	}
	for _, v := range nums2 {
		if num, ok := map1[v]; ok {
			rs = append(rs, v)
			map1[v]--
			if num == 1 {
				delete(map1, v)
			}
		}
	}
	return rs
}
