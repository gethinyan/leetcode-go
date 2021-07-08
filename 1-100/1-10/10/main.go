// 10. 正则表达式匹配（https://leetcode-cn.com/problems/regular-expression-matching/）
package main

import (
	"fmt"
)

func main() {
	s := "aa"
	p := "a"
	fmt.Println(isMatch(s, p))
}
func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	dp := make([][]bool, sLen+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true
	for j := 1; j <= pLen; j++ {
		if j > 1 && p[j-1] == '*' && p[j-2] != '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if j > 1 && p[j-1] == '*' && p[j-2] != '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2] || dp[i][j-1]
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j-1] || dp[i-1][j]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[sLen][pLen]
}
