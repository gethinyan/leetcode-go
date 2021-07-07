// 20. 有效的括号（https://leetcode-cn.com/problems/valid-parentheses/）
package main

import (
	"fmt"
)

func main() {
	s := "()"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	match := func(a rune, b rune) bool {
		switch {
		case a == '[' && b == ']':
			return true
		case a == '(' && b == ')':
			return true
		case a == '{' && b == '}':
			return true
		default:
			return false
		}
	}
	var i int
	stack := make([]rune, len(s)+1)
	for _, char := range s {
		if match(stack[i], char) {
			i--
		} else {
			i++
			stack[i] = char
		}
	}
	if i != 0 {
		return false
	}
	return true
}
