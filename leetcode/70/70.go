package main

import "log"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

func main() {
	log.Println(climbStairs(5))
}

func climbStairs2(n int) int {
	// dp[n] = dp[n-1]+dp[n-2]
	if n < 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	// log.Println(dp)
	return dp[n]
}

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	ppre := 0
	pre := 1
	ans := 0
	for i := 1; i <= n; i++ {
		ans = ppre + pre
		ppre = pre
		pre = ans
	}
	return ans
}
