package main

import "log"

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

func main() {
	m1 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	log.Println(spiralOrder(m1))
	m2 := [][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15}}
	log.Println(spiralOrder(m2))
	m3 := [][]int{{1, 2}, {3, 4}}
	log.Println(spiralOrder(m3))
	m4 := [][]int{{1}, {3}}
	log.Println(spiralOrder(m4))
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	var out []int
	left := 0
	right := n - 1
	up := 0
	bottom := m - 1
	for left <= right && up <= bottom {
		log.Println(left, right, up, bottom)
		for i := left; i <= right; i++ {
			log.Println("up", matrix[left][i])
			out = append(out, matrix[left][i])
		}
		for i := up + 1; i <= bottom; i++ {
			log.Println("right", matrix[i][right])
			out = append(out, matrix[i][right])
		}
		if left < right && up < bottom {
			for i := right - 1; i >= left; i-- {
				log.Println("bottom", matrix[bottom][i])
				out = append(out, matrix[bottom][i])
			}
			for i := bottom - 1; i >= up+1; i-- {
				log.Println("left", matrix[i][left])
				out = append(out, matrix[i][left])
			}
		}

		left++
		right--
		up++
		bottom--
	}
	return out
}

func spiralOrder1(matrix [][]int) []int {
	var out []int
	cycle(matrix, &out, 0)
	return out
}

func cycle(matrix [][]int, out *[]int, deep int) {
	matrixLen := len(matrix)
	matrixLen0 := len(matrix[0])
	// if deep ==
	// n := 2*matrixLen0 + matrixLen - 2*(matrixLen-2)
	for i := deep; i < matrixLen0-deep; i++ {
		up := matrix[deep][i]
		log.Println("up", up)
		*out = append(*out, up)
	}
	isRight := false
	for i := deep + 1; i < matrixLen-(deep+1); i++ {
		right := matrix[i][matrixLen0-1-deep]
		log.Println("right", right)
		// isRight = true
		*out = append(*out, right)
	}

	for i := matrixLen0 - 1 - deep; i >= deep; i-- {
		if !isRight {
			// return
		}
		down := matrix[matrixLen-1-deep][i]
		log.Println("down", down)
		*out = append(*out, down)
	}
	for i := matrixLen - 1 - (deep + 1); i > deep; i-- {
		left := matrix[i][deep]
		log.Println("left", left)
		*out = append(*out, left)
	}
	cycle(matrix, out, deep+1)
}
