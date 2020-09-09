package t283

func moveZeroes(nums []int) {
	l := len(nums)
	zeroI := -1
	for i := 0; i < l; {
		if nums[i] == 0 && zeroI < 0 {
			zeroI = i // 记录索引
			i++
			continue
		}
		if nums[i] != 0 && zeroI >= 0 {
			nums[zeroI], nums[i] = nums[i], nums[zeroI]
			i, zeroI = zeroI, -1
		}
		i++
	}
}

//func moveZeroes(nums []int) {
//	l := len(nums)
//	zeroI := 0
//	for i := 0; i < l; i++ {
//		if nums[i] != 0 {
//			if i != zeroI {
//				nums[zeroI] = nums[i]
//			}
//			zeroI++
//		}
//	}
//
//	for ; zeroI < l; zeroI++ {
//		nums[zeroI] = 0
//	}
//}
