package o6

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	return intoArr(head, []int{})
}

func intoArr(head *ListNode, arr []int) []int {
	if head == nil {
		return arr
	}
	arr = intoArr(head.Next, arr)
	arr = append(arr, head.Val)
	return arr
}
