package main

import (
	"log"
	// "math"
)

/*
最低票价

1,4,6,7,8,20
30
7+1+1
6*1
1+7+1

1,2,3,4,5,6,7,8,9,10,30,31
12*1
30+1
7+7+1+1

[1,2,3,4,6,8,9,10,13,14,16,17,19,21,24,26,27,28,29]
[3,14,50]


1. 如果当天需要出行，则有以下3种购票策略
	1. 只买1天，不包1天以后的票
	2. 买7天，不包7天以后的票
	3. 买30天，不包30天以后的票
2. 当天和之后几天是否需要买票，依赖后几天是否需要出行，对后面的依赖，可以考虑从后往前计算。所以有以下推导公式：
dp[i] =	{
	min(dp[i+1]+c[0],dp[i+7]+c[1],d[i+30]+c[2]) // 表示当天需要出行
	dp[i+1] // 表示当天不需要出行
}
dp[i]表示从第i天到最后的最低花费。

3. dp[days[len(days)-1]] = min(c[0],c[1],c[2]) // len(days)-1 表示最后一天出行

dp[20] = 1
dp[8] =  min(dp[8+1]+c[0],dp[8+7]+c[1],d[8+30]+c[2]) => min(dp[9]+c[0],dp[15]+c[1],d[38]+c[2])
dp[9] = min(dp[10]+c[0],dp[16]+c[1],d[39]+c[2])
	.
	.
	.
dp[19] = min(dp[19+1]+c[0],dp[19+7]+c[1],d[19+30]+c[2]) => min(dp[20]+c[0],dp[26]+c[1],d[39]+c[2])
dp[20] = min(dp[20+1]+c[0],dp[20+7]+c[1],d[20+30]+c[2]) => min(dp[21]+c[0],dp[27]+c[1],d[50]+c[2]) =>  min(c[0],c[1],c[2]) 假设c[0]最小
dp[i > 20] = 0
记忆化日期花费
daysConst = {
	>0: 0
	20: c[0]
	19: min(dp[20],dp[26],dp[49]) => min(dp[20],dp[26],dp[49]) // 第19天不需要买
	  .
	  .
	15: min(dp[16]+c[0],dp[22]+c[1],d[45]+c[2]) => min(dp[16]+c[0],0+c[1],0+c[2])
	  .
	8 : min(dp[9]+c[0],dp[15]+c[1],d[38]+c[2]) => min(dp[9]+c[0],dp[15]+c[1],d[38]+c[2])
}

dp[24] = min(min(dp[25]+costs[0], dp[31]+costs[1]), dp[54]+costs[2])
*/

func main() {
	// log.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	log.Println(mincostTickets2([]int{3, 4, 6, 8, 9, 10, 13, 14, 16, 17, 19, 21, 24, 26, 27, 28, 29}, []int{3, 14, 50}))
}
func mincostTickets(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}
	daysM := map[int]bool{}
	for _, v := range days {
		daysM[v] = true
	}

	dp := make([]int, days[len(days)-1]+30+1)
	dp[days[len(days)-1]] = min(min(costs[0], costs[1]), costs[2])
	for i := days[len(days)-1] + 1; i < len(dp); i++ {
		dp[i] = 0
	}
	for i := days[len(days)-1] - 1; i >= 0; i-- {
		if daysM[i] {
			dp[i] = min(min(dp[i+1]+costs[0], dp[i+7]+costs[1]), dp[i+30]+costs[2])
		} else {
			dp[i] = dp[i+1]
		}
	}
	for i, v := range dp {
		log.Println(i, v)
	}
	return dp[1]
}

func mincostTickets2(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}

	dp := make([]int, days[len(days)-1]+30+1)
	for i := days[len(days)-1] + 1; i < len(dp); i++ {
		dp[i] = 0
	}
	for i, j := days[len(days)-1], len(days)-1; i >= 0 && j >= 0; i-- {
		if i == days[j] {
			dp[i] = min(min(dp[i+1]+costs[0], dp[i+7]+costs[1]), dp[i+30]+costs[2])
			j--
		} else {
			dp[i] = dp[i+1]
		}
	}
	for i, v := range dp {
		log.Println(i, v)
	}
	return dp[days[0]]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
