package main

import "log"

func main() {
	// longestPalindrome("ada")
	// longestPalindrome("bda")
	log.Println(longestPalindrome("ababa"))
	log.Println(povit("ababa"))
}

/*
后面的可以依赖前面的判断，可以考虑动态规划
动态规划
ababa
	j	0	1	2	3	4
i
0		1	0	1	0	1
1			1	0	1	0
2				1	0	1
3					1	0
4						1

1. dp[i][j]：确定dp数组及i，j的表达，判断i到j区间是否是回文子字符串
2. 确定递推公式
2.1 如果s[i]!=s[j]，那么dp[i][j]=false
2.2 如果s[i]==s[j]，又分为i=j,j-i=1两种情况
2.3 而s[i]==s[j]后，可以由dp[i+1][j-1]判断内部是否相等
3. dp二维数组如何初始化
4. 确定遍历顺序，因为是依赖dp[i+1][j-1]，是先依赖内部的，所以二维表也是从下往上，从左往右，

*/
func longestPalindrome(s string) string {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ { // 初始化二维数组
		dp[i] = make([]int, len(s))
	}
	log.Println(dp)
	maxLen := 0
	left, right := 0, 0                // 表示至少有一个长度为1的
	for i := len(s) - 1; i >= 0; i-- { // 遍历二维数组
		for j := i; j < len(s); j++ {
			if s[i] != s[j] {
				dp[i][j] = 0
			} else {
				if i == j {
					dp[i][j] = 1
				} else if dp[i+1][j-1] == 1 {
					dp[i][j] = 1
				}
				if dp[i][j] == 1 && j-i+1 > maxLen { // 因为这里每次都是计算当前区间的i,j，所以可以直接计算maxLen
					maxLen = j - i + 1
					left, right = i, j
				}
			}
		}
	}
	log.Println(dp)

	// 从二维数组中找最大长度，优先从二维数组右上角找
	// for i := 0; i < len(s); i++ {
	// 	for j := len(s) - 1; j >= 0; j-- {
	// 		if dp[i][j] == 1 && j-i+1 > maxLen {
	// 			maxLen = j - i + 1
	// 			left, right = i, j
	// 		}
	// 	}
	// }
	log.Println(left, right)
	return s[left : right+1]
}

// 中心外扩法，找到一个中心点，向外扩散，判断前后指针是否相同
func povit(s string) string {
	maxLen, left, right := 0, 0, 0
	for i := 0; i < len(s)-1; i++ {
		m1, l1, r1 := extend(s, i, i)
		m2, l2, r2 := extend(s, i, i+1)
		if m1 > maxLen {
			maxLen, left, right = m1, l1, r1
		}
		if m2 > maxLen {
			maxLen, left, right = m2, l2, r2
		}
	}
	return s[left : right+1]
}

// 外扩，可以是奇数外扩，也可以是两个相同值外扩，、
// 这里有一个优化方法，当找到最大长度为s的一半或者以上，就可以自动排除剩下未遍历的
func extend(s string, i, j int) (int, int, int) {
	maxLen, left, right := 0, 0, 0
	for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
		if j-i+1 > maxLen { // 巧妙点：每次相等的时候，都计算maxLen，每次计算虽然有点浪费，但是可以排除很多边界问题
			maxLen = j - i + 1
			left, right = i, j
		}
		i--
		j++
	}
	return maxLen, left, right
}
