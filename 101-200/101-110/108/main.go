// 108. 将有序数组转换为二叉搜索树（https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Printf("%+v", sortedArrayToBST(nums))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	mid := len(nums) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
