package main

import "log"

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

思路：
方法一：
1. 如果该方格是1，则把遍历的方格改成2，继续遍历该方格的上下左右方格，如果上下左右方格有存在为1的，继续遍历该格子的上下左右，直到遇到边界或者水
2. 如果递归相邻的方格都是水，则表明该方格及其相邻方格是岛屿

方法二：
由方法一，把遍历过的方格及其相邻方格是1的改为2，继续遍历网格，计算未标记为1的方格，即是岛屿数量。


*/

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '1'},
		{'1', '1', '0', '1', '1'},
		{'0', '0', '1', '1', '1'},
		{'0', '0', '0', '1', '1'},
	}
	log.Println(numIslands2(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				if dfs(grid, i, j) {
					total++
				}
			}
		}
	}
	return total
}

// 判断该位置上下左右是否都是水
func dfs(grid [][]byte, i, j int) bool {
	if !isArea(grid, i, j) {
		return true
	}
	if grid[i][j] != '1' {
		return true
	}

	grid[i][j] = '2'
	up := dfs(grid, i-1, j)
	down := dfs(grid, i+1, j)
	left := dfs(grid, i, j-1)
	right := dfs(grid, i, j+1)
	if up && down && left && right {
		return true
	}
	return false
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs2(grid, i, j)
				total++
			}
		}
	}
	return total
}

// 把该方格及其相邻方格是1的标记为2
func dfs2(grid [][]byte, i, j int) {
	if !isArea(grid, i, j) {
		return
	}
	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func isArea(grid [][]byte, i, j int) bool {
	if i < 0 || i > len(grid)-1 {
		return false
	}
	if j < 0 || j > len(grid[i])-1 {
		return false
	}
	return true
}
