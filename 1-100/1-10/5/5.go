// 5. 最长回文子串（https://leetcode-cn.com/problems/longest-palindromic-substring/）
package main

import (
	"fmt"
)

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	newS := "^#"
	lenS := len(s)
	// 先插入#
	for i := 0; i < lenS; i++ {
		newS += string(s[i]) + "#"
	}
	lenNewS := len(newS)
	newS += "$"
	p := make([]int, lenNewS)
	index, mx, maxLen, maxStr := 0, 1, -1, ""
	for i := 1; i < lenNewS; i++ {
		if i < mx {
			p[i] = p[2*index-i]
			if p[2*index-i] > mx-i {
				p[i] = mx - i
			}
		} else {
			p[i] = 1
		}
		for {
			if newS[i-p[i]] != newS[i+p[i]] {
				break
			}
			p[i]++
		}
		if mx < i+p[i] {
			index = i
			mx = i + p[i]
		}
		if maxLen < p[i]-1 {
			maxLen = p[i] - 1
			maxStr = newS[i-maxLen : i+maxLen+1]
		}
	}
	result := ""
	lenMaxStr := len(maxStr)
	for i := 0; i < lenMaxStr; i++ {
		if string(maxStr[i]) != "#" {
			result += string(maxStr[i])
		}
	}

	return result
}
