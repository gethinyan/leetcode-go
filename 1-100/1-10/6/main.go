// 6. Z 字形变换（https://leetcode-cn.com/problems/zigzag-conversion/）
package main

import (
	"fmt"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	var ans string
	if numRows <= 1 {
		return s
	}
	numRows--
	for i := 0; i <= numRows; i++ {
		for j := i; j < len(s); {
			ans += string(s[j])
			j += 2 * (numRows - j%numRows)
		}
	}
	return ans
}
