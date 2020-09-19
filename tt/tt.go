package tt

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

func maxArea(height []int) int {
	area, l := 0, len(height)-1
	i, j := 0, l
	for i < j {
		h := min(height[i], height[j])
		area = max(area, h*(j-i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return area
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

type Event struct {
	//ID        string            `josn:"_id"`
	EventName string                 `json:"eventName"`
	Timestamp string                 `json:"timestamp"`
	Time      int64                  `json:"time"`
	EventData map[string]interface{} `json:"eventData"`
	Info      struct {
		UserID string `json:"userId"`
		RoomId string `json:"roomId"`
		Role   string `json:"role"`
	} `json:"info"`
	IsSuccess bool `json:"isSuccess"`
}

func Sort(jsons ...string) string {
	objs := make([]Event, 0)
	for _, s := range jsons {
		arr := make([]Event, 0)
		if err := json.Unmarshal([]byte(s), &arr); err != nil {
			fmt.Printf("%v\n", err)
		}
		objs = append(objs, arr...)
	}
	for i := range objs {
		t, err := strconv.ParseInt(objs[i].Timestamp, 10, 64)
		fmt.Printf("time: %d\n", t)
		objs[i].Time = t
		if err != nil {
			fmt.Printf("%v\n", err)
			//return ""
		}
	}
	sort.Slice(objs, func(i, j int) bool {
		return objs[i].Time < objs[j].Time
	})
	rs, _ := json.Marshal(objs)
	return string(rs)
}

func Json(str string) Event {
	var rs Event
	if err := json.Unmarshal([]byte(str), &rs); err != nil {
		fmt.Printf("%v\n", err)
	}
	return rs
}

func LargestRectangleArea(heights []int) int {
	l := len(heights)
	if len(heights) == 0 {
		return 0
	}
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	l += 2
	stack := make([]int, 0)
	maxArea := 0

	for i := 0; i < l; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for {
			if len(stack) == 0 || heights[i] <= heights[stack[len(stack)-1]] {
				stack = append(stack, i)
				break
			}

			// 出栈  并计算面积
			area := 0
			stack, area = pop(stack, heights, i)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func pop(stack, heights []int, right int) ([]int, int) {
	cur := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	if heights[cur] == 0 {
		return stack, 0
	}
	left := stack[len(stack)-1]

	area := heights[cur] * (right - left)
	return stack, area
}

type Node struct {
	Val      int
	Children []*Node
}

func Bfs(node *Node) []int {
	if node == nil {
		return []int{}
	}
	var arr []int
	queue := []*Node{node}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		arr = append(arr, cur.Val)
		queue = append(queue, cur.Children...)
	}
	return arr
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
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
