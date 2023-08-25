package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	depth int // 全局变量 最大深度
	res   int
)

func findBottomLeftValue(root *TreeNode) int {
	depth, res = 0, 0 // 初始化
	dfs(root, 1)
	return res
}

func dfs(root *TreeNode, d int) {
	if root == nil {
		return
	}
	// 因为先遍历左边，所以左边如果有值，右边的同层不会更新结果
	if root.Left == nil && root.Right == nil && depth < d {
		depth = d
		res = root.Val
	}
	dfs(root.Left, d+1) // 隐藏回溯
	dfs(root.Right, d+1)
}
