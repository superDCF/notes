package main

import "log"

func main() {
	findMedianSortedArrays([]int{2, 4, 6, 8, 10, 12}, []int{3, 6, 9, 12, 15})
}

func findMedianSortedArrays(nums1 []int, nums2 []int) {
	numsAll := make([]int, 0, len(nums1)+len(nums2))
	isMore := len(nums1) >= len(nums2)
	for i := 0; i < len(nums1); i++ {
		if isMore {
			if len(nums2) > i {
				if nums1[i] > nums2[i] {
					numsAll = append(numsAll, nums2[i])
				} else {
					numsAll = append(numsAll, nums1[i])
				}
			} else {
				numsAll = append(numsAll, nums1[i:]...)
			}
		} else {
			if i != len(nums1)-1 {
				if nums1[i] > nums2[i] {
					numsAll = append(numsAll, nums2[i])
				} else {
					numsAll = append(numsAll, nums1[i])
				}
			} else {
				numsAll = append(numsAll, nums1[2:]...)
			}
		}
	}
	log.Println(numsAll)
}
