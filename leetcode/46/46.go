/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

package main

import "log"

func main() {
	log.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	numsLen := len(nums)
	var ret [][]int
	// used := make([]bool, numsLen)
	// backTrace(nums, nil, &ret, used, 0, numsLen)
	backTrace1(nums, &ret, 0, numsLen)
	return ret
}

func backTrace(nums, out []int, ret *[][]int, used []bool, deep, numsLen int) {
	if deep == numsLen {
		_out := append([]int{}, out...)
		*ret = append(*ret, _out)
		return
	}

	for i := 0; i < numsLen; i++ {
		if !used[i] {
			out = append(out, nums[i])
			used[i] = true
			// log.Printf("backTrace1 i=%v deep=%v out=%v used=%v", i, deep, out, used)
			backTrace(nums, out, ret, used, deep+1, numsLen)
			used[i] = false
			out = out[:len(out)-1]
			// log.Printf("backTrace2 i=%v deep=%v out=%v used=%v", i, deep, out, used)
		}
	}
}

func backTrace1(nums []int, ret *[][]int, deep, numsLen int) {
	if deep == numsLen {
		_out := append([]int{}, nums...)
		*ret = append(*ret, _out)
		return
	}
	for i := deep; i < numsLen; i++ {
		// if i != deep {
		log.Printf("backTrace1 deep=%v i=%v nums=%v", deep, i, nums)
		nums[i], nums[deep] = nums[deep], nums[i]
		// }
		backTrace1(nums, ret, deep+1, numsLen)
		// if i != deep {
		log.Printf("backTrace2 deep=%v i=%v nums=%v", deep, i, nums)
		nums[i], nums[deep] = nums[deep], nums[i]
		// }
	}
}
