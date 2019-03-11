// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
package main

import . "LeetCode/ListNode"

func swapPairs(head *ListNode) *ListNode {
	tmphead := &ListNode{Val: 0, Next: head}
	tmp := tmphead
	for tmp.Next != nil && tmp.Next.Next != nil {
		a := tmp.Next
		b := a.Next
		c := b.Next
		tmp.Next = b
		b.Next = a
		a.Next = c
		tmp = tmp.Next.Next
	}
	return tmphead.Next
}

func main() {
	head := MakeList([]int{1, 2, 3, 4, 5, 6, 7})
	// head := MakeList([]int{1})
	PrintList(head)
	swaphead := swapPairs(head)
	PrintList(swaphead)
}
