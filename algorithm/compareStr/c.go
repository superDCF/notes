/*
输入一个字符串，判断在字符“people”的任意位置，增加任意字符，是否能够等于输入的字符串
1. 如果输入的长度小于“people”，直接返回false
2. 判断是否输入的字符串是否不存在“people”其中之一的
3. 判断输入字符出现的顺序是否和people一致，双指针法
*/

package main

import "log"

func main() {
	log.Println(compare("appaeople"))
	log.Println(compare("paaaeople"))
	log.Println(compare("aaaeople"))
}

const people = "people"

func compare(str string) bool {
	if len(str) < len(people) {
		return false
	}
	s1, s2 := 0, 0
	for _ = range str {
		if str[s1] == people[s2] {
			s1++
			s2++
		} else {
			s1++
		}
	}
	if s1-s2 == len(str)-len(people) {
		return true
	}
	return false
}
