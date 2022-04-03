package main

import "log"

func main() {
	head := &ListNode{
		1,
		&ListNode{
			2, &ListNode{
				3, &ListNode{},
			},
		},
	}
	hasCycle(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
