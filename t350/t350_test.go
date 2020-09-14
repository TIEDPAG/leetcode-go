package t350

import "testing"

func Test1(t *testing.T) {
	arr1, arr2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4}

	t.Log(intersect(arr1, arr2))
}
