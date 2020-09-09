package t283

import "testing"

func Test1(t *testing.T) {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	t.Log(arr)
}
