/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

1,2,3,4,5,6,7,8,9

1,2,3
8,9,4
7,6,5
*/
package main

import "log"

func main() {
	log.Println(generateMatrix(1))
}

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	left := 0
	right := n - 1
	up := 0
	down := n - 1
	first := 1
	for left <= right && up <= down {
		for i := left; i <= right; i++ {
			matrix[left][i] = first
			first++
		}
		for i := up + 1; i <= down; i++ {
			matrix[i][right] = first
			first++
		}

		for i := right - 1; i >= left; i-- {
			matrix[down][i] = first
			first++
		}
		for i := down - 1; i >= up+1; i-- {
			matrix[i][left] = first
			first++
		}
		left++
		right--
		up++
		down--
	}
	return matrix
}
