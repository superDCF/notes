package main

import (
	"log"
	"strings"
)

/*
给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
输入：s = "Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
*/

func main() {
	log.Println(reverseWords("abc 啊哈哈 cba"))
}
func reverseWords(s string) string {
	sb := strings.Split(s, " ")
	for i, v := range sb {
		sb[i] = reverseStr([]byte(v))
	}
	return strings.Join(sb, " ")
}

func reverseStr(s []byte) string {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		s[i], s[sLen-1-i] = s[sLen-1-i], s[i]
	}
	return string(s)
}
