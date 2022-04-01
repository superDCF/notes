/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
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
	l3 := mergeTwoLists(list1, list2)
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
		node.Next = &ListNode{} // 把上一次的变量的next赋为下一个要生成的指针，相当于继续向后走
		if list1.Val > list2.Val {
			node.Next.Val = list2.Val
			list2 = list2.Next
		} else {
			node.Next.Val = list1.Val
			list1 = list1.Next
		}
		// log.Printf("node=%+v p=%p next=%v", node, node, node.Next)
		node = node.Next
	}
	if list1 == nil {
		node.Next = list2
	} else if list2 == nil {
		node.Next = list1
	}
	return list3.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
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
