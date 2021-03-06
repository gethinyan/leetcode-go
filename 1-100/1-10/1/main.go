// 1. 两数之和（https://leetcode-cn.com/problems/two-sum/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	slice := make([]int, 0, 2)
	for key, value := range nums {
		if _, ok := numsMap[target-value]; ok {
			return append(slice, key, numsMap[target-value])
		}
		numsMap[value] = key
	}

	return nil
}
