/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/

package main

import "log"

func main() {
	s := "abcd"
	log.Println(s[:4])
	log.Println(longestCommonPrefix2([]string{"flower", "flow", "flight"}))
	log.Println(longestCommonPrefix2([]string{"flower", "flower", "flower", "flower"}))
	log.Println(longestCommonPrefix2([]string{"dog", "racecar", "car"}))
	log.Println(longestCommonPrefix2([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minLenIndex := 0
	for i := 1; i < len(strs); i++ {
		if minLenIndex > len(strs[i]) { // 明显有bug，还能测试通过
			minLenIndex = i
		}
	}
	if strs[minLenIndex] == "" {
		return ""
	}
	var prefix = ""
	for j := 0; j < len(strs[0]); j++ {
		prefix = string(strs[minLenIndex][:j+1])
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) < len(prefix) {
				return string(prefix[:len(prefix)-1])
			}
			if strs[i][:len(prefix)] != prefix {
				return string(prefix[:len(prefix)-1])
			}
		}
	}

	return string(prefix)
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLenIndex := 0
	for i := 1; i < len(strs); i++ {
		if len(strs[minLenIndex]) > len(strs[i]) {
			minLenIndex = i
		}
	}
	if strs[minLenIndex] == "" {
		return ""
	}
	lenth := 0
	for ; lenth < len(strs[minLenIndex]); lenth++ {
		char := strs[minLenIndex][lenth]
		for i := 0; i < len(strs); i++ {
			if i == minLenIndex {
				continue
			}
			if strs[i][lenth] != char {
				return strs[minLenIndex][:lenth]
			}
		}
	}
	return strs[minLenIndex][:lenth]
}
