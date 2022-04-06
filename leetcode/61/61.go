package main

import (
	"log"
	"strconv"
)

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
*/

func main() {
	l := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	rotateRight(l, 5).String()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() {
	node := l
	s := ""
	log.Println(node.Val)
	for node != nil {
		s += strconv.Itoa(node.Val)
		node = node.Next
	}
	log.Println(s, "1"+"2")
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	h0 := head
	node := head
	var newhead *ListNode
	var tail *ListNode
	n := 0
	total := 0
	for node != nil {
		total++
		tail = node
		node = node.Next
	}
	move := k % total
	if move == 0 {
		return head
	}
	// log.Printf("total=%v move=%v tail=%v", total, move, *tail)
	var preNode *ListNode
	node = head
	for node != nil {
		n++
		if n == total-move+1 {
			newhead = node
			tail.Next = h0
			preNode.Next = nil
			// log.Printf("total=%v move=%v tail=%v newhead=%v", total, move, *tail, *newhead)
			return newhead
		}
		preNode = node
		node = node.Next
	}
	return newhead
}
