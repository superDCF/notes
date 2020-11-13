package main

import "log"

/*
快速排序原理：
1. 对一个给定的数组排序，选取一个pivot（分区点）；pivot的选取，非常影响排序的性能
2. 把大于pivot的元素放在数组的右边，把小于pivot的元素放在数组的左边
3. 递归处理由pivot分割的左右子数组，直到子数组的长度为1

phony:
	QuickSort(a){
		m := povit(a)
		sort(a,f,m-1)
		sort(a,m+1,l)
	}

	1,4,7,9,5,3,8

	1,4,7,5,3,8,9

	1,3,4,7,5

		4,5,7

*/

func main() {
	a := []int{6, 1, 4, 7, 9, 5, 3, 8}
	log.Println(QuickSort(a, 0, 7))
}

func QuickSort(a []int, left, right int) []int {
	log.Println("ZERO", left, right, a)
	if right > left {
		i := povit(a, left, right)
		QuickSort(a, left, i-1)
		QuickSort(a, i+1, right)
	}
	return a
}

// 在a的f和l之间，选出一个povit
func povit(a []int, left, right int) int {
	log.Println("first", left, right, a)
	tmp := a[left] // 基准数据
	for left < right {
		log.Println("second", left, right, a)
		for left < right && tmp <= a[right] {
			right--
		}
		a[left] = a[right]
		for left < right && a[left] <= tmp {
			left++
		}
		a[right] = a[left]
	}
	a[left] = tmp
	return left
}
