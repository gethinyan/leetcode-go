// 7. 整数反转（https://leetcode-cn.com/problems/reverse-integer/）
package main

import (
	"fmt"
	"math"
)

func main() {
	x := 123
	fmt.Println(reverse(x))
}
func reverse(x int) int {
	var ans int
	flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	for x >= 1 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > math.MaxInt32 {
		ans = 0
	}
	return ans * flag
}
