/*
扫雷-布雷算法
方格 m x n，有k个雷
5*4 6

5*4 k=5
 1 -1  3  2
 1  2 -1 -1
 2  3  3  2
-1 -1  1  0
 2  2  1  0

 5*4 k=10
 -1 -1  3  2
 -1  4 -1 -1
  3  4  4 -1
 -1 -1  4  2
  3 -1 -1  1


  -1 -1 -1  2
  -1 -1  7 -1
  -1 -1 -1 -1
  -1 -1 -1 -1
  -1 -1 -1 -1
*/

package main

import (
	"log"
	"math/rand"
)

const mineFlag = -1 // 是雷的标识

func main() {
	log.Println(placeMine(5, 4, 6))
}

// placeMine 在m*n方格上，布置k个雷
func placeMine(m, n, k int) [][]int {
	if m <= 0 || n <= 0 || k <= 0 || k >= m*n {
		return nil
	}
	// step1:先生成m*n个方格，使用二维数组实现
	mines := make([][]int, m)
	for i := 0; i < m; i++ {
		mines[i] = make([]int, n)
	}

	// step2:把雷随机等概率的放入方格，
	genMines(mines, m, n, k)
	// step3: 计算方格周围雷数
	calcMine(mines, m, n)
	return mines
}

// genMines 随机生成雷
func genMines(mines [][]int, m, n, k int) {
	bits := make([]int, m*n)
	for i := 0; i < k; i++ {
		bits[i] = mineFlag
	}
	// 洗牌算法
	for i := m*n - 1; i >= 0; i-- {
		index := rand.Intn(i + 1) // 随机获取一个从0～i的值
		bits[i], bits[index] = bits[index], bits[i]
	}

	// 把bits上的雷映射到方格上
	for i := 0; i < m*n; i++ {
		if bits[i] == mineFlag {
			row := i / n
			column := i - row*n
			mines[row][column] = mineFlag
		}
	}
}

// calc 计算方格周围雷数
func calcMine(mines [][]int, m, n int) {
	x := []int{-1, 0, 1} // 计算方格周围上下左右斜的坐标表达，x轴
	y := []int{-1, 0, 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mines[i][j] == mineFlag { // 只对是雷的周围做处理
				for _, vx := range x {
					for _, vy := range y {
						if i+vy < 0 || i+vy > m-1 { // 在行的方向上超出方格面板的大小
							continue
						}
						if j+vx < 0 || j+vx > n-1 {
							continue
						}
						if vx == 0 && vx == vy { // 表示方格自己
							continue
						}
						if mines[i+vy][j+vx] != mineFlag {
							mines[i+vy][j+vx]++
						}
					}
				}
			}
		}
	}
}
