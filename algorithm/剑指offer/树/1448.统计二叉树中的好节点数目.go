package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	var inOrder func(*TreeNode, int)
	inOrder = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if root.Val >= max {
			ans++
			max = root.Val
		}
		inOrder(root.Left, max)
		inOrder(root.Right, max)
	}
	inOrder(root, root.Val)
	return ans
}
