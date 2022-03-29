// 29. 两数相除（https://leetcode-cn.com/problems/divide-two-integers/）
package main

import (
	"fmt"
	"math"
)

func main() {
	dividend := 10
	divisor := 3
	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	symbol := "+"
	if dividend < 0 {
		dividend = -dividend
		divisor = -divisor
	}
	if divisor < 0 {
		symbol = "-"
		divisor = -divisor
	}
	if divisor > dividend {
		return 0
	}
	multiDivisor := divisor
	res := 1
	for dividend > multiDivisor+multiDivisor {
		res += res
		multiDivisor += multiDivisor
	}
	res += divide(dividend-multiDivisor, divisor)
	if symbol == "-" {
		res = -res
	}
	// 如果除法结果溢出，则返回 231 − 1
	if res > math.MaxInt32 || res < math.MinInt32 {
		return math.MaxInt32
	}

	return res
}
