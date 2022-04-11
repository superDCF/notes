/*
33. 搜索旋转排序数组
*/

package main

import "log"

func main() {
	log.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	// log.Println([]int{4, 5, 6, 7, 0, 1, 2}[0:0])
}

func search(nums []int, target int) int {
	for i, v := range nums {
		if target == v {
			return i
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	k := 0
	for i, v := range nums {
		if v < nums[k] {
			k = i
			break
		}
	}
	index := bSearch(nums, 0, k-1, target)
	if index == -1 {
		return bSearch(nums, k, len(nums)-1, target)
	}
	return index
}

func bSearch(nums []int, l, r, target int) int {
	if len(nums) == 0 || l > r {
		return -1
	}
	log.Printf("l=%v r=%v nums=%v", l, r, nums)
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return bSearch(nums, l, mid-1, target)
	}
	return bSearch(nums, mid+1, r, target)
}
