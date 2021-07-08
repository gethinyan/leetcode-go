// 18. 四数之和（https://leetcode-cn.com/problems/4sum/）
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	numsLen := len(nums)
	for i := 0; i < numsLen-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < numsLen-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := numsLen - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case sum == target:
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] && nums[r] == nums[r-1] {
						l++
						r--
					}
					l++
					r--
				case sum < target:
					l++
				case sum > target:
					r--
				}
			}
		}
	}
	return result
}
