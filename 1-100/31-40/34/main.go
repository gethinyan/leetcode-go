// 34. 在排序数组中查找元素的第一个和最后一个位置（https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			l, r = mid, mid
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	for l > 0 && nums[l-1] == target {
		l--
	}
	for r < len(nums)-1 && nums[r+1] == target {
		r++
	}
	if len(nums) == 0 || nums[l] != target || nums[r] != target {
		return []int{-1, -1}
	}
	return []int{l, r}
}
