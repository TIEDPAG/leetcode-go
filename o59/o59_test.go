package o59

import "testing"

func Test1(t *testing.T) {
	arr, k := []int{1, 3, -1, -3, 5, 3, 6, 7}, 3
	t.Log(maxSlidingWindow(arr, k))
}

func Test2(t *testing.T) {
	arr, k := []int{1, 3, 1, 2, 0, 5}, 3
	t.Log(maxSlidingWindow(arr, k))
}

func Test3(t *testing.T) {
	arr, k := []int{7, 2, 4}, 2
	t.Log(maxSlidingWindow(arr, k))
}
