package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 3,
			},
		},
	}
	fmt.Println(levelOrder(root))
}

// 二叉树的层序遍历

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func levelOrder(root *Node) int {
	if root == nil {
		return 0
	}
	r := 0
	queue := list.New()
	queue.PushBack(root)
	fmt.Println(queue.Len(),r)
	for queue.Len() > 0 {
		front := queue.Remove(queue.Front())
		node := front.(*Node)
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
