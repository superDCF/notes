package main

import (
	"log"
)

/*
给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
*/

func main() {
	log.Println(repeatedSubstringPattern("abab"))
}
func repeatedSubstringPattern(s string) bool {
	sLen := len(s)
	if sLen < 2 {
		return false
	}
	s2 := s + s
	s2 = s2[1 : len(s2)-1]
	for i := 0; i < len(s2)-sLen+1; i++ {
		if s2[i:i+sLen] == s {
			return true
		}
	}
	return false
}
