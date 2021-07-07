// 4. 寻找两个正序数组的中位数（https://leetcode-cn.com/problems/median-of-two-sorted-arrays/）
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsAll := []int{}
	startNum1, startNum2 := 0, 0
	for {
		if len(nums1[startNum1:]) == 0 {
			numsAll = append(numsAll, nums2[startNum2:]...)
			break
		}
		if len(nums2[startNum2:]) == 0 {
			numsAll = append(numsAll, nums1[startNum1:]...)
			break
		}
		if nums1[startNum1] < nums2[startNum2] {
			numsAll = append(numsAll, nums1[startNum1])
			startNum1++
		} else {
			numsAll = append(numsAll, nums2[startNum2])
			startNum2++
		}
	}
	lenNumsAll := len(numsAll)
	return float64(numsAll[lenNumsAll/2]+numsAll[(lenNumsAll-1)/2]) / float64(2)
}
