package main

import "log"

// import "log"

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

*/

func main() {
	cur := []int{1, 2, 3}
	pre := cur[:]
	pre[0] = -1
	log.Println(cur, pre)
	log.Println(uniquePaths2(3, 7))
}

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
			}
		}
	}
	return matrix[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	pre := make([]int, n)
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = 1
		cur[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			log.Printf("j=%v cur[j-1]=%v pre[j]=%v", j, cur[j-1], pre[j])
			cur[j] = cur[j-1] + pre[j]
		}
		pre = cur[:]
	}
	return pre[n-1]
}

func uniquePaths1(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	matrix := make([][]int, m)
	val := 1
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			tmp[j] = val
			val++
		}
		matrix[i] = tmp
	}
	right := 0
	down := 0
	// path:=0
	var tree []int
	for right <= n && down <= m {
		tree = append(tree, matrix[right][down])
		down++
	}
	return 0
}
