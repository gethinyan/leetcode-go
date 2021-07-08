// 48. 旋转图像（https://leetcode-cn.com/problems/rotate-image/）
package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)-i-1; i++ {
		for j := i; j < len(matrix)-i-1; j++ {
			// 交换四个边对应位置的元素
			matrix[i][j], matrix[j][len(matrix)-1-i], matrix[len(matrix)-1-i][len(matrix)-1-j], matrix[len(matrix)-1-j][i] = matrix[len(matrix)-1-j][i], matrix[i][j], matrix[j][len(matrix)-1-i], matrix[len(matrix)-1-i][len(matrix)-1-j]
		}
	}
}
