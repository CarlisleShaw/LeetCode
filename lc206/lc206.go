package main

import . "LeetCode/ListNode"

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		// temp := cur.Next
		// cur.Next = prev
		// prev = cur
		// cur = temp
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func main() {
	head := MakeList([]int{1, 2, 3, 4, 5, 6, 7})
	PrintList(head)
	rhead := reverseList(head)
	PrintList(rhead)
}
