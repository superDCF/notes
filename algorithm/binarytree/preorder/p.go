package main

import (
	"container/list"
	"log"
)

/*
			1
		2		3
	4	  5   6		7

*/
func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	preOrderRecur(root)
	preOrder(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 前序遍历 1->2->4->5->3->6->7
// 前序递归遍历
func preOrderRecur(root *Node) {
	if root == nil {
		return
	}
	log.Println("preOrderRecur", root.Val)
	preOrderRecur(root.Left)
	preOrderRecur(root.Right)
}

// 前序迭代遍历
func preOrder(root *Node) {
	if root == nil {
		return
	}
	stack := list.New()
	stack.PushFront(root)
	for stack.Len() != 0 {
		front := stack.Remove(stack.Front()).(*Node)
		log.Println("preOrder", front.Val)
		if front.Right != nil {
			stack.PushFront(front.Right)
		}
		if front.Left != nil {
			stack.PushFront(front.Left)
		}
	}
}
