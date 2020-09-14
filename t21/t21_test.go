package t21

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}

	p1 := arrToList(arr1)
	p2 := arrToList(arr2)
	printList(p1)
	printList(p2)
	printList(mergeTwoLists(p1, p2))

	//arr1 = []int{}
	//arr2 = []int{}
	//
	//p1 = arrToList(arr1)
	//p2 = arrToList(arr2)
	//printList(mergeTwoLists(p1, p2))
	//
	//arr1 = []int{1, 3, 5, 7, 9}
	//arr2 = []int{2, 4}
	//
	//p1 = arrToList(arr1)
	//p2 = arrToList(arr2)
	//printList(mergeTwoLists(p1, p2))
	//
	//arr1 = []int{1, 2, 3, 4}
	//arr2 = []int{}
	//
	//p1 = arrToList(arr1)
	//p2 = arrToList(arr2)
	//printList(mergeTwoLists(p1, p2))
	//
	//arr1 = []int{1, 2, 3, 4}
	//arr2 = []int{5}
	//
	//p1 = arrToList(arr1)
	//p2 = arrToList(arr2)
	//printList(mergeTwoLists(p1, p2))
}

func arrToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	p := &ListNode{
		Val:  arr[0],
		Next: arrToList(arr[1:]),
	}
	return p
}

func printList(head *ListNode) {
	if head != nil {
		fmt.Printf("%d, ", head.Val)
		printList(head.Next)
		return
	}
	fmt.Println("nil")
}
