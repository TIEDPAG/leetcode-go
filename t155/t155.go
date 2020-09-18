package t155

type MinStack struct {
	stack []int
	minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		minStack: make([]int, 0),
	}
}


func (s *MinStack) Push(x int)  {
	s.stack = append(s.stack, x)
	if len(s.minStack) == 0 {
		s.minStack = append(s.minStack, x)
		return
	}

	// 判断新元素是否为最小
	if x < s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, x)
	} else {
		s.minStack = append(s.minStack, s.minStack[len(s.minStack)-1])
	}
}


func (s *MinStack) Pop()  {
	// if len(this.minStack) == 0 {
	//     return
	// }
	s.minStack = s.minStack[:len(s.minStack)-1]
	s.stack = s.stack[:len(s.stack)-1]
}


func (s *MinStack) Top() int {
	// if len(this.stack) == 0 {
	//     return 0
	// }
	return s.stack[len(s.stack)-1]
}


func (s *MinStack) GetMin() int {
	// if len(this.stack) == 0 {
	//     return 0
	// }
	return s.minStack[len(s.minStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */