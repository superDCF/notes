package main

import "log"

func main() {
	arr0 := []int{}
	arr1 := []int{1}
	arr2 := []int{1, 2, 3, 4, 5, 6}
	arr3 := []int{1, 2, 6, 4, 5}
	arr4 := []int{1, 2, 3, 4, 5, 6, 100}
	i0 := BinarySearch(arr0, 6)
	i1 := BinarySearch(arr1, 6)
	i2 := BinarySearch(arr2, 6)
	i3 := BinarySearch(arr3, 6)
	i4 := BinarySearch(arr4, 1)
	log.Println(i0)
	log.Println(i1)
	log.Println(i2)
	log.Println(i3)
	log.Println(i4)
}

/*
n/2 * n/2^2 * n/2^3 .... n/2^k
n/2^k = 1
O(logn)
*/

func BinarySearch(arr []int, target int) (index int) {
	arrLen := len(arr)
	left := 0
	right := arrLen - 1
	for left <= right {
		index = (left + right) >> 1
		if target == arr[index] {
			return index
		} else if target > arr[index] {
			left = index + 1
		} else {
			right = index - 1
		}
	}
	return -1
}
