package main

import . "LeetCode/ListNode"

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	slow := head
	fast := head.Next.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func main() {

}
