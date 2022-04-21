package main

import (
	"log"
)

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
[3,2,1,5,6,4]
[5,6,4,3,2,1]
[6,5,4,3,2,1]

[3,2,3,1,2,4,5,5,6]
[2,4,5,5,6,3,2,3,1]
[6,5,5,4,2,3,2,3,1]
*/

func main() {
	// nums := []int{7, 6, 5, 4, 3, 2, 1}
	// findKthLargest(nums, 5)
	// log.Printf("1 %v %p %p", nums, nums, &nums, nums[5-1])
	nums := []int{3, 2, 1, 5, 6, 4}
	quickSort3(nums, 0, len(nums)-1)
	log.Println(nums)
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	quickSortTopK(nums, k-1, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort2(nums []int, k, l, r int) {
	log.Printf("2 %v %p %p", nums, nums, &nums)
	nums = append(nums, k, l, r)
	// nums[l], nums[r] = nums[r], nums[l]
	log.Printf("3 %v %p %p", nums, nums, &nums)
}

func quickSortTopK(nums []int, k, l, r int) {
	log.Println(nums, l, r)
	if l >= r {
		return
	}
	left := l
	right := r
	m := (l + r) / 2
	povite := nums[m]
	for left < right {
		for left < right && nums[right] < povite {
			right--
		}
		if left <= right {
			nums[m] = nums[right]
		}
		for left < right && nums[left] >= povite {
			left++
		}
		if left < right {
			nums[right] = nums[left]
		}
		nums[left] = povite
		m = left
	}
	if left == k {
		return
	} else if left > k {
		quickSortTopK(nums, k, l, left-1)
	} else {
		quickSortTopK(nums, k, left+1, r)
	}
}

func quickSort(nums []int, l, r int) {
	log.Println(nums, l, r)
	if l >= r {
		return
	}
	left := l
	right := r
	m := (l + r) / 2
	povite := nums[m]
	for left < right {
		for left < right && nums[right] <= povite {
			right--
		}
		if left < right {
			nums[m] = nums[right]
		}
		for left < right && nums[left] >= povite {
			left++
		}
		if left < right {
			nums[right] = nums[left]
		}
		nums[left] = povite
		m = left
	}
	quickSort(nums, l, left-1)
	quickSort(nums, left+1, r)
}

func quickSort3(nums []int, l, r int) {
	if l >= r {
		return
	}
	left := l
	right := r
	m := (l + r) / 2
	povite := nums[m]
	for left < right {
		for nums[right] < povite {
			right--
		}
		for nums[left] > povite {
			left++
		}
		if left <= right {
			nums[right], nums[left] = nums[left], nums[right]
			left++
			right--
		}
	}
	// 移动过左右两边的值之后，就会确定较大的值的范围和较小的值的范围，然后递归分别再对这两个区间内部做交换。但是这种不适合找TopK，因为这里分区的位置可能不对
	quickSort(nums, l, right)
	quickSort(nums, left, r)
}

func quick_sort_new(nums []int, l, r int) {
	for l < r {
		x := l
		y := r
		z := nums[l+1]
		for x <= y {
			for nums[x] > z {
				x++
			}
			for nums[y] < z {
				y--
			}
			if x <= y {
				nums[x], nums[y] = nums[y], nums[x] // 这里交换完povit左右的值，其实并不满足左边的都比povit大，右边的都比povit小，但是把小值移动到后面，把大值移动到前面，然后也不是根据基准值分左右，而是根据分过的大小值分左右，再做排序
				x++
				y--
			}
		}
		quick_sort_new(nums, x, r)
		r = y
	}
}
