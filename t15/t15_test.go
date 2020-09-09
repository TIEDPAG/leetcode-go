package t15

import "testing"

func Test1(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	arr = []int{-2,0,0,2,2}
	t.Log(threeSum(arr))
}
