package main

import "log"

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

1. 把比较后，需要插入的值，从该位置以后，元素整体往后移动一位。有内存移动。
2. 从后往前比较，这样就会先把大的数，直接填充到最后位。


*/

func main() {
	merge2([]int{3, 4, 7, 0, 0, 0, 0}, 3, []int{1, 2, 3, 5}, 4)
	merge2([]int{1, 3, 5, 0, 0, 0}, 3, []int{2, 4, 6}, 3)
	merge2([]int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}, 5, []int{-1, -1, 0, 0, 1, 2}, 6)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n {
		return
	}
	first := 0
	second := 0
	for first < m+n && second < n {
		if first > m-1+second && nums1[first] == 0 {
			nums1[first] = nums2[second]
			first++
			second++
		} else if nums1[first] > nums2[second] {
			copy(nums1[first+1:], nums1[first:])
			nums1[first] = nums2[second]
			first++
			second++
		} else {
			first++
		}
		log.Println(nums1, first, second)
	}

	log.Println(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	first := m - 1
	second := n - 1
	thrid := m + n - 1
	for second >= 0 && first >= 0 {
		if nums1[first] > nums2[second] {
			nums1[thrid] = nums1[first]
			first--
			thrid--
		} else {
			nums1[thrid] = nums2[second]
			second--
			thrid--
		}
	}

	if first < 0 {
		copy(nums1[0:second+1], nums2[0:second+1])
	}
	log.Println(nums1)
}
