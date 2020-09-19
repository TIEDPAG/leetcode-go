package o40

import "container/heap"

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	l := len(arr)
	if l == k {
		return arr
	}
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < l; i++ {
		if h.Len() < k {
			heap.Push(h, arr[i])
			continue
		}
		if arr[i] < (*h)[0] {
			heap.Push(h, arr[i])
			heap.Pop(h)
		}
	}

	return *h
}

// 以下借由golang官方库的heap interface接口来实现堆  下面堆的实现来自golang官方的测试用例

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// 这里可以控制是大根堆还是小根堆
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
