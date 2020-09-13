package t15

import "testing"

func Test1(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	arr = []int{0, 0, 0, 0}
	t.Log(threeSum(arr))
}
