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
	"time"
)

func main() {
	log.Println(placeMine(5, 4, 18))
}

const (
	mineFlag = -1 // 雷标识
)

// placeMine 布雷
func placeMine(m, n, k int) [][]int {
	if m <= 0 || n <= 0 || m*n < k {
		return nil
	}
	// 生成雷方格
	mines := make([][]int, m)
	for i := 0; i < m; i++ {
		mines[i] = make([]int, n)
	}
	if k <= 0 {
		return mines
	}
	// 在方格上公平放置雷
	bits := make([]int, m*n)
	for i := 0; i < k; i++ {
		bits[i] = mineFlag
	}
	// 把雷随机分布
	for i := 0; i < k; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		index := rand.Intn(m * n)
		log.Println("index", index)
		bits[i], bits[index] = bits[index], bits[i]
	}
	log.Println(bits)

	for i, v := range bits {
		if v == mineFlag {
			row := i / n
			coulmn := i - row*n
			mines[row][coulmn] = mineFlag
		}
	}

	// 计算雷周围小方格的雷数
	x := []int{-1, 0, 1} // x轴方向，代表以方格为原点，如果原点是雷，则把其周围上下左右斜8个格子数加1
	y := []int{-1, 0, 1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mines[i][j] == mineFlag {
				for _, vx := range x {
					for _, vy := range y {
						if vx == 0 && vx == vy { // 表示原点
							continue
						}
						if i+vy < 0 || i+vy > m-1 { // 超出边界
							continue
						}
						if j+vx < 0 || j+vx > n-1 {
							continue
						}
						if mines[i+vy][j+vx] != mineFlag { // 如果该方格不是雷，才加
							mines[i+vy][j+vx]++
						}
					}
				}
			}
		}
	}
	return mines
}
