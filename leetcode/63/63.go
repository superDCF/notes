package main

import "log"

// import "log"

/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

*/

func main() {

}
func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen < 2 {
		return 0
	}
	profit := 0
	minPriceIndex := 0
	for i := 1; i < pricesLen; i++ {
		if prices[minPriceIndex] > prices[i] {
			minPriceIndex = i
		}else if prices[i]-prices[minPriceIndex] > profit {
			profit = prices[i] - prices[minPriceIndex]
		}
	}
	return profit
}
