// 23. 合并K个升序链表（https://leetcode-cn.com/problems/merge-k-sorted-lists/）
package main

import (
	"fmt"
	"strconv"
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
	lists := []*ListNode{l1, l2}

	fmt.Println(mergeKLists(lists))
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

func mergeKLists(lists []*ListNode) *ListNode {
	result := &ListNode{}
	resultCopy := result
	var key, minVal int
	var flag bool
	for {
		key = -1
		minVal = -1
		flag = false
		var newLists []*ListNode
		for _, list := range lists {
			if list == nil {
				continue
			}
			newLists = append(newLists, list)
		}
		lists = newLists
		if len(lists) <= 0 {
			break
		}
		resultCopy.Next = &ListNode{}
		resultCopy = resultCopy.Next
		for index, list := range lists {
			if list == nil {
				lists = append(lists[:index], lists[index+1:]...)
				continue
			}
			if minVal == -1 && flag == false {
				minVal = list.Val
				key = index
				flag = true

				continue
			}
			if list.Val < minVal {
				minVal = list.Val
				key = index
			}
		}
		if key != -1 {
			lists[key] = lists[key].Next
		}
		resultCopy.Val = minVal
	}

	return result.Next
}
