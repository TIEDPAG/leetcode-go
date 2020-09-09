package t416

import "testing"

func Test1(t *testing.T) {
	t.Log(canPartition([]int{1, 5, 11, 5}))
	t.Log(canPartition([]int{1, 2, 3, 4, 5}))
}
