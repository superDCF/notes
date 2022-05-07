package main

import (
	"log"
)

/*
'M' 代表一个 未挖出的 地雷，
'E' 代表一个 未挖出的 空方块，
'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的 已挖出的 空白方块，
数字（'1' 到 '8'）表示有多少地雷与这块 已挖出的 方块相邻，
'X' 则表示一个 已挖出的 地雷。

根据给定的扫雷面板，返回计算后的面板
1. 面板上方格的数字表示方格周围的雷数。
3. 如果点击到雷，游戏就结束了- 把该点击位改为 'X'
4. 如果一个 没有相邻地雷 的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的 未挖出 方块都应该被递归地揭露，如果其相邻的方块周围有雷，则相邻的方块停止递归揭露。
5. 如果一个 有相邻地雷 的空方块被挖出，则修改其数字，不递归挖

如果一个点击到一个未被挖出的方块，如果该方块是雷，按规则3，如果不是雷，就要判断其周围雷数，如果雷数是0，就要执行规则4。
*/

var (
	x = []int{-1, 0, 1}
	y = []int{-1, 0, 1}
)

func main() {
	log.Println("B=", 'B', "E=", 'E', "M=", 'M', "X=", 'X')
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}
	click := []int{3, 0}

	// board := [][]byte{
	// 	{'B', 1, 'E', 1, 'B'},
	// 	{'B', 1, 'M', 1, 'B'},
	// 	{'B', 1, 1, 1, 'B'},
	// 	{'B', 'B', 'B', 'B', 'B'}}
	// click := []int{1, 2}
	// log.Println(board)
	log.Println(updateBoard(board, click))
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(click) != 2 {
		return board
	}
	switch board[click[0]][click[1]] {
	case 'E':
		dfs(board, click[0], click[1])
	case 'M': // 如果点击到雷，游戏就结束了- 把该点击位改为 'X'
		board[click[0]][click[1]] = 'X'
	}
	return board
}

func dfs(board [][]byte, i, j int) {
	if board[i][j] != 'E' { //
		return
	}
	m := len(board)
	n := len(board[0])
	cnt := byte(0) // 方格i,j周围雷的数量
	for _, vx := range x {
		for _, vy := range y {
			if vx == 0 && vx == vy {
				continue
			}
			if i+vy < 0 || i+vy > m-1 {
				continue
			}
			if j+vx < 0 || j+vx > n-1 {
				continue
			}
			if board[i+vy][j+vx] == 'M' {
				cnt++
			}
			// else if board[i+vy][j+vx] == 'E' {
			// 	dfs(board, i+vy, j+vx)
			// }
		}
	}
	if cnt == 0 {
		board[i][j] = 'B'
		for _, vx := range x {
			for _, vy := range y {
				if vx == 0 && vx == vy {
					continue
				}
				if i+vy < 0 || i+vy > m-1 {
					continue
				}
				if j+vx < 0 || j+vx > n-1 {
					continue
				}
				if board[i+vy][j+vx] == 'E' {
					dfs(board, i+vy, j+vx)
				}
			}
		}
	} else {
		board[i][j] = cnt + 48 // 48是0的ascii码
	}
}
