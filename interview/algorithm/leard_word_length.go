/*
给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。
假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。
注意：每次拼写（指拼写词汇表中的一个单词）时，chars 中的每个字母都只能用一次。
返回词汇表 words 中你掌握的所有单词的 长度之和。



示例 1：
输入：words = ["cat","bt","hat","tree"], chars = "atach"
输出：6
解释：
可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。
示例 2：
输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"
输出：10
解释：
可以形成字符串 "hello" 和 "world"，所以答案是 5 + 5 = 10。

提示：
1 <= words.length <= 1000
1 <= words[i].length, chars.length <= 100
所有字符串中都仅包含小写英文字母

*/

package main

import (
	"log"
)

func main() {
	// log.Println(WordsCharLen([]string{"cat", "bt", "hat", "tree"}, "atach"))
	log.Println(WordsCharLen([]string{"hello","world","leetcode"}, "welldonehoneyr"))
}

func WordsCharLen(words []string, chars string) int {
	if len(words) == 0 || chars == "" {
		return 0
	}
	totalLen := 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		charsMap := NewCharsMap(chars) // todo
		isLearn := true
		for j := 0; j < len(word); j++ {
			if count, ok := charsMap[string(word[j])]; ok && count > -1 {
				charsMap[string(word[j])] = charsMap[string(word[j])] - 1
			} else { // 表示charsMap不存在该字符，直接break，循环判断下一个单词
				isLearn = false
				break
			}
		}
		if isLearn {
			totalLen += len(word)
		}
	}
	return totalLen
}

func NewCharsMap(chars string) map[string]int {
	charsMap := make(map[string]int)
	for _, v := range chars {
		if _, ok := charsMap[string(v)]; ok {
			charsMap[string(v)]++
		} else {
			charsMap[string(v)] = 1
		}
	}
	return charsMap
}
