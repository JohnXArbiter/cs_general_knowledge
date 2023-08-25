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

/*
使用的是后序遍历，因为后序求高度，我们做一个判断，能获取最大高度，
最大高度就是深度
*/
func maxDepth(root *TreeNode) (depth int) {
	return postOrder(root)
}

func postOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1 := postOrder(root.Left) + 1
	d2 := postOrder(root.Right) + 1
	if d1 > d2 {
		return d1
	}
	return d2
}

/*
N叉树的最大深度，层序遍历
*/
func maxDepth1(root *Node) (depth int) {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		for i := 0; i < queue.Len(); i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for j := range node.Children {
				queue.PushBack(node.Children[j])
			}
		}
		depth++
	}
	return depth
}
