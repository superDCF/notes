package main

import (
	"fmt"
	"log"

	. "./pkg"
)

/*
归并排序
就是把一个数组分成两段，再对两段各自排序

*/

func main() {
	// log.Println(MergeSort([]int{2, 4, 6, 5, 3, 1, 100}))

	log.Println(Merge([]int{2, 4, 5}, []int{2, 4, 6, 100}))
	log.Println(Merge([]int{2, 4, 6}, []int{1, 3, 5, 100}))
	log.Println(Merge2([]int{2, 4, 5}, []int{1, 3, 6, 100}))
	log.Println(Merge2([]int{2, 4, 6}, []int{1, 3, 5, 100}))
	log.Println(Merge([]int{2, 4, 5, 102, 104}, []int{1, 3, 6, 100}))
	log.Println(Merge([]int{2, 4, 6, 102, 104}, []int{1, 3, 5, 100}))
	log.Println(Merge2([]int{2, 4, 5, 102, 104}, []int{1, 3, 6, 100}))
	log.Println(Merge2([]int{2, 4, 6, 102, 104}, []int{1, 3, 5, 100}))

	// log.Println(Merge(Prev, Next))
	// log.Println(len(Merge(Prev, Next)))
	// log.Println(Merge(Next, Prev))
	// log.Println(Merge2(Prev, Next))
	// log.Println(Merge2(Next, Prev))
	// log.Println(fmt.Sprintf("%v", []int{2, 4, 6, 102, 104}) == fmt.Sprintf("%v", []int{2, 4, 6, 102, 104}))
	log.Println(fmt.Sprintf("%v", Merge(Prev, Next)) == fmt.Sprintf("%v", Merge(Next, Prev)))
	log.Println(fmt.Sprintf("%v", Merge2(Prev, Next)) == fmt.Sprintf("%v", Merge2(Next, Prev)))
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// m := len(arr) / 2
	return arr
}

func ESort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	m := len(arr) / 2
	return Merge(ESort(arr[:m]), ESort(arr[m:]))
}
