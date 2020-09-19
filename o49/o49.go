package o49

import "container/heap"

func nthUglyNumber(n int) int {
	h := &IntHeap{1}
	set := make(map[int]bool)
	count := 0

	for count < n {
		top := (heap.Pop(h)).(int)
		count++
		if count == n {
			return top
		}

		if _, ok := set[top*2]; !ok {
			heap.Push(h, top*2)
			set[top*2] = true
		}
		if _, ok := set[top*3]; !ok {
			heap.Push(h, top*3)
			set[top*3] = true
		}
		if _, ok := set[top*5]; !ok {
			heap.Push(h, top*5)
			set[top*5] = true
		}
	}
	return 0
}

// 以下借由golang官方库的heap interface接口来实现堆  下面堆的实现来自golang官方的测试用例

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// 这里可以控制是大根堆还是小根堆
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
