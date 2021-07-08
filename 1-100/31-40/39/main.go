// 39. 组合总和（https://leetcode-cn.com/problems/combination-sum/）
package main

import (
	"fmt"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] == target {
			result = append(result, []int{candidates[i]})
		} else if candidates[i] < target {
			res := combinationSum(candidates[i:], target-candidates[i])
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
