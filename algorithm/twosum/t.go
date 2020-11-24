package main

import (
	"log"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	go func() {
		for {
			twoSum([]int{2, 7, 11, 55}, 9)
		}
	}()

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Minute)
}




func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	viewMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if index, ok := viewMap[target-v]; ok {
			return []int{index, i}
		}
		viewMap[v] = i
	}
	return nil
}
