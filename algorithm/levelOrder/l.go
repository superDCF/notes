package main

import (
	"container/list"
	"fmt"
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
	fmt.Println(preOrder(root))
}

// 二叉树的前序遍历

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func preOrder(root *Node) int {
	if root == nil {
		return 0
	}
	r := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		front := queue.Remove(queue.Front())
		node := front.(*Node)
		fmt.Println(node.Val)
		r += node.Val
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return r
}
