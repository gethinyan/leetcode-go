// 剑指 Offer 03. 数组中重复的数字（https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/）

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
	numsMap := map[int]int{}
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			return num
		}

		numsMap[num] = 1
	}

	return -1
}
