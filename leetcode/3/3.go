package main

import "log"

func main() {
	// log.Println(lengthOfLongestSubstring("dvdf")) // "pwwkew" // abcabcbb //"au"
	// s := "pwwkew11发"
	// log.Println(len(s), s[0], s[len(s)-1], s[len(s)-2], s[len(s)-3])
	// log.Printf("%T", s[len(s)-2])
	log.Println(lengthOfLongestSubstring2("a"))
}

// s是 由英文字母、数字、符号和空格组成，所以不用考虑多字节字符
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var m = map[byte]int{} // 复用，存储每段不重复字符串的字符
	start := 0
	index := map[int]int{} // 用来存储不重复字符串的开始位置，和长度，map[长度]开始位置
	m[s[0]] = 0
	for i := 1; i < len(s); i++ {
		log.Println(start, i, m, index)
		if s[i-1] != s[i] {
			if k, ok := m[s[i]]; ok { // 判断是否重复
				index[i-start] = start
				start = k
				m[s[i]] = i
			} else {
				if start == i-1 {
					start = i - 1
				}
				m[s[i]] = i
				if i == len(s)-1 { // 保证最后一个加入
					index[i-start+1] = start
				}
			}

		} else {
			index[i-start] = start
			m = map[byte]int{
				s[i]: i,
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

// 滑动窗口
// abcabcbb
// abcabcbbace
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 1
	ret := string(s[0])
	for i := 1; i < len(s); i++ {
		existIndex := -1
		for j := 0; j < len(ret); j++ {
			if s[i] == ret[j] {
				existIndex = j
				break
			}
		}
		if existIndex > -1 {
			ret = ret[existIndex+1:] + string(s[i])
		} else {
			ret += string(s[i])
		}
		if len(ret) > maxLen {
			maxLen = len(ret)
		}
	}
	return maxLen
}
