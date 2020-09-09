package t11

import "testing"

func Test1(t *testing.T) {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(arr))
}
