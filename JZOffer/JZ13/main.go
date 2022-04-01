// 剑指 Offer 13. 机器人的运动范围（https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/）

package main

import (
	"fmt"
)

func main() {
	m, n, k := 16, 8, 4
	fmt.Println(movingCount(m, n, k))
}

func movingCount(m int, n int, k int) int {
	flag := make([][]bool, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]bool, n)
	}

	flag[0][0] = true

	count := 1
	for i := 0; i < m; i++ {
		sumI := sumDigit(i)
		for j := 0; j < n; j++ {
			if (i > 0 && flag[i-1][j]) || (j > 0 && flag[i][j-1]) {
				sumJ := sumDigit(j)
				if sumI+sumJ <= k {
					flag[i][j] = true
					count++
				}
			}
		}
	}

	return count
}

func sumDigit(number int) int {
	sum := 0
	for number >= 1 {
		sum += number % 10
		number /= 10
	}

	return sum
}
