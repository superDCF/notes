package main

import (
	"log"
	"math"
)

/*
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

思路：
1. 每一位都有一个最大盛水
*/

func main() {
	log.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// dp
func maxArea2(height []int) int {
	if len(height) < 2 {
		return 0
	}

	max := 0
	i := 0
	j := len(height) - 1
	for i != j {
		val, isFirst := min(height[i], height[j])
		m := (j - i) * val
		if m > max {
			max = m
		}
		if isFirst {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(v1, v2 int) (val int, isFirst bool) {
	if v1 > v2 {
		return v2, false
	}
	return v1, true
}

// 暴力求解，dp[i][j]表示第i到第j之间的盛水量
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			m := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if m > max {
				max = m
			}
		}
	}

	return max
}
