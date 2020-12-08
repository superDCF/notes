package main

import (
	"fmt"
	"log"
)

func main() {
	// addTwoNumbers
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	log.Printf("%s", addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for l != nil {
		log.Println(l.Val)
		s = fmt.Sprintf("%s->%v", s, l.Val)
		l = l.Next
	}
	return s
}

// 单链表只能向后移动，没法向前移动
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var call int
	var tail, head *ListNode
	for l1 != nil || l2 != nil || call >0 {
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
		if tail == nil {
			head = &ListNode{ // 先固定头节点
				Val: (l1.Val + l2.Val + call) % 10,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: (l1.Val + l2.Val + call) % 10,
			}
			tail = tail.Next // 每次移动尾节点
		}
		call = (l1.Val + l2.Val + call) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head
}
