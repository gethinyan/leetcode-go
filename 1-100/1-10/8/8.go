// 8. 字符串转换整数（https://leetcode-cn.com/problems/string-to-integer-atoi/）
package main

import (
	"fmt"
	"math"
)

func main() {
	str := "4193 with words"
	fmt.Println(myAtoi(str))
}
func myAtoi(str string) int {
	var ans int
	flag := 1
	for i := 0; i < len(str); i++ {
		if ans > math.MaxInt32 || ans < math.MinInt32 {
			break
		}
		if (str[i] == ' ' || str[i] == '+') && (i == 0 || str[i-1] == ' ') {
			continue
		} else if str[i] == '-' && (i == 0 || str[i-1] == ' ') {
			flag = -1
		} else if str[i] >= '0' && str[i] <= '9' {
			ans = 10*ans + int(str[i]-'0')
			fmt.Println(ans)
		} else {
			break
		}
	}
	ans = ans * flag
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}
