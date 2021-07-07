// 13. 罗马数字转整数（https://leetcode-cn.com/problems/roman-to-integer/）
package main

import (
	"fmt"
)

func main() {
	s := "XIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	var ans int
	romanConvert := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if _, ok := romanConvert[string(s[i])+string(s[i+1])]; ok {
				ans += romanConvert[string(s[i])+string(s[i+1])]
				i++
				continue
			}
		}
		if _, ok := romanConvert[string(s[i])]; ok {
			ans += romanConvert[string(s[i])]
		}
	}
	return ans
}
