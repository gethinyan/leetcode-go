// 14. 最长公共前缀（https://leetcode-cn.com/problems/longest-common-prefix/）
package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	var ans string
	if len(strs) <= 0 {
		return ans
	}
	minLen := len(strs[0])
	// 先找最短字符串长度
	for i := 1; i < len(strs); i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
Loop:
	for j := 0; j < minLen; j++ {
		str := string(strs[0][j])
		for i := 1; i < len(strs); i++ {
			if string(strs[i][j]) != str {
				break Loop
			}
		}
		ans += str
	}
	return ans
}
