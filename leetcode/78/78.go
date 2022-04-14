package main

import "log"

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

回溯算法：
1. 首先要满足，解不重复
2. 该算法可以暴力得出所有解。
*/

func main() {
	log.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	out := []int{}
	ans := new([][]int)
	handler(nums, out, 0, ans)
	return *ans
}

// 1,2,3,4,5
func handler(nums, out []int, cur int, ans *[][]int) {
	// log.Println("handler", cur, out)
	if cur == len(nums) {
		tmp := append([]int{}, out...)
		*ans = append(*ans, tmp)
		return
	}
	out = append(out, nums[cur])
	handler(nums, out, cur+1, ans)
	out = out[:len(out)-1]
	handler(nums, out, cur+1, ans)
}
