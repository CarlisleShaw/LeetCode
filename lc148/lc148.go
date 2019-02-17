package main

import "fmt"
// import "math"
// import "sort"

type ListNode struct {
	Val int
	Next *ListNode
}

func appendNode(head *ListNode, val int) *ListNode {
	n := head
	for ; n.Next != nil; n = n.Next { }
	n.Next = new(ListNode)
	n = n.Next
	n.Val = val
	n.Next = nil
	return head
}

func insertNode(node *ListNode, val int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = node.Next
	node.Next = newNode
	return node
}

func initHead(val int) *ListNode {
	head := new(ListNode)
	head.Val = val
	head.Next = nil
	return head
}

func mergeSortedList(L1 *ListNode, L2 *ListNode) *ListNode{
	dummy := initHead(-1)
	dummy.Next = L1
	p1 := dummy
	p2 := L2

	for ; p1.Next != nil && p2 != nil; {
		if p1.Next.Val >= p2.Val {
			L2 = p2.Next
			p2.Next = p1.Next
			p1.Next = p2
			p1 = p2
			p2 = L2
		} else {
			p1 = p1.Next
		}
	}

	if p1.Next == nil {
		p1.Next = p2
	}

	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head

	for ; fast.Next != nil && fast.Next.Next != nil; {
		fast = fast.Next.Next
		slow = slow.Next
	}

	leftHead := head
	rightHead := slow.Next
	slow.Next = nil

	leftHead = sortList(leftHead)
	rightHead = sortList(rightHead)

	return mergeSortedList(leftHead, rightHead)
}


func main() {
	// fmt.Println("")
	head := initHead(0)
	appendNode(head, 4)
	appendNode(head, 2)
	appendNode(head, 1)
	appendNode(head, 5)
	appendNode(head, 3)
	appendNode(head, 7)
	appendNode(head, 6)

	for n := head; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
	fmt.Println("")
	head = sortList(head)

	for n := head; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}

