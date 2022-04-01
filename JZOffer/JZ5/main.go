// 剑指 Offer 05. 替换空格（https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/）

package main

import (
	"fmt"
)

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	newS := ""
	for _, c := range s {
		if ' ' == c {
			newS += "%20"
		} else {
			newS += string(c)
		}
	}
	return newS
}
