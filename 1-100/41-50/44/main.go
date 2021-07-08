// 44. 通配符匹配（https://leetcode-cn.com/problems/wildcard-matching/）
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
	// 初始化
	dp := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true
	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	// s第i个位置，p第j个位置
	// 如果s[i]==p[j]||p[j]=="?"，则dp[i][j]=dp[i-1][j-1]
	// 如果p[j]=="*"，则dp[i][j]=dp[i-1][j]||dp[i][j-1]
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[sLen][pLen]
}
