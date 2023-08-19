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
层序遍历：
使用队列将当前这层的节点除掉，同时在后面补上他的子节点
能保证每次循环时，只存在当前层的节点数，所一开始要获取队列长度，再进行内层循环
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var (
		ans   [][]int
		queue = list.New()
	)

	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		var level []int
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		size *= 2
		ans = append(ans, level)
	}
	return ans
}
