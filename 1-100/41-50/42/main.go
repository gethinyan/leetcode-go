// 42. 接雨水（https://leetcode-cn.com/problems/trapping-rain-water/）
package main

import (
	"fmt"
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r, maxL, maxR, capacity := 0, len(height)-1, height[0], height[len(height)-1], 0
	for l < r {
		if height[l] < height[r] {
			if maxL > height[l] {
				capacity += maxL - height[l]
			} else {
				maxL = height[l]
			}
			l++
		} else {
			if maxR > height[r] {
				capacity += maxR - height[r]
			} else {
				maxR = height[r]
			}
			r--
		}
	}
	return capacity
}
