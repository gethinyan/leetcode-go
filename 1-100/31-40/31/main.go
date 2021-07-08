// 31. 下一个排列（https://leetcode-cn.com/problems/next-permutation/）
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	if i != 0 {
		// 如果整个数组不是倒序的，那么从断点处后面找出一个刚好比断点处值大的数交换
		for j := len(nums) - 1; j >= i; j-- {
			if nums[j] > nums[i-1] {
				nums[j], nums[i-1] = nums[i-1], nums[j]
				break
			}
		}
	}
	// 交换后断点处后面的元素是降序的，需要调整为升序
	for k := i; k <= (len(nums)-1+i)/2; k++ {
		nums[k], nums[len(nums)-1+i-k] = nums[len(nums)-1+i-k], nums[k]
	}
}
