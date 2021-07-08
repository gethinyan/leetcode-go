// 11. 盛最多水的容器（https://leetcode-cn.com/problems/container-with-most-water/）
package main

import (
	"fmt"
)

func main() {
	height := []int{4, 3, 2, 1, 4}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	var max, current int
	for l < r {
		if height[r] >= height[l] {
			current = height[l] * (r - l)
		} else {
			current = height[r] * (r - l)
		}
		if current > max {
			max = current
		}
		if height[r] >= height[l] {
			l++
		} else {
			r--
		}
	}
	return max
}
