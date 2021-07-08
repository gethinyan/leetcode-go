// 27. 移除元素（https://leetcode-cn.com/problems/remove-element/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	var i, j int
	for j < len(nums) {
		if nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	return i
}
