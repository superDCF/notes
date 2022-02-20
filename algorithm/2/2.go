package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	root := &ListNode{
		-1, nil,
	}
	root = addNode(root)
	for root != nil {
		log.Println(root.Val)
		root = root.Next
	}
	l1 := &ListNode{
		9,
		&ListNode{
			9,
			&ListNode{
				9, nil},
		},
	}
	l2 := &ListNode{
		9,
		&ListNode{
			9,nil,
		},
	}
	list := addTwoNumbers(l1, l2)
	for list != nil {
		log.Println(list.Val)
		list = list.Next
	}
}


/* 
1. 考虑如何向链表添加后一个节点
2. 考虑如何遍历链表，以最长为准还是最短
3. 考虑哪些是重置变量，以及重置变量出现的顺序
4. 考虑边界条件
*/


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	node := root
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2
		node.Next = new(ListNode)
		sum = sum  + carry // 有坑的先标出来
		node.Next.Val = sum % 10
		carry = sum / 10
		log.Println(v1, v2, sum, carry)
		node = node.Next
	}
	return root.Next
}

func addNode(root *ListNode) *ListNode {
	node := root
	for i := 0; i < 3; i++ {
		tmp := node
		node = new(ListNode)
		node.Val = i
		tmp.Next = node
		log.Println(tmp, tmp.Next)
	}
	return root
}
