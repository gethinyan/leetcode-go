// 54. 螺旋矩阵（https://leetcode-cn.com/problems/spiral-matrix/）
package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}
	l, t, r, b := 0, 0, len(matrix[0])-1, len(matrix)-1
	for l <= r && t <= b {
		// 遍历上面的行
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[t][i])
		}
		t++
		// 遍历右边的列
		for i := t; i <= b; i++ {
			ans = append(ans, matrix[i][r])
		}
		r--
		// 遍历下面的行，需避免重复遍历
		if t <= b {
			for i := r; i >= l; i-- {
				ans = append(ans, matrix[b][i])
			}
		}
		b--
		// 遍历左边的列，需避免重复遍历
		if l <= r {
			for i := b; i >= t; i-- {
				ans = append(ans, matrix[i][l])
			}
		}
		l++
	}
	return ans
}
