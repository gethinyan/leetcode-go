// 57. 插入区间（https://leetcode-cn.com/problems/insert-interval/）
package main

import (
	"fmt"
)

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	newIntervals := [][]int{}
	if (len(intervals) == 0 || len(intervals[0]) == 0) && len(newInterval) == 0 {
		return ans
	}
	if len(newInterval) != 0 {
		if len(intervals) == 0 || len(intervals[0]) == 0 {
			newIntervals = [][]int{newInterval}
		}
		for i := 0; i < len(intervals); i++ {
			if intervals[i][0] > newInterval[0] {
				newIntervals = append(append(newIntervals, newInterval), intervals...)
				break
			} else if i == len(intervals)-1 {
				newIntervals = append(intervals, newInterval)
				break
			} else if intervals[i][0] <= newInterval[0] && intervals[i+1][0] >= newInterval[0] {
				newIntervals = append(append(append(newIntervals, intervals[:i+1]...), newInterval), intervals[i+1:]...)
				break
			}
		}
	}
	ans = append(ans, []int{newIntervals[0][0], newIntervals[0][1]})
	for i := 1; i < len(newIntervals); i++ {
		if ans[len(ans)-1][1] >= newIntervals[i][0] {
			if ans[len(ans)-1][1] < newIntervals[i][1] {
				ans[len(ans)-1][1] = newIntervals[i][1]
			}
		} else {
			ans = append(ans, []int{newIntervals[i][0], newIntervals[i][1]})
		}
	}
	return ans
}
