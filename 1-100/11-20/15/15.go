// 15. 三数之和（https://leetcode-cn.com/problems/3sum/）
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			switch {
			case nums[l]+nums[r]+nums[i] == 0:
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l+1] == nums[l] && nums[r-1] == nums[r] {
					l++
					r--
				}
				l++
				r--
			case nums[l]+nums[r]+nums[i] < 0:
				l++
			case nums[l]+nums[r]+nums[i] > 0:
				r--
			}
		}
	}
	return result
}
