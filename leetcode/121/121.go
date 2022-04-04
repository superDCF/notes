package main

import "log"

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。


*/

func main() {
	log.Println(maxProfit([]int{1, 4, 3, 6, 7, 4, 5}))
	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Println(maxProfit([]int{5, 2, 1, 2, 1, 16, 0, 20}))
}

// maxProfit 只能买卖一次
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	minP := 0 // 最小值 // 直到遇到更小的值之前的最大值的利润比较，也就是找某范围的最大差值
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[minP] {
			minP = i
		} else if prices[i]-prices[minP] > profit {
			profit = prices[i] - prices[minP]
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// maxProfit2 可以买卖多次
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	cur := -1
	for i := 1; i < len(prices); i++ {
		if cur == -1 && prices[i-1] < prices[i] {
			cur = i - 1
		} else if cur > -1 && prices[i-1] > prices[i] {
			profit += prices[i-1] - prices[cur]
			cur = -1
		}
	}
	if cur > -1 {
		profit += prices[len(prices)-1] - prices[cur]
	}
	return profit
}
