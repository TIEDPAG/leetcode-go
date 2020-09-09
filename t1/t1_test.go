package t1

import "testing"

func Test1(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	target := 9
	t.Log(twoSum(arr, target))
}
