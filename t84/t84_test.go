package t84

import "testing"

func Test1(t *testing.T) {
	t.Log(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	t.Log(largestRectangleArea([]int{1, 1}))
	t.Log(largestRectangleArea([]int{2, 1, 2}))
}
