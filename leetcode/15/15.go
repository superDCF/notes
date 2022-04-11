/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组
*/

package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(threeSum3([]int{-1, 0, 1, 2, -1, -4}))
	log.Println(threeSum3([]int{-1, 0, 1, 2, 2, 3, -1, -4}))
	log.Println(threeSum3([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
	log.Println(threeSum3([]int{0, 0, 0, 0}))
	log.Println(threeSum3([]int{-2, 0, 1, 1, 2}))

}

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	var sums [][]int
	left := 0
	right := len(nums) - 1
	for left < right {
		log.Printf("left=%v right=%v sums=%v", left, right, sums)
		for i := left + 1; i < right; i++ {
			if nums[left]+nums[right]+nums[i] == 0 {
				sums = append(sums, []int{nums[left], nums[i], nums[right]})
			}
		}
		if nums[left]+nums[right] >= 0 {
			if nums[left]+nums[right]+nums[left] >= 0 {
				right--
			} else {
				left++
			}
		} else {
			if nums[left]+nums[right]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	if len(sums) == 0 {
		return nil
	}
	desc := [][]int{sums[0]}
	for i := 1; i < len(sums); i++ {
		isDiff := false
		for j := 0; j < len(sums[i]); j++ {
			if sums[i-1][j] != sums[i][j] {
				isDiff = true
				break
			}
		}
		if isDiff {
			desc = append(desc, sums[i])
		}
	}

	return desc
}

// 时间复杂度太高
func threeSum2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	var sums [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					sums = append(sums, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	log.Println("sums", sums)
	if len(sums) == 0 {
		return nil
	}
	desc := [][]int{sums[0]}
	for i := 1; i < len(sums); i++ {
		isDiff := false
		for j := 0; j < len(sums[i]); j++ {
			if desc[len(desc)-1][j] != sums[i][j] {
				isDiff = true
				break
			}
		}
		if isDiff {
			desc = append(desc, sums[i])
		}
	}
	return desc
}

func threeSum3(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	var sums [][]int
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right { // 其实就是两数之和，找第三个数
			if left > i+1 && nums[left] == nums[left-1] { // 防止left在遍历过程中遇到相同的值，只取第一个相同的值
				left++
				continue
			}
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				sums = append(sums, []int{nums[i], nums[left], nums[right]})
				left++ // 直接向后走，即使有重复的，可以通过i往后走，防止有遗漏
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}
	}
	return sums
}
