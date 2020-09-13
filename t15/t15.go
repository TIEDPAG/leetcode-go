package t15

import "sort"

//func threeSum(nums []int) [][]int {
//	rs := make([][]int, 0)
//	l := len(nums)
//	if l < 3 {
//		return rs
//	}
//	firstL := l - 2
//
//	sort.Ints(nums)
//	// 当同符号则直接返回
//	if nums[0] > 0 || nums[l-1] < 0 {
//		return rs
//	}
//	for i := 0; i < firstL; i++ {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		if nums[i] > 0 {
//			break
//		}
//
//		left, right := i+1, l-1
//		for left < right {
//			sum := nums[i] + nums[left] + nums[right]
//			if sum == 0 {
//				rs = append(rs, []int{nums[i], nums[left], nums[right]})
//				for left < right {
//					if nums[left] != nums[left+1] {
//						left++
//						break
//					}
//					left++
//				}
//				for left < right {
//					if nums[right] != nums[right-1] {
//						right--
//						break
//					}
//					right--
//				}
//			} else if sum < 0 {
//				left++
//			} else {
//				right--
//			}
//		}
//	}
//	return rs
//}

func threeSum(nums []int) [][]int {
	l := len(nums)
	rs := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < l-2; {
		t := -nums[i]
		f, s := i+1, l-1
		for f < s {
			tt := nums[f] + nums[s]
			// 满足条件  记录
			if tt == t {
				rs = append(rs, []int{nums[i], nums[f], nums[s]})
				for f++; f < s && nums[f] == nums[f-1]; f++ {}
				for s--; f < s && nums[s] == nums[s+1]; s-- {}
			}
			// 小了 移动左指针
			if tt < t {
				for f++; f < s && nums[f] == nums[f-1]; f++ {}
			}
			// 大了 移动右指针
			if tt > t {
				for s--; f < s && nums[s] == nums[s+1]; s-- {}
			}
		}

		for i++; i < l-2 && nums[i] == nums[i-1]; i++ {}
	}
	return rs
}


