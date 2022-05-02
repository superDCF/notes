package main

import (
	"log"
)

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

思路：
计算左子节点的最长路径和右子节点的最长路径。
*/

func main() {
	// root := &TreeNode{
	// 	1,
	// 	&TreeNode{
	// 		2,
	// 		nil,
	// 		nil,
	// 	},
	// 	nil,
	// }
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				&TreeNode{
					6, nil, nil,
				},
				&TreeNode{
					7,
					&TreeNode{
						12, nil, nil,
					},
					&TreeNode{
						13, nil, nil,
					},
				},
			},
			&TreeNode{
				5,
				&TreeNode{
					8, nil, nil,
				},
				&TreeNode{
					9,
					&TreeNode{
						10, nil, nil,
					},
					&TreeNode{
						11, nil, nil,
					},
				},
			},
		},
		&TreeNode{
			3, nil, nil,
		},
	}
	log.Println(diameterOfBinaryTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func diameterOfBinaryTree(root *TreeNode) int {
    var maxN = 0
	calcNode(root,&maxN)
	return maxN
}

func calcNode(root *TreeNode,maxN *int) int {
	if root == nil {
		return 0
	}
	ln := calcNode(root.Left,maxN)

	rn := calcNode(root.Right,maxN)
	// log.Println(root.Val, ln, rn)
	if ln+rn > *maxN {
		*maxN = ln + rn
	}
	if ln > rn {
		return ln + 1
	}
	return rn + 1
}