package main

import (
	"fmt"
	"log"
	"sync"
)

/*
Author: Super
Date: 2022-01-18
// Task: Implement a struct named 'RangeList'
// A pair of integers define a range, for example: [1, 5). This range includes integers: 1, 2, 3, and 4. // A range list is an aggregate of these ranges: [1, 5), [10, 11), [100, 201)
// NOTE: Feel free to add any extra member variables/functions you like.

// 思路：
用有序数组存储区间的值
*/

// RangeList 底层是类bitmap的实现
type RangeList struct {
	list [][2]int
	sync.Map
}

// Add 添加前闭后开区间元素
func (rangeList *RangeList) Add(rangeElement [2]int) error {
	if err := rangeList.checkArgs(rangeElement); err != nil {
		return err
	}
	if len(rangeList.list) == 0 {
		rangeList.list = append(rangeList.list, value{val: rangeElement[0], typ: true}, value{val: rangeElement[1], typ: false})
	} else {
		var start, end = -1, -1
		for i, v := range rangeList.list {
			if rangeElement[0] < v.val {
				start = i
			}
			if rangeElement[1] < v.val {
				end = i
			}
		}
	}

	if rangeElement[1]+1 > len(rangeList.bits) {
		newBits := make([]uint8, rangeElement[1]-len(rangeList.bits)+1)
		rangeList.bits = append(rangeList.bits, newBits...)
	}
	// 为区间赋值
	for i := rangeElement[0]; i < rangeElement[1]; i++ {
		rangeList.bits[i] = 1
	}
	return nil
}

// Remove 移除前闭后开区间元素
func (rangeList *RangeList) Remove(rangeElement [2]int) error {
	if err := rangeList.checkRemoveArgs(rangeElement); err != nil {
		return err
	}
	for i := rangeElement[0]; i < rangeElement[1]; i++ {
		rangeList.bits[i] = 0
	}
	return nil
}

// Print 打印当前rangeList的有效区间索引
func (rangeList *RangeList) Print() error {
	indexs := make([][2]int, 0)
	start := 0
	for i, v := range rangeList.bits {
		if v == 1 && start == 0 {
			start = i
		} else if v == 0 && start != 0 {
			indexs = append(indexs, [2]int{start, i})
			start = 0
		}
	}
	result := ""
	for _, v := range indexs {
		result = fmt.Sprintf("%s [%d, %d)", result, v[0], v[1])
	}
	fmt.Printf("%s\n", result)
	return nil
}

// checkArgs 检查参数是否合法
func (rangeList *RangeList) checkArgs(rangeElement [2]int) error {
	if rangeElement[0] < 0 {
		return fmt.Errorf("checkArgs params must be unsigned integers! rangeElement=%+v", rangeElement)
	}
	if rangeElement[0] > rangeElement[1] {
		return fmt.Errorf("checkArgs params is invalid; rangeElement=%+v", rangeElement)
	}
	return nil
}

// checkRemoveArgs 检查remove函数参数是否合法
func (rangeList *RangeList) checkRemoveArgs(rangeElement [2]int) error {
	if err := rangeList.checkArgs(rangeElement); err != nil {
		return err
	}
	// 输入的范围，如果大于bits长度，
	if rangeElement[1] > len(rangeList.bits) || rangeElement[0] > len(rangeList.bits) {
		return fmt.Errorf("checkRemoveArgs params=%+v more than rangelist length=%d,", rangeElement, len(rangeList.bits))
	}
	return nil
}

func main() {
	rl := RangeList{}

	rl.Add([2]int{1, 5})
	rl.Print()
	// Should display: [1, 5)
	rl.Add([2]int{10, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 21})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{2, 4})
	rl.Print()
	// // Should display: [1, 5) [10, 21)
	rl.Add([2]int{3, 8})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 10})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 11})
	rl.Print()
	// Should display: [1, 8) [11, 21)
	rl.Remove([2]int{15, 17})
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)
	rl.Remove([2]int{3, 19})
	rl.Print()
	// Should display: [1, 3) [19, 21)
}
