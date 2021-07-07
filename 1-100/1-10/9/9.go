// 9. 回文数（https://leetcode-cn.com/problems/palindrome-number/）
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	for l, r := 0, len(str)-1; l < r; l, r = l+1, r-1 {
		if str[l] != str[r] {
			return false
		}
	}
	return true
}
