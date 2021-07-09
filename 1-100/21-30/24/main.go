// 24. 两两交换链表中的节点（https://leetcode-cn.com/problems/swap-nodes-in-pairs/）
package main

import (
	"fmt"
	"strconv"
)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(swapPairs(head))
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	var next string
	if nil != node.Next {
		next = node.Next.String()
	} else {
		next = "nil"
	}
	return "{Val:" + strconv.Itoa(node.Val) + " Next:" + next + "}"
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
