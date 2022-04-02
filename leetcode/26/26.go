/*
 删除有序数组中的重复项
*/

package main

import "log"

func main() {
	log.Println(removeDuplicates2([]int{-1, 0, 1, 1, 1, 2, 2}))
	log.Println(removeDuplicates2([]int{1, 1, 2}))
	log.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen < 2 {
		return numsLen
	}
	si := -1
	ei := -1
	count := 0
	l := 1
	r := numsLen
	for l < r {
		log.Printf("l=%v r=%v count=%v si=%v ei=%v numsLen=%v nums=%v", l, r, count, si, ei, numsLen, nums)
		if nums[l] == nums[l-1] {
			if si == -1 {
				si = l
			}
			ei = l
		} else {
			if ei != -1 {
				_count := ei - si + 1
				count += _count
				r = r - _count
				l = l - _count - 1
				nums = append(nums[0:si], nums[ei+1:]...)
			}
			si = -1
			ei = -1
		}
		l++
	}
	if ei != -1 {
		count += ei - si + 1
		nums = append(nums[0:si], nums[ei+1:]...)
	}
	return numsLen - count
}

func removeDuplicates2(nums []int) int {
	numsLen := len(nums)
	if numsLen < 2 {
		return numsLen
	}
	i := 0
	for j := 1; j < numsLen; j++ {
		if nums[i] != nums[j] {
			i++
			if j > i {
				nums[i] = nums[j]
			}
		}
	}
	return i + 1
}
