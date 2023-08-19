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
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	inOrder(root, &ans)
	return ans
}

func inOrder(cur *TreeNode, ans *[]int) {
	if cur == nil {
		return
	}
	inOrder(cur.Left, ans)
	*ans = append(*ans, cur.Val)
	inOrder(cur.Right, ans)
}

/*
迭代法：
中序遍历是左根右，所以我们需要一直往左子树找直到为空，
前序遍历的迭代法更好处理，是因为当你遍历到根节点时，要处理的也是当前根节点，
但是中序遍历不同，我们反而要先保存根节点，也就是先入栈。
当我们的当前节点为空，也就是上一个根节点的左孩子为空，那么我们就pop这个根节点，
并加入结果集，再去遍历其右子树。
我们循环的条件是当前不为空或者栈长度不为0，
当当前节点为空时，栈不为空，代表左子树遍历结束，准备遍历父亲节点的右子树，
当当前节点不为空，但是栈为空时，也是可以继续的，比如前面都出栈了，但是还有右子树可以继续遍历。
当两个条件都不满足时，说明肯定遍历结束。
*/
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		ans   []int
		stack = list.New()
		cur   = root
	)

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur.Right)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}
