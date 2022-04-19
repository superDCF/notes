package main

import "log"

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
在 O(n log n) 时间复杂度和常数级空间复杂度下
1. 归并排序
2. 常数空间复杂度，说明要自底向上归并
3. 从中分割链表，快慢指针，找到中间节点
*/

func main() {
	// l1 := &ListNode{2, &ListNode{4, &ListNode{5, nil}}}
	// l2 := &ListNode{1, &ListNode{3, &ListNode{6, nil}}}
	// head := merge(l1, l2)
	// head.str("hhh")
	l3 := &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}}
	sortList(l3).str("sort")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) str(s string) {
	l0 := l
	for l0 != nil {
		log.Printf("str%v=%v", s, l0.Val)
		l0 = l0.Next
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	slow := head
	length := 0
	for slow != nil {
		length++
		slow = slow.Next
	}
	for step := 1; step < length; step *= 2 {
		next := dummyHead.Next
		tail := dummyHead
		for next != nil {
			left := next
			right := cut(left, step)
			// left.str("left")
			// right.str("right1")
			next = cut(right, step)
			// right.str("right2")
			// next.str("next")
			tail.Next = merge(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummyHead.Next
}

// 合并两个有序的链表，并返回链表的头节点和尾节点的下一个节点
func merge(l1, l2 *ListNode) (dummyHead *ListNode) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var tail = &ListNode{}
	dummyHead = tail
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tail.Next = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			l1 = l1.Next
		}
		tail = tail.Next
	}
	if l1 == nil {
		tail.Next = l2
	} else if l2 == nil {
		tail.Next = l1
	}
	return dummyHead.Next
}

//
func cut(l *ListNode, n int) *ListNode {
	var tail *ListNode
	for l != nil && n != 0 {
		n--
		tail = l
		l = l.Next
	}
	if tail == nil {
		return nil
	}
	tail.Next = nil
	return l
}


