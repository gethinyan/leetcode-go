// 剑指 Offer 07. 重建二叉树（https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/）

package main

import (
	"fmt"
)

func main() {
	preorder, inorder := []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, inorder))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := preorder[0]
	// 找到中序遍历分割点
	var mid int
	for i, item := range inorder {
		if item == root {
			mid = i
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  buildTree(preorder[1:mid+1], inorder[:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:]),
	}
}
