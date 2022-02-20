package main

import (
	"fmt"
	"log"
	"math"
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
}

// Add 添加前闭后开区间元素
func (rangeList *RangeList) Add(rangeElement [2]int) error {
	if err := rangeList.checkArgs(rangeElement); err != nil {
		return err
	}
	// 当为空和列表最后的范围小于输入的区间的第一个值，则直接append list最后
	if len(rangeList.list) == 0 || rangeList.list[len(rangeList.list)-1][1] < rangeElement[0] {
		rangeList.list = append(rangeList.list, rangeElement)
		return nil
	} else if rangeList.list[0][0] > rangeElement[1] { // 当列表的第一个区间范围的第一个元素比输入的范围的第二个元素大，就把输入的区间放到列表的最前面
		rangeList.list = append([][2]int{rangeElement}, rangeList.list...)
		return nil
	} else { // 然后就是交叉取或集和两区间范围内的并集
		index := -1
		for i := 0; i < len(rangeList.list); i++ {
			if rangeElement[0] > rangeList.list[i][0] {
				index++
				break
			}
		}
		if index == -1 {
			
		}
		rangeList.list = append(rangeList.list[0:index+1], rangeList.list[index:]...)
	}

	return nil
}

// Remove 移除前闭后开区间元素
func (rangeList *RangeList) Remove(rangeElement [2]int) error {

	return nil
}

// Print 打印当前rangeList的有效区间索引
func (rangeList *RangeList) Print() error {

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
