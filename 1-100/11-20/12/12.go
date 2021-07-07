// 12. 整数转罗马数字（https://leetcode-cn.com/problemset/all/）
package main

import (
	"fmt"
)

func main() {
	num := 16
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	var ans string
	decimalNums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNums := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(decimalNums); i++ {
		for num >= decimalNums[i] {
			num -= decimalNums[i]
			ans += romanNums[i]
		}
	}
	return ans
}
