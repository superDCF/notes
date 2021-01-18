package main

import (
	"container/list"
	"log"
)

func main() {
	log.Println(SingleStack([]int{73, 74, 75, 71, 69, 72, 76, 76, 73}))
	log.Println(SingleStack2([]int{73, 74, 75, 71, 69, 72, 76, 76, 73}))
}

// 单调栈处理 Next greater number
/*
给一个数组 T = [73, 74, 75, 71, 69, 72, 76, 73]，这个数组存放的是近几天的天气气温（这气温是铁板烧？不是的，这里用的华氏度）。你返回一个数组，计算：对于每一天，你还要至少等多少天才能等到一个更暖和的气温；如果等不到那一天，填 0 。
举例：给你 T = [73, 74, 75, 71, 69, 72, 76, 73]，你返回 [1, 1, 4, 2, 1, 1, 0, 0]。
解释：第一天 73 华氏度，第二天 74 华氏度，比 73 大，所以对于第一天，只要等一天就能等到一个更暖和的气温，到75度那天，还要等4天才是76度。后面的同理。
*/

func SingleStack(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	stack := list.New()
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		interval := 0
		node := stack.Front()
		for node != nil {
			interval++
			if val := (node.Value).(int); val <= nums[i] {
				node = node.Next()
			} else {
				break
			}
		}
		if node == nil {
			interval = 0
		}
		stack.PushFront(nums[i])
		result[i] = interval
	}
	return result
}

// SingleStack2 stack store index
// [73, 74, 75, 71, 69, 72, 76, 76, 73]
func SingleStack2(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	stack := list.New()
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		var val int
		for stack.Len() != 0 {
			if val = (stack.Front().Value).(int); nums[val] <= nums[i] {
				stack.Remove(stack.Front())
			} else {
				break
			}
		}
		interval := 0
		if stack.Len() != 0 {
			interval = val - i
		}
		stack.PushFront(i)
		result[i] = interval
	}
	return result
}
