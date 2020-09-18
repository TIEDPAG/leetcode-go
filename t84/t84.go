package t84

import "fmt"

//func largestRectangleArea(heights []int) int {
//	n := len(heights)
//	left, right := make([]int, n), make([]int, n)
//	mono_stack := make([]int, 0)
//	for i := 0; i < n; i++ {
//		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
//			mono_stack = mono_stack[:len(mono_stack)-1]
//		}
//		if len(mono_stack) == 0 {
//			left[i] = -1
//		} else {
//			left[i] = mono_stack[len(mono_stack)-1]
//		}
//		mono_stack = append(mono_stack, i)
//	}
//	mono_stack = []int{}
//	for i := n - 1; i >= 0; i-- {
//		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
//			mono_stack = mono_stack[:len(mono_stack)-1]
//		}
//		if len(mono_stack) == 0 {
//			right[i] = n
//		} else {
//			right[i] = mono_stack[len(mono_stack)-1]
//		}
//		mono_stack = append(mono_stack, i)
//	}
//	ans := 0
//	for i := 0; i < n; i++ {
//		ans = max(ans, (right[i]-left[i]-1)*heights[i])
//	}
//	return ans
//}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	maxArea := 0

	for i := 0; i < l; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for {
			if len(stack) == 0 || heights[stack[len(stack)-1]] <= heights[i] {
				stack = append(stack, i)
				break
			}
			// 元素出栈 并计算面积
			area := 0
			stack, area = pop(stack, heights, i)

			maxArea = max(maxArea, area)
		}
	}

	for len(stack) > 0 {
		area := 0
		stack, area = pop(stack, heights, l)

		maxArea = max(maxArea, area)
	}
	return maxArea
}

func pop(stack, heights []int, right int) ([]int, int) {
	e := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	left := -1
	if len(stack) != 0 {
		left = stack[len(stack)-1]
	}

	area := (right - left - 1) * heights[e]
	fmt.Println(heights[e], area)
	return stack, area
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
