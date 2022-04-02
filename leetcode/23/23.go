/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/

package main

import "log"

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	list4 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	l3 := mergeKLists([]*ListNode{list1, list2, list4})
	log.Printf("%+v", l3)
	for l3 != nil {
		log.Println(l3)
		l3 = l3.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	node := &ListNode{}
	list3 := node
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			node.Next = list2
			list2 = list2.Next
		} else {
			node.Next = list1
			list1 = list1.Next
		}
		node = node.Next
	}
	if list1 == nil {
		node.Next = list2
	} else if list2 == nil {
		node.Next = list1
	}
	return list3.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	listsLen := len(lists)
	if listsLen == 0 {
		return nil
	} else if listsLen == 1 {
		return lists[0]
	} else if listsLen == 2 {
		return mergeTwoLists(lists[0], lists[1])
	} else {
		index := len(lists) / 2
		l1 := mergeKLists(lists[0:index])
		l2 := mergeKLists(lists[index:])
		return mergeTwoLists(l1, l2)
	}
}
