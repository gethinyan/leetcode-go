package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsAll := []int{}
	startNums1, startNums2 := 0, 0
	for {
		if len(nums1[startNums1:]) == 0 {
			numsAll = append(numsAll, nums2[startNums2:]...)
			break
		}
		if len(nums2[startNums2:]) == 0 {
			numsAll = append(numsAll, nums1[startNums1:]...)
			break
		}
		if nums1[startNums1] < nums2[startNums2] {
			numsAll = append(numsAll, nums1[startNums1])
			++startNums1
		} else {
			numsAll = append(numsAll, nums2[startNums2])
			++startNums2
		}
	}
	lenNumsAll := len(numsAll)
	return float64(numsAll[lenNumsAll/2]+numsAll[(lenNumsAll-1)/2]) / float64(2)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
