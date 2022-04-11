/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
[{()}]

*/

package main

import "log"

func main() {
	log.Println(isValid("{{}[]}"))
}

var table = map[byte]byte{
	']': '[',
	')': '(',
	'}': '{',
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else if s[i] == '(' {
			stack = append(stack, ')')
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
