package main

import "log"

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

func main() {
	log.Println(majorityElement([]int{6, 5, 5}))
}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	val := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			val = v
			count = 1
			continue
		}
		if v == val {
			count++
		} else {
			count--
		}
	}
	return val
}
