package o59

type queue struct {
	arr []int
	len int
}

func (q *queue) in(value int) {
	q.arr = append(q.arr, value)
	q.len++
}

func (q *queue) out() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	q.len--
	v := q.arr[0]
	q.arr = q.arr[1 : q.len+1]
	return v, true
}

func (q *queue) outLast() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	q.len--
	v := q.arr[q.len]
	q.arr = q.arr[:q.len]
	return v, true
}

func (q *queue) get() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	return q.arr[0], true
}

func (q *queue) getLast() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	return q.arr[q.len-1], true
}

//func maxSlidingWindow(nums []int, k int) []int {
//	l := len(nums)
//	if l == 0 || k == 0 {
//		return nil
//	}
//	rs := make([]int, 0, l-k+1)
//	queue := queue{
//		arr: make([]int, 0),
//		len: 0,
//	}
//	for i, j := -k+1, 0; j < l; i, j = i+1, j+1 {
//		if i > 0 {
//			v, _ := queue.get()
//			if v == nums[i-1] {
//				queue.out()
//			}
//		}
//		f, ok := queue.getLast()
//		fmt.Printf("%v\t%d\n", queue.arr, i)
//		if !ok {
//			queue.in(nums[j])
//		} else {
//			for f < nums[j] {
//				queue.outLast()
//				if f, ok = queue.getLast(); !ok {
//					break
//				}
//			}
//			queue.in(nums[j])
//		}
//		if i >= 0 {
//			v, _ := queue.get()
//			rs = append(rs, v)
//		}
//	}
//	return rs
//}


func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	deq := make([]int, 0)
	l := len(nums)

	for i := 0; i < l; i++ {
		for deqL := len(deq); deqL != 0 && nums[i] >= nums[deq[deqL-1]]; deqL = len(deq) {
			deq = deq[:deqL-1]
		}
		deq = append(deq, i)

		if deq[0] == i-k {
			deq = deq[1:]
		}

		if i >= k-1 {
			res = append(res, nums[deq[0]])
		}
	}
	return res
}
