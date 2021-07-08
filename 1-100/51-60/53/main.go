// 53. 最大子序和（https://leetcode-cn.com/problems/maximum-subarray/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
