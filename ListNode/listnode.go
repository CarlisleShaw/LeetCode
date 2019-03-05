package listnode

import "fmt"

// ListNode is singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(array []int) *ListNode {
	head := &ListNode{array[0], nil}
	node := head
	for i := 1; i < len(array); i++ {
		node.Next = &ListNode{array[i], nil}
		node = node.Next
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" ")
		head = head.Next
	}
	fmt.Println("")
}
