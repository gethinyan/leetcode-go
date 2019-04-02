package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	L := &ListNode{}
	current := L
	carry := 0
	nullNode := &ListNode{}
	for {
		if l1 == nil {
			l1 = nullNode
		}
		if l2 == nil {
			l2 = nullNode
		}
		current.Val = l1.Val + l2.Val + carry
		carry = 0
		if current.Val >= 10 {
			current.Val -= 10
			carry = 1
		}
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		current.Next = &ListNode{}
		current = current.Next
	}

	return L
}

func main() {
	l1 := &ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}
	l2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}
	fmt.Println(addTwoNumbers(l1, l2))
}