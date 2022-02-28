/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
[7,1,5,3,6,4] => 5
*/
package main

import "log"

func main() {
	log.Println(getFit([]int{7, 2, 5, 1, 6, 4}))
	log.Println(getFit2([]int{7, 2, 5, 1, 6, 4}))
}

func getFit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxFit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxFit {
				maxFit = prices[j] - prices[i]
			}
		}
	}
	return maxFit
}

// 如果后面有比当前买入的价格还低的价格，就再计算买入价格，然后比较收益
func getFit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buyPrice := prices[0]
	maxFit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			if prices[i]-buyPrice > maxFit {
				maxFit = prices[i] - buyPrice
			}
		}
	}
	return maxFit
}