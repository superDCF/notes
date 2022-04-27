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
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	l3 := reverseKGroup(list1, 2)
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

func (l *ListNode) str(f string) {
	l0 := l
	for l0 != nil {
		log.Println(f, l0.Val)
		l0 = l0.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	k0 := k
	var pre *ListNode
	dh := &ListNode{0, head}
	cur := head
	tail := dh
	for cur != nil && k != 0 {
		// dh := cur
		for tail != nil && k != 0 {
			if tail == nil {
				return dh.Next
			}
			k--
			tail = tail.Next
		}
		nxt := tail.Next
		h, t := reverseList(cur)
		t.Next = nxt
		cur = nxt
		pre = h

	}
	pre.str("pre=")
	cur.str("cur=")
	tail.str("tail=")
	return dh.Next
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre, head
}
