// 55. 跳跃游戏（https://leetcode-cn.com/problems/jump-game/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	for i := 0; i < len(nums)-1; {
		if i+nums[i] >= len(nums)-1 {
			return true
		}
		max := 0
		t := i
		for j := t + 1; j <= t+nums[t] && j < len(nums); j++ {
			if nums[j]+j-t > max {
				max = nums[j] + j - t
				i = j
			}
		}
		if i == t {
			return false
		}
	}
	return false
}
