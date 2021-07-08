// 41. 缺失的第一个正数（https://leetcode-cn.com/problems/first-missing-positive/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	var i, j int
	for ; i < len(nums); i++ {
		for nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for ; j < len(nums) && j+1 == nums[j]; j++ {
	}
	return j + 1
}
