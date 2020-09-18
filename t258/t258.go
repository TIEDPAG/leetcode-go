package t258

//func addDigits(num int) int {
//	if num < 10 {
//		return num
//	}
//	sum := 0
//	for num > 9 {
//		sum += num % 10
//		num = num / 10
//	}
//	sum += num
//	return addDigits(sum)
//}

func addDigits(num int) int {
	return (num-1)%9 + 1
}
