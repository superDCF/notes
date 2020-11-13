package merge

// var Prev = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172, 174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196, 198}
// var Next = []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99, 102, 105, 108, 111, 114, 117, 120, 123, 126, 129, 132, 135, 138, 141, 144, 147, 150, 153, 156, 159, 162, 165, 168, 171, 174, 177, 180, 183, 186, 189, 192, 195, 198, 201, 204, 207, 210, 213, 216, 219, 222, 225, 228, 231, 234, 237, 240, 243, 246, 249, 252, 255, 258, 261, 264, 267, 270, 273, 276, 279, 282, 285, 288, 291, 294, 297}
var Prev = GenSeqNum(10000, 3)
var Next = GenSeqNum(10000, 2)

func GenSeqNum(n, factor int) []int {
	container := make([]int, n)
	for i := 1; i < n; i++ {
		container[i] = i * factor
	}
	return container
}

// 哨兵
func Merge(prev, next []int) []int {
	pL := len(prev)
	nL := len(next)
	i, p, n := 0, 0, 0
	container := make([]int, pL+nL)

	if prev[pL-1] > next[nL-1] {
		for n < nL {
			minV := prev[p]
			if minV > next[n] {
				minV = next[n]
				n++
			} else {
				p++
			}
			container[i] = minV
			i++
		}
		for i < pL+nL {
			container[i] = prev[p]
			p++
			i++
		}
	} else {
		for p < pL {
			minV := prev[p]
			if minV > next[n] {
				minV = next[n]
				n++
			} else {
				p++
			}
			container[i] = minV
			i++
		}
	}
	return container
}

// common
func Merge2(prev, next []int) []int {
	pL := len(prev)
	nL := len(next)
	container := make([]int, pL+nL)
	i, p, n := 0, 0, 0
	for p < pL && n < nL {
		minV := prev[p]
		if minV > next[n] {
			minV = next[n]
			n++
		} else {
			p++
		}
		container[i] = minV
		i++
	}
	if p == pL {
		for i < pL+nL {
			container[i] = next[n]
			n++
			i++
		}

	} else {
		for i < pL+nL {
			container[i] = prev[p]
			p++
			i++
		}
	}

	return container
}

// func GenRandomNum()
