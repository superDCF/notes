package main

import (
	"log"
	"math"
	"sort"
)

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

dp[0][0] = 0
dp[1][0] = -1/1
if amount-coins[i] == coins[i]{
	dp[amount][i] = dp[amount-coins[i]][i]+1
}
	0	1	2	3	4	5	6
0 1	0	1	2	3	4	5	6
1 2	0	1	1	2	2	3	3
2 5	0	1	1	2	2	1	2

	1	2	3	4	5	6
2	-1	1	-1	2	-1	3
5	-1	1	-1	2	1	3
*/

func main() {
	// log.Println(coinChange([]int{83, 186, 408, 419}, 6249))
	log.Println(coinChange([]int{2, 5}, 11))
	// 419*13 = 5447 6249-5447=802
	//
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	coinsLen := len(coins)
	if coinsLen == 0 {
		return -1
	}
	amount = amount + 1
	dp := make([]int, amount)
	for i := 0; i < amount; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := 1; i < amount; i++ {
		for j := 0; j < coinsLen; j++ {
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt {
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	if dp[amount-1] == math.MaxInt {
		return -1
	}
	return dp[amount-1]
}

func min(a, b int) int {

	if a > b {
		return b
	}
	return a
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	coinsLen := len(coins)
	if coinsLen == 0 {
		return -1
	}
	sort.Ints(coins)
	a0 := amount
	n := 0
	for j := 0; j < coinsLen; j++ {
		for i := coinsLen - 1 - j; i >= 0; i-- {
			log.Println(j, i, amount, coins[i], n)
			for amount > coins[i] {
				amount = amount - coins[i]
				n++
			}
			if amount == coins[i] {
				amount = 0
				n++
				return n
			}
		}
		n = 0
		amount = a0
	}
	return -1
}
