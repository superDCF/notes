package main

import "log"

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
0->1->2->3->4
*/

func main() {
	l := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}
	reverseList(l).str()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) str() {
	l0 := l
	for l0 != nil {
		log.Println(l0.Val)
		l0 = l0.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	return reverse(head, nil)
}

func reverse(cur, pre *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	reverse(tmp, cur)
	return cur
}
