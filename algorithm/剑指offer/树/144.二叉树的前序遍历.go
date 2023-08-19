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
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	preOrder(root, &ans)
	return ans
}

func preOrder(cur *TreeNode, ans *[]int) {
	if cur == nil {
		return
	}
	*ans = append(*ans, cur.Val)
	preOrder(cur.Left, ans)
	preOrder(cur.Right, ans)
}

/*
	 迭代法
		其实就是我们自己使用栈来模拟底层的递归，

当我们遍历到某个根节点时，按照前序遍历，将其放入我们的答案中，
然后将他非空的孩子节点加入栈中，先加入右孩子，再加入左孩子，
因为对于栈的结构，先进后出，并且我们本来就应该先找左孩子。
遍历开始，我们先将整棵树的根节点加入栈中，这样能保证栈中一直有元素，
因为我们的循环条件就是栈非空，为什么呢，因为上面我们要把后摇进行遍历的节点先入栈吗，
所以栈不可能为空，为空则代表元素都被遍历完了
*/
func preorderTraversal2(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

// 有问题，但思想差不多会了
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		ans   []int
		stack []*TreeNode
	)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}
	return ans
}
