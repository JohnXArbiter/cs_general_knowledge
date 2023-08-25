package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) (ans int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}

	var inOrder func(*TreeNode, bool)
	inOrder = func(root *TreeNode, left bool) {
		if root == nil {
			return
		}
		if left {
			if root.Left == nil && root.Right == nil {
				ans += root.Val
			}
		}
		inOrder(root.Left, true)
		inOrder(root.Right, false)
	}
	inOrder(root.Left, true)
	inOrder(root.Right, false)
	return
}

func sumOfLeftLeaves1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left) // 左

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val // 中
	}

	rightValue := sumOfLeftLeaves(root.Right) // 右

	return leftValue + rightValue
}

func sumOfLeftLeaves2(root *TreeNode) (ans int) {

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			ans += node.Left.Val
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}
