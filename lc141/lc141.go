// https://leetcode-cn.com/problems/linked-list-cycle/submissions/
package main

import (
	. "LeetCode/ListNode"

	. "github.com/deckarep/golang-set"
)

// 快慢指针有环必相遇
// 时间复杂度：O(n)
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	// 点前面的对象必须判断是否为nil
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 走过的节点存入Set，查询经过的节点在Set中是否存在
// 时间复杂度：O(n*1)
func hasCycleSet(head *ListNode) bool {
	var nodeSet Set
	for head != nil {
		if nodeSet.Contains(head) {
			return true
		}
		nodeSet.Add(head)
		head = head.Next
	}
	return false
}

func main() {

}
