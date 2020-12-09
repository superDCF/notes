package main

import "log"

func main() {
	log.Println(lengthOfLongestSubstring("dvdf")) // "pwwkew" // abcabcbb //"au"
	s := "pwwkew11发"
	log.Println(len(s), s[0], s[len(s)-1], s[len(s)-2], s[len(s)-3])
	log.Printf("%T", s[len(s)-2])
	// for i, v := range s {
	// 	// log.Println(i, v)
	// }
}

// s是 由英文字母、数字、符号和空格组成，所以不用考虑多字节字符
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var m = map[byte]struct{}{} // 复用，存储每段不重复字符串的字符
	start := 0
	index := map[int]int{} // 用来存储不重复字符串的开始位置，和长度，map[长度]开始位置
	m[s[0]] = struct{}{}
	for i := 1; i < len(s); i++ {
		log.Println(start, i, m, index)
		if s[i-1] != s[i] {
			if _, ok := m[s[i]]; ok { // 判断是否重复
				index[i-start] = start
				start = i
				m = map[byte]struct{}{}
			} else {
				if start == i-1 {
					start = i - 1
				}
				m[s[i]] = struct{}{}
				if i == len(s) - 1 {
					index[i-start+1] = start
				}
			}

		} else {
			index[i-start] = start
			m = map[byte]struct{}{
				s[i]: struct{}{},
			}
			start = i
		}
		
	}
	maxLen := 1
	for k := range index {
		if k > maxLen {
			maxLen = k
		}
	}
	return maxLen
}
