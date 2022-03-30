/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。
*/

package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(0b0110 &^ 0b1011)
	// log.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	// log.Println(threeSumClosest([]int{-4, -2, 1, 2, 5}, 1))
	// log.Println(threeSumClosest([]int{0, 0, 0}, 1))
	log.Println(threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
	log.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, -1))
	log.Println(threeSumClosest([]int{-3, 0, 1, 2}, 1))
	log.Println(threeSumClosest([]int{-1, -1, 1, 1, 3}, -1))
}

// 所有数字都要遍历到
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	var sum = nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			_sum := nums[i] + nums[left] + nums[right] // 改动某个参数，是否影响其他参数
			if abs(_sum-target) < abs(sum-target) {
				sum = _sum
			}
			if _sum > target {
				right--
			} else if _sum < target {
				left++
			} else {
				return _sum
			}
		}
	}
	return sum
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
