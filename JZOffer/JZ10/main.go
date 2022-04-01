// 剑指 Offer 10- I. 斐波那契数列（https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/）
// 剑指 Offer 10- II. 青蛙跳台阶问题（https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/）

package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(95))
	fmt.Println(numWays(10))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	f0, fn := 0, 1
	n -= 2
	for ; n >= 0; n-- {
		f0, fn = fn, f0+fn
		fn %= 1000000007
	}

	return fn
}

func numWays(n int) int {
	f0, fn := 1, 1
	n -= 2
	for ; n >= 0; n-- {
		f0, fn = fn, f0+fn
		fn %= 1000000007
	}

	return fn
}
