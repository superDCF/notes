package main

import "log"

/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
*/

func main() {
	// 4-3-8-2-10-9-1-7-6-5
	t := &TreeNode{
		7,
		&TreeNode{
			4,
			&TreeNode{
				2,
				&TreeNode{1, nil, nil},
				&TreeNode{3, nil, nil},
			},
			&TreeNode{
				6,
				&TreeNode{5, nil, nil},
				nil,
			},
		},
		&TreeNode{
			10,
			&TreeNode{
				9,
				&TreeNode{8, nil, nil},
				nil},
			nil},
	}
	log.Println(middleTree2(t, 20))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return 0
}

func middleTree(root *TreeNode) {
	if root == nil {
		return
	}
	middleTree(root.Left)
	log.Println(root.Val)
	if root.Val == 7 {
		return
	}
	middleTree(root.Right)
}

func middleTree2(root *TreeNode, k int) int {
	var stack = []*TreeNode{nil}
	for len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		if root == nil {
			return -2
		}
		stack = stack[:len(stack)-1]
		// log.Println(k, root.Val)
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}
