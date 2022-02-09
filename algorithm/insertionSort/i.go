package main

import "log"

func main() {
	log.Println(InsertSort([]int{2, 4, 6, 5, 3, 1, 100}))
}

func InsertSort(arr []int) []int {
	arrL := len(arr)
	if arrL <= 1 {
		return arr
	}
	for i := 1; i < arrL; i++ {
		v := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > v {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = v
	}
	return arr
}
