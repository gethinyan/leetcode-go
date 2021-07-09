// 25. K 个一组翻转链表（https://leetcode-cn.com/problems/reverse-nodes-in-k-group/）
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
	k := 3
	fmt.Println(reverseKGroup(head, k))
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	node1 := &ListNode{}
	node2 := newHead
	var i int
	for i = 0; i < k && nil != head; i++ {
		node1 = &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  head.Val,
				Next: node1.Next,
			},
		}
		head = head.Next
	}
	// 如果是达到 k 个节点退出的，才翻转
	if i == k {
		for i := k; i > 0; i-- {
			node2.Next.Val = node1.Next.Val
			node1 = node1.Next
			node2 = node2.Next
		}
		node1.Next = reverseKGroup(head, k)
	}
	return newHead.Next
}
