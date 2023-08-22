package main

import "math"

/*
也就是后续遍历计算高度，如果当前节点左右子树的高度差大于1，那么就不是平衡二叉树了
*/
func isBalanced(root *TreeNode) bool {
	_, res := q(root)
	return res
}

func q(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, lb := q(root.Left)
	r, rb := q(root.Right)
	if !lb || !rb || math.Abs(float64(l-r)) > 1 {
		return 0, false
	}
	if l > r {
		return l + 1, true
	}
	return r + 1, true
}

func isBalanced1(root *TreeNode) bool {
	h := getHeight(root)
	if h == -1 {
		return false
	}
	return true
}

// 返回以该节点为根节点的二叉树的高度，如果不是平衡二叉树了则返回-1
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := getHeight(root.Left), getHeight(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	return max(l, r) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
