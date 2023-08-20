package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			st = append(st, node.Children[i])
		}
	}
	return
}

func preorder2(root *Node) (ans []int) {
	dfs1(root, &ans)
	return
}

func dfs1(root *Node, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	for i := 0; i < len(root.Children); i++ {
		dfs1(root.Children[i], ans)
	}
}
