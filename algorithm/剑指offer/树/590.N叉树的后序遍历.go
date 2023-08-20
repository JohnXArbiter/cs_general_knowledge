package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for i := 0; i < len(node.Children); i++ {
			dfs(node.Children[i])
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}

/*
迭代法：
我们使用 map 记录某个根节点被遍历到的第几棵子树
因为是后序遍历，所以我们从先向后找当前根的子树，
所以比如第一次，我们就需要找到最左的子树，我们将每一次遍历到的根都保存在栈中，直到为空。
那我就拿出栈顶节点，也就是最左的那个，为了通用性，我们也会判断他的子树的情况，
那么情况是他没有子树，所以我们就把他加入到结果集中。
接下来，我们的栈顶是就是刚刚那个节点的父节点，所以，我们就能遍历到他的兄弟节点了，
因为我们不是用 map 保存了每个节点当前遍历子节点的下标吗，所以如果从 map 中拿出的下标小于栈顶节点的孩子长度，
那我们就将当前 node 变成孩子，继续去遍历
*/
func postorder1(root *Node) (ans []int) {
	if root == nil {
		return []int{}
	}

	var (
		m     = map[*Node]int{}
		stack []*Node
		node  = root
	)

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			if len(node.Children) <= 0 {
				break
			}
			m[node] = 1
			node = node.Children[0]
		}
		node = stack[len(stack)-1]
		i := m[node]
		if i < len(node.Children) {
			m[node]++
			node = node.Children[i]
		} else {
			ans = append(ans, node.Val)
			stack = stack[:len(stack)-1]
			delete(m, node)
			node = nil
		}
	}
	return
}
