package main

import (
	"log"
)

/*

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。
*/

func main() {
	root := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				17,
				nil,
				nil,
			},
		},
	}
	_ = &TreeNode{
		1,
		nil,
		&TreeNode{
			2, nil, nil,
		},
	}
	log.Println(helper2(root))
	// log.Println(helper2(root2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	helper(root, 0)
	return count
}

var count int

func helper(root *TreeNode, n int) int {
	if root == nil {
		if count > n {
			return 0
		}
		count = n
		// n--
		return count
	}
	n++
	helper(root.Left, n)
	helper(root.Right, n)
	return n
}

func helper2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	log.Println(root.Val)
	l := helper2(root.Left)
	r := helper2(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
