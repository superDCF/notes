/*
二分法
*/
package main

import "log"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 5}
	log.Println(binarySearch(arr, 3))
	log.Println(binarySearch(arr, 0))
	log.Println(binarySearch(arr, 6))
}

func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1
	index := -1
	for right >= left {
		// index, left, right = get(arr, left, right, target)
		povit := (left + right) / 2
		if arr[povit] == target {
			return povit
		} else if arr[povit] > target {
			right = povit - 1
		} else {
			left = povit + 1
		}
	}
	log.Println("aa", left, right)
	return index
}
