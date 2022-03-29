// 28. 实现 strStr（https://leetcode-cn.com/problems/implement-strstr/）
package main

import (
	"fmt"
)

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	lenHaystack := len(haystack)
	lenNeedle := len(needle)
	if 0 == lenNeedle {
		return 0
	}
	for i := 0; i < lenHaystack; i++ {
		if haystack[i] != needle[0] {
			continue
		}
		flag := false
		for j := 0; j < lenNeedle && i+j < lenHaystack; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == lenNeedle-1 {
				flag = true
			}
		}
		if true == flag {
			return i
		}
	}

	return -1
}
