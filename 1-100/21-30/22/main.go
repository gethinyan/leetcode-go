// 22. 括号生成（https://leetcode-cn.com/problems/generate-parentheses/）
package main

import (
	"fmt"
)

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	result := addCharacter(0, 0, n, []string{""})

	return result
}

func addCharacter(left, right, n int, data []string) []string {
	if left >= n && right >= n {
		return data
	}
	result := []string{}
	// 右括号 +1（前提是左括号数量大于右括号）
	if right < left {
		tmp := []string{}
		tmp2 := data
		for _, item2 := range tmp2 {
			tmp = append(tmp, item2+")")
		}
		result = append(result, addCharacter(left, right+1, n, tmp)...)
	}
	// 左括号 +1
	if left < n {
		tmp := []string{}
		tmp1 := data
		for _, item1 := range tmp1 {
			tmp = append(tmp, item1+"(")
		}
		// fmt.Println(left)
		result = append(result, addCharacter(left+1, right, n, tmp)...)
	}
	return result
}
