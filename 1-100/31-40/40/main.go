// 40. 组合总和 II（https://leetcode-cn.com/problems/combination-sum-ii/）
package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	for i := 0; i < len(candidates); i++ {
		if i > 0 && (candidates[i] == candidates[i-1]) {
			continue
		}
		if candidates[i] == target {
			result = append(result, []int{candidates[i]})
		} else if candidates[i] < target {
			res := combinationSum2(candidates[i+1:], target-candidates[i])
			for j := range res {
				if len(res[j]) == 0 {
					continue
				}
				result = append(result, append([]int{candidates[i]}, res[j]...))
			}
		}
	}
	return result
}
