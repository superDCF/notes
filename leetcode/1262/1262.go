package main

import (
	"log"
	"math"
)

/*
给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。
输入：nums = [3,6,5,1,8]
输出：18
解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。
*/

func main() {

	// log.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	log.Println(maxSumDivThree([]int{2, 19, 6, 16, 5, 10, 7, 4, 11, 6}))
}

func maxSumDivThree(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, numsLen)
	}
	dp[0][0] = 0
	dp[1][0] = math.MinInt
	dp[2][0] = math.MinInt
	dp[nums[0]%3][0] = nums[0]
	for i := 1; i < numsLen; i++ {
		r := nums[i] % 3
		if r == 0 {
			dp[0][i] = max(dp[0][i-1], dp[0][i-1]+nums[i])
			dp[1][i] = max(dp[1][i-1], dp[1][i-1]+nums[i])
			dp[2][i] = max(dp[2][i-1], dp[2][i-1]+nums[i])
		} else if r == 1 {
			dp[0][i] = max(dp[0][i-1], dp[2][i-1]+nums[i])
			dp[1][i] = max(dp[1][i-1], dp[0][i-1]+nums[i])
			dp[2][i] = max(dp[2][i-1], dp[1][i-1]+nums[i])
		} else if r == 2 {
			dp[0][i] = max(dp[0][i-1], dp[1][i-1]+nums[i])
			dp[1][i] = max(dp[1][i-1], dp[2][i-1]+nums[i])
			dp[2][i] = max(dp[2][i-1], dp[0][i-1]+nums[i])
		}
	}
	if dp[0][numsLen-1] > math.MinInt {
		return dp[0][numsLen-1]
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
