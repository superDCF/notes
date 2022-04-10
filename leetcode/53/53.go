/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

[-2,1,-3,4,-1,2,1,-5,4] => [4,-1,2,1] => 6

[-2,1,-3,4,-1,2,1,-1,4] => [4,-1,2,1,-1,4] = > 9

思路：找连续子数组最大值，可以类比为找数组中某个元素的最大值。
*/

package main

import "log"

func main() {
	// log.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -1, 4}))
	log.Println(maxSubArray([]int{-2, -1, -10}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	res := sum
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		res = max(sum, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
