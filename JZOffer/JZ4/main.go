// 剑指 Offer 04. 二维数组中的查找（https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/）

package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(matrix, 11))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	height := len(matrix)
	if height == 0 {
		return false
	}
	width := len(matrix[0])
	if width == 0 {
		return false
	}
	i := height - 1
	j := 0
	for i >= 0 && j < len(matrix[i]) {
		if target == matrix[i][j] {
			return true
		} else if target < matrix[i][j] {
			i--
		} else if target > matrix[i][j] {
			j++
		}
	}

	return false
}
