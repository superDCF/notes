package main

import (
	"fmt"
	"log"
)

/*
时间复杂度：O(n) O(n^2) O(n^2)
原地排序算法，O(1)
稳定
 */

func main() {
	fmt.Println(BubbleSort([]int{2, 4, 6, 5, 3, 1, 100}))
}

func BubbleSort(src []int) (des []int) {
	for i := 0; i < len(src)-1; i++ {
		for j := 0; j < len(src)-i-1; j++ {
			log.Println(j)
			if src[j] > src[j+1] {
				tmp := src[j]
				src[j] = src[j+1]
				src[j+1] = tmp
			}
		}
	}
	return src
}
