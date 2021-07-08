// 45. 跳跃游戏 II（https://leetcode-cn.com/problems/jump-game-ii/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	step := 1
	for i := 0; i < len(nums)-1; {
		if i+nums[i] >= len(nums)-1 {
			break
		}
		max := 0
		t := i
		for j := t + 1; j <= t+nums[t] && j < len(nums); j++ {
			if nums[j]+j-t > max {
				max = nums[j] + j - t
				i = j
			}
		}
		step++
	}
	return step
}
