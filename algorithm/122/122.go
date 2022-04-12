/*
给定一个数组 prices ，其中 prices[i] 表示股票第 i 天的价格。

在每一天，你可能会决定购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以购买它，然后在 同一天 出售。
返回 你能获得的 最大 利润 。
prices = [7,1,5,3,6,4] => 7
*/

package main

import "log"

func main() {
	log.Println(getFit([]int{7, 1, 5, 3, 7, 8}))
}

func getFit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxFit := 0
	buyPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > buyPrice {
			maxFit += prices[i] - buyPrice
			buyPrice = prices[i]
		} else {
			buyPrice = prices[i]
		}
	}
	return maxFit
}
