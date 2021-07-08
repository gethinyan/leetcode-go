// 35. 搜索插入位置（https://leetcode-cn.com/problems/search-insert-position/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			l, r = mid, mid
		}
		if nums[mid] < target {
			l = mid + 1
		}
		if nums[mid] > target {
			r = mid - 1
		}
	}
	if target <= nums[l] {
		return l
	}
	return l + 1
}
