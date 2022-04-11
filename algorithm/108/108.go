/*
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree

输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案。

		0								0
	-3		5						-3		9
-10				9 				-10	  	  5

思路：
题目给的是中序遍历的结果，因为高度平衡，拿到中间的元素，往左右两边递归。
如果用数组存储树，那么树的构造，某节点的左节点在数组中的索引是2n，右节点是2n+1，n从1开始，表示第一个节点。
高度平衡，说明，每个子节点也是高度平衡，递归调用，用某一个区间表达某个区间范围的构造。

如果是前序遍历呢？
		-10						-10
	-3		9				-3		5
0	   5				0				9

-10,-3,9,0,5
-10,-3,5,0,null,null,9
*/

package main

func main() {
	SortedArrayToBST([]int{-10, -3, 0, 5, 9})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	return midHelper(nums, 0, len(nums)-1)
}

// 中序遍历
func midHelper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = midHelper(nums, left, mid-1)
	root.Right = midHelper(nums, mid+1, right)
	return root
}
