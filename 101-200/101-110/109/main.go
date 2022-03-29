// 109. 有序链表转换二叉搜索树（https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/）
package main

import (
	"fmt"
)

func main() {
	head := &ListNode{
		Val: 0,
		Next: &ListNode{
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
		},
	}
	fmt.Println(sortedListToBST(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 快慢指针
func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func buildTree(left *ListNode, right *ListNode) *TreeNode {
	if left == nil || left == right {
		return nil
	}
	slow, fast := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return &TreeNode{
		Val:   slow.Val,
		Left:  buildTree(left, slow),
		Right: buildTree(slow.Next, right),
	}
}
