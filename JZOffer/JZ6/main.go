// 剑指 Offer 06. 从尾到头打印链表（https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/）

package main

import (
	"fmt"
)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	fmt.Println(reversePrint(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
