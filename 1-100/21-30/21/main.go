// 21. 合并两个有序链表（https://leetcode-cn.com/problems/merge-two-sorted-lists/）
package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(mergeTwoLists(l1, l2))
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := &ListNode{}
	if l1.Val >= l2.Val {
		result.Val = l2.Val
		result.Next = mergeTwoLists(l1, l2.Next)
	} else {
		result.Val = l1.Val
		result.Next = mergeTwoLists(l1.Next, l2)
	}

	return result
}
