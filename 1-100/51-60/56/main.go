// 56. 合并区间（https://leetcode-cn.com/problems/merge-intervals/）
package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return ans
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	ans = append(ans, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < len(intervals); i++ {
		if ans[len(ans)-1][1] >= intervals[i][0] {
			if ans[len(ans)-1][1] < intervals[i][1] {
				ans[len(ans)-1][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return ans
}
