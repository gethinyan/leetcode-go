// 3. 无重复字符的最长子串（https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/）
package main

import (
	"fmt"
)

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLength := 0
	sMap := make(map[rune]int)
	for key, value := range []rune(s) {
		if lastKey, ok := sMap[value]; ok && lastKey >= start {
			start = lastKey + 1
		}
		if (key - start + 1) > maxLength {
			maxLength = key - start + 1
		}
		sMap[value] = key
	}

	return maxLength
}
