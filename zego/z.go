/*
输入字符串，判断在“people”任意加字符的情况下，能否等于输入的字符串
*/

// step1: 先判断输入的字符串有没有people这几个字符串，如果没有直接返回false
// step2：如果有，再对比一下字符的顺序是否和people中字符出现的顺序一致，如果不一致直接返回false
// people => true
// ppeople-> true

package main

func main() {

}

var people = "people"

func input(str string) bool {
	if len(str) < len(people) { // 如果长度不够，肯定是不够
		return false
	}
	phash {
		p:
	}
	equalIndex := make([]int, 0, len(people))
	// step1
	for _, s := range str {
		equal := false
		for i, p := range people { // 如果p里没有s可以加，如果s没有p要直接返回
			if p == s && !equal { // 一个字符只判断一次，防止index重复写
				equal = true
				equalIndex = append(equalIndex, i)
			}
		}
		if !equal { // 如果s没有p要直接返回
			return false
		}
	}

	// step2

}
