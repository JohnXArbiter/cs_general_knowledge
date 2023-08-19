package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		ans   []int
		stack = list.New()
	)

	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Back().Value.(*TreeNode)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		} else {
			ans = append(ans, node.Val)
			stack.Remove(stack.Back())
		}
	}
	return ans
}
