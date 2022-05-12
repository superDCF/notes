package main

import (
	"log"
)

/*

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
inorder 和 postorder 都由 不同 的值组成

输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

			20
	3		  		7				3
9		15						9

			3
	7			20
9		15

			3
		9		20
			  15   7

思路：
1. 根据后续遍历，可知尾节点为根结点，例如3是根结点；
2. 根据中序遍历，可知根结点的左边是左子树，右边是右子树，9是3的左子树，15，20，7是3的右子树；
3. 对于剩余元素，递归的求出根节点和左右子树，对于例子的右子树部分，从后序遍历可知顺序是15，7，20，可知右子树的根节点是20，根据中序遍历可知15是左子树，7是右子树
4. 根据中序遍历左节点的元素的个数，在后序遍历中，存在同样的区分左右子树的索引，例如3的左子树个树是1，在后序遍历中，从0到索引1，也都是左子树的节点，同理右子树差不多
*/

func main() {

	instr(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	// log.Println(helper2(root2))
	root := &TreeNode{
		3,
		&TreeNode{
			9,
nil,nil,
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
	instr(root)
	preStr(root)
	poststr(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//非递归前序遍历
func preStr(t *TreeNode) {
	var orders []int
	var nodes []*TreeNode
	t0 := t
	for t0 != nil || len(nodes) != 0 {
		left := t0
		for left != nil {
			orders = append(orders, left.Val)
			nodes = append(nodes, left)
			left = left.Left
		}
		if len(nodes) != 0 {
			node := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			t0 = node.Right
		}
	}
	log.Println(orders)
}

// 非递归中序遍历
func instr(t *TreeNode) {
	var orders []int
	var nodes []*TreeNode
	t0 := t
	for t0 != nil || len(nodes) != 0 {
		left := t0
		for left != nil {
			nodes = append(nodes, left)
			left = left.Left
		}
		if len(nodes) != 0 {
			node := nodes[len(nodes)-1]
			orders = append(orders, node.Val)
			nodes = nodes[:len(nodes)-1]
			t0 = node.Right
		}
	}
	log.Println(orders)
}

// 非递归后序遍历
func poststr(t *TreeNode) {
	var orders []int
	var nodes []*TreeNode
	t0 := t
	var last *TreeNode
	for t0 != nil || len(nodes) != 0 {
		left := t0
		for left != nil {
			nodes = append(nodes, left)
			left = left.Left
		}
		node := nodes[len(nodes)-1]
		if node.Right == nil || node.Right == last {
			orders = append(orders, node.Val)
			nodes = nodes[:len(nodes)-1]
			last = node
			t0 = nil
		} else {
			t0 = node.Right
		}
	}
	log.Println(orders)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rootVal,
	}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			root.Left = buildTree(inorder[0:i], postorder[0:i])                  // 递归构建左子树
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1]) // 递归构建右子树
			break
		}
	}
	return root
}
