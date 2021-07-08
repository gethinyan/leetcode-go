// 17. 电话号码的字母组合（https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/）
package main

import (
	"fmt"
)

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	var ans []string
	numConvert := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	for i := 0; i < len(digits); i++ {
		var tmp []string
		if _, ok := numConvert[digits[i]]; !ok {
			break
		}
		for j := 0; j < len(numConvert[digits[i]]); j++ {
			if len(ans) <= 0 {
				tmp = append(tmp, string(numConvert[digits[i]][j]))
			} else {
				for k := 0; k < len(ans); k++ {
					tmp = append(tmp, ans[k]+string(numConvert[digits[i]][j]))
				}
			}
		}
		ans = tmp
	}
	return ans
}
