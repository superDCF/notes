package main

import "log"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/

func main() {
	log.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	log.Println(minSubArrayLen(2, []int{1, 4, 4}))
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum := nums[0]
	minLen := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		for end < len(nums)-1 {
			if sum < target {
				end++
				sum += nums[end] // 这种方式决定end的边界条件
			} else {
				break
			}
		}
		if minLen == 0 && sum >= target {
			minLen = end - i + 1
		} else if end-i+1 < minLen && sum >= target {
			minLen = end - i + 1
		} else if sum < target {
			return minLen
		}
		// log.Printf("i=%v end=%v sum=%v  minLen=%v", i, end, sum, minLen)
		sum -= nums[i]
	}
	return minLen
}
