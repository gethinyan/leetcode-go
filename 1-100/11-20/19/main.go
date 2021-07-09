// 19. 删除链表的倒数第 N 个结点（https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/）
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
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	n := 2
	fmt.Println(removeNthFromEnd(head, n))
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	first, second := dummy, dummy
	for n >= 0 {
		first = first.Next
		n--
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
