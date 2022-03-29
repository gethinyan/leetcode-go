// 110. 平衡二叉树（https://leetcode-cn.com/problems/balanced-binary-tree/）
package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isBalanced(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(height(root.Left)), float64(height(root.Right)))) + 1
}
