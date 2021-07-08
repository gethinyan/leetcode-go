// 26. 删除有序数组中的重复项（https://leetcode-cn.com/problems/reverse-nodes-in-k-group/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	var i, j int
	numLen := len(nums)
	if numLen < 1 {
		return 0
	}
	for ; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
