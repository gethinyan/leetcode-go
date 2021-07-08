// 16. 最接近的三数之和（https://leetcode-cn.com/problems/3sum-closest/）
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	var closest int
	sort.Ints(nums)
	if len(nums) >= 3 {
		closest = nums[0] + nums[1] + nums[2]
	}
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			s := nums[i] - target + nums[l] + nums[r]
			if math.Abs(float64(s)) < math.Abs(float64(closest-target)) {
				closest = s + target
			}
			if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return closest
}
