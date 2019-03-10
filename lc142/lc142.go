// https://leetcode-cn.com/problems/linked-list-cycle-ii/
package main

import (
	. "LeetCode/ListNode"
)

// O(n)解法
// x1 := head到环入口节点的距离
// x2 := 环入口节点到快慢指针相遇节点的距离
// x3 := 快慢指针相遇节点往后走到环入口节点的距离
// 慢指针走过的距离为 x1 + x2
// 快指针走过的距离为 x1 + x2 + x3 + x2
// 又：快指针走过的距离 == 2 * 慢指针走过的距离
// => x1 + x2 + x3 + x2 == 2 * (x1 + x2)
// => x1 == x3
// 即：一指针从head每次往后走一步。另一指针从快慢指针相遇节点每次往后走一步
// 两指针必然在环入口节点相遇
func detectCycle(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	hasCycle := false
	for p1 != nil && p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	for p1 = head; p1 != p2; {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

// O(n^2)解法，朴素，容易想到
// 先用快慢指针判断是否有环，并记录慢指针经过的节点数
// 从head开始遍历，查询每一个节点往后有没有相同的节点，查询次数为2*慢指针经过的节点数
// func detectCycleOn2(head *ListNode) *ListNode {
// 	loop := hasCycle(head) * 2
// 	test1 := head
// 	for j := loop; j > 0; j-- {
// 		test2 := test1.Next
// 		for i := loop; i > 0; i-- {
// 			if test2 == test1 {
// 				return test2
// 			}
// 			test2 = test2.Next
// 		}
// 		test1 = test1.Next
// 	}
// 	return nil
// }

// func hasCycle(head *ListNode) int {
// 	slow := head
// 	fast := head
// 	for i := 1; slow != nil && fast != nil && fast.Next != nil; i++ {
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		if slow == fast {
// 			return i
// 		}
// 	}
// 	return -1
// }

func main() {

}
