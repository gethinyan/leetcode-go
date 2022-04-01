// 剑指 Offer 11. 旋转数组的最小数字（https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/）

package main

import (
	"fmt"
)

func main() {
	numbers := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(numbers))
}

func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		mid := low + (high-low)/2
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}

	return numbers[low]
}
