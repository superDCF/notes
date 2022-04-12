package main

import "log"

func main() {
	nums1 := []int{1, 3, 4, 9}
	nums2 := []int{2, 4, 6, 8, 10,12}
	arr := mergeArrays(nums1, nums2)
	log.Println(arr)
	log.Println(findMedianSortedArrays(nums1, nums2))
}

/*
	1,3,9
	2,4,6,8
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	arr := mergeArrays(nums1, nums2)
	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2
	} else {
		return float64(arr[len(arr)/2])
	}
}

func mergeArrays(nums1 []int, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	ret := make([]int, len1+len2)
	i1, i2 := 0, 0
	for i := 0; i < len1+len2; i++ {
		if i1 < len1 && i2 < len2 {
			if nums1[i1] < nums2[i2] {
				ret[i] = nums1[i1]
				i1++
			} else {
				ret[i] = nums2[i2]
				i2++
			}
		}
	}
	if i1 == len1 {
		copy(ret[i1+i2:], nums2[i2:])
	} else {
		copy(ret[i1+i2:], nums1[i1:])
	}
	return ret
}
