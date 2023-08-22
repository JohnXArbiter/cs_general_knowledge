package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
深度是根节点到叶子节点的距离，叶子节点没有孩子，
所以需要判断，就是左边空，取右边的值，右边空，取左边的值
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
