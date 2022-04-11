/* 双指针法：原地移除数组中给予的某一个元素
[0,2,4,3,6,2,5] 移除2，在O(1)空间复杂度，O(n)时间复杂度下操作
O1空间，说明不需要再开辟新数组
On时间，说明至多只有一个for循环
快慢指针法：一开始，快慢指针一起移动，当找到一个符合的元素，慢指针停下，快指针继续for循环往下走一位，然后把慢指针所在的索引位用快指针所在的索引位值替换掉，以此往后循环
不和fast相等，就移动fast
*/

package main

import "log"

func main() {
	log.Println(remove([]int{0, 1, 2, 2, 2, 4, 5}, 2))
	log.Println(remove([]int{0, 1, 2, 3, 2, 4, 5}, 2))
	log.Println(remove([]int{2, 1, 2, 3, 4, 4, 5}, 4))
}

func remove(arr []int, val int) []int {
	if len(arr) == 0 {
		return nil
	}
	slow := 0 // 快慢指针法：当找到目标值之后，慢指针不往后移动，快指针往后移动，
	for fast := 0; fast < len(arr); fast++ {
		if arr[fast] != val {
			arr[slow] = arr[fast]
			slow++
		}
	}
	return arr[:slow]
}
