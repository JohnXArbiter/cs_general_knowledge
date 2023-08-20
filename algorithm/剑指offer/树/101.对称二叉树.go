package main

/*
递归判断，只有四种情况：
当左右子树都为空，那肯定对称 true
当有一边为空另一边不空，肯定不对称 false
然后就只能比较当前节点的值了
*/
func isSymmetric(root *TreeNode) bool {
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right != nil {
			return false
		} else if left.Val != right.Val {
			return false
		}

		outside := compare(left.Left, right.Right)
		inside := compare(left.Right, right.Left)
		return outside && inside
	}
	return compare(root.Left, root.Right)
}

/*
不是层序遍历，但是一层层的从外子树到内子树的加入到队列中
*/
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left != nil || right != nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
