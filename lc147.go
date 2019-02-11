package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

// 思路：新建一个新链表，存放已排序的部分节点，
// 源链表的每一个节点（外循环）依次与新链表中已排序的每个节点比大小（内循环），
// 找到适合的位置插入新链表。
func insertionSortList(head *ListNode) *ListNode {
	// 情景：需要在头结点前插入节点的情况。
	// 出现问题：最后难以确定链表的头指针，而该方法最终需要返回链表的头指针。
	// 解决方案：新建dummy结点作为辅助头结点，dummy.Next 即为最终返回的链表的头指针。
    dummy := new(ListNode)

	for p2 := head; p2 != nil ; {
		// 寻找合适的插入位置
		p1 := dummy
		for ; p1.Next != nil && p2.Val >= p1.Next.Val; p1 = p1.Next { }

		if p1.Next == nil {
			// 要插入的值比链表中的值都要大，插在尾部
			head = p2.Next	// 借用head暂存下一轮循环会用到的源链表的头指针
			p1.Next = p2	// 与尾节点建立链接
			p2.Next = nil	// 只链接该节点，断后，不能拖着后面的其他节点
			p2 = head		// 恢复p2为源链表的头指针
		} else {
			// 找到链表中间的某一位置，插入，处理好前后关系
			head = p2.Next
			p2.Next = p1.Next
			p1.Next = p2
			p2 = head
		}
	}

	return dummy.Next
}

func initNode(val int) *ListNode {
	p := new(ListNode)
	p.Val = val
	p.Next = nil

	return p
}

func initNodeArray(val []int) *ListNode {
	p := new(ListNode)
	p.Val = val[0]
	p.Next = nil
	p = appendNode(p, val[1:])
	return p
}

func appendNode(head *ListNode, val []int) *ListNode {
	p := head
	for ; p.Next != nil; p = p.Next { }
	for i := 0; i < len(val); i++ {
		p.Next = initNode(val[i])
		p = p.Next
	}

	return head
}

func main() {
	head := initNodeArray([]int{11,1,3,5,2,4,9,6,8,7})

	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

	head = insertionSortList(head)

	fmt.Println("after sorted")
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}

}