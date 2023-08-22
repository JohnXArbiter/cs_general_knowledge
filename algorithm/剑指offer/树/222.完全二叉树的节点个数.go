package main

/*
这样写是后序遍历
递归终止条件肯定还是当前节点为空返回0，
因为每次进入函数都是先走左边（右边也行），所以就会遍历到最下层，
这时候再返回当前统计的节点数即可，
为什么不是前序：因为前序在下潜的过程中统计了，但还没进入下一次递归你都不知道下面有多少个节点
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

/*
利用完全二叉树的性质：
如果从某一个节点开始，左边和右边层数一样，那一定是个满二叉树，因为完全二叉树中间不可能缺
如果是，那么对应当前节点一共的节点数就是 2^高度 - 1，
那么只存在是满二叉树或者不是的情况，不是的时候就直接按照以前的写法返回结果，这样可以减少遍历次数
*/
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}
	if leftH == rightH {
		return (2 << leftH) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
