// 剑指 Offer 12. 矩阵中的路径（https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/）

package main

import (
	"fmt"
)

func main() {
	// board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	// word := "ABCCED"
	board := [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}
	word := "AAB"
	fmt.Println(exist(board, word))
	fmt.Println(exist1(board, word))
}

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	if n*m < len(word) {
		return false
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 终止条件
		if k >= len(word) {
			return true
		}
		// 终止条件，坐标范围错误或者 i j 坐标对应 board 的值不等于 k 坐标对应 word 的值
		if i < 0 || j < 0 || i >= n || j >= m || board[i][j] != word[k] {
			return false
		}
		// 如果往回前找，不会找到相同字符，如 ABAB，k = 2 时，往前往后都是 B ；
		// 将 B 修改（剪枝）为不存在的字符，杜绝往回找成功的可能性。
		board[i][j] = '0'
		res := dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i+1, j, k+1) || dfs(i-1, j, k+1)
		// 找完了就改回来
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func exist1(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}
	for i, row := range board {
		for j, _ := range row {
			if true == nearbyExist(board, word, i, j) {
				return true
			}
		}
	}

	return false
}

func nearbyExist(board [][]byte, word string, i int, j int) bool {
	if word == "" {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || word[0] != board[i][j] {
		return false
	}
	board[i][j] = ' '
	result := nearbyExist(board, word[1:], i-1, j) ||
		nearbyExist(board, word[1:], i+1, j) ||
		nearbyExist(board, word[1:], i, j-1) ||
		nearbyExist(board, word[1:], i, j+1)
	board[i][j] = word[0]

	return result

}
