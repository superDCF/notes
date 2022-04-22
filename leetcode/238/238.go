package main

import "log"

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请不要使用除法，且在 O(n) 时间复杂度内完成此题。
输入: nums = [1,2,3,4]
输出: [24,12,8,6]
*/

func main() {
	log.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	val := 1
	zeroIndex1 := -1
	zeroIndex2 := -1
	for i, v := range nums {
		if v == 0 {
			if zeroIndex1 == -1 {
				zeroIndex1 = i
			} else {
				zeroIndex2 = i
			}
			continue
		}
		val *= v
	}
	ret := make([]int, len(nums))
	for i := range ret {
		if zeroIndex1 > -1 && zeroIndex2 > -1 || (zeroIndex1 > -1 && zeroIndex1 != i) {
			ret[i] = 0
		} else if zeroIndex1 > -1 && zeroIndex1 == i {
			ret[i] = val
		} else if zeroIndex1 != i {
			ret[i] = val / nums[i]
		}
	}
	return ret
}
