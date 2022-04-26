package main

import "log"

/*
给你一个字符串 s，找到 s 中最长的回文子串。
*/
func main() {
	log.Println(longestPalindrome("cbbd"))
}

// 最长回文字符串
func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}
	start := 0
	end := 0
	for i := 0; i < sLen; i++ {
		left, right := expandCompare(s, i, i)
		log.Println(1, left, right)
		if right-left > end-start {
			start = left
			end = right
		}
		left, right = expandCompare(s, i, i+1)
		log.Println(2, left, right)
		if right-left > end-start {
			start = left
			end = right
		}
	}
	log.Println(start, end)
	return string(s[start : end+1])
}

// 中心扩散法，指定一个待确认的字符串，然后向两边扩散
func expandCompare(s string, left, right int) (int, int) {
	sLen := len(s)
	for left >= 0 && right <= sLen-1 {
		if s[left] != s[right] {
			return left + 1, right - 1
		}
		left--
		right++
	}
	return left + 1, right - 1
}
