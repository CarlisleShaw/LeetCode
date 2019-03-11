// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

package main

import (
	. "LeetCode/ListNode"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0
	for ; curr != nil && count != k; count++ {
		curr = curr.Next
	}
	if count != k {
		return head
	}
	curr = reverseKGroup(curr, k)
	for ; count > 0; count-- {
		// tmp := head.Next
		// head.Next = curr
		// curr = head
		// head = tmp
		head.Next, curr, head = curr, head, head.Next
	}
	head = curr
	return head
}

func main() {

}
