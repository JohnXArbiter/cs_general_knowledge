package main

import "fmt"

type TreeNode2 struct {
	e        any
	left     *TreeNode2
	right    *TreeNode2
	leftTag  int8
	rightTag int8 // 标志位，如果为1表示这一边指针指向的是线索，不为1就是正常的孩子结点
}

func main() {
	a := new(TreeNode2)
	b := new(TreeNode2)
	c := new(TreeNode2)
	d := new(TreeNode2)
	e := new(TreeNode2)
	f := new(TreeNode2)

	a.e = 32
	b.e = 43
	c.e = 5
	d.e = 123
	e.e = 6
	f.e = 8
	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	InOrderThreaded(a)
	InOrder2(a)
}

// 存储前序遍历时，当前节点的前一个节点
var pre = new(TreeNode2)

// PreOrderThreaded 前序遍历线索化函数
func PreOrderThreaded(root *TreeNode2) {
	if root == nil {
		return
	}
	// 如果左儿子为空，那指向前一个节点
	if root.left == nil {
		root.left = pre
		root.leftTag = 1
	}
	// 如果前一个节点的右孩子为空，那就指向当前节点（如果当前是一个右孩子的话，会让左边的兄弟节点指向自己）
	if pre.right == nil {
		pre.right = root
		pre.rightTag = 1
	}
	pre = root
	// 只有当线索为0的时候去执行，不然就死循环了
	if root.leftTag == 0 {
		PreOrderThreaded(root.left)
	}
	if root.rightTag == 0 {
		PreOrderThreaded(root.right)
	}
}

func PreOrder2(root *TreeNode2) {
	if root == nil {
		return
	}
	fmt.Println(root.e)
	if root.leftTag == 0 {
		PreOrder2(root.left)
	} else {
		PreOrder2(root.right)
	}
}

var pre2 *TreeNode2

func InOrderThreaded(root *TreeNode2) {
	if root == nil {
		return
	}
	if root.leftTag == 0 {
		PreOrderThreaded(root.left)
	}
	if root.left == nil {
		root.left = pre2
		root.leftTag = 1
	}
	if pre2 != nil || pre2.right == nil {
		pre2.right = root
		pre2.rightTag = 1
	}
	pre2 = root
	if root.rightTag == 0 {
		PreOrderThreaded(root.right)
	}
}

func InOrder2(root *TreeNode2) {
	for root != nil {
		for root != nil && root.leftTag == 0 {
			root = root.left
		}
		fmt.Println(root.e)
		for root != nil && root.rightTag == 1 {
			root = root.right
			fmt.Println(root.e)
		}
		root = root.right
	}
}

//type TreeNode3 struct {
//	e        any
//	left     *TreeNode3
//	right    *TreeNode3
//	parent   *TreeNode3
//	leftTag  int8
//	rightTag int8 // 标志位，如果为1表示这一边指针指向的是线索，不为1就是正常的孩子结点
//}
//
//var pre3 *TreeNode3
//
//func PostOrderThreaded(root *TreeNode3) {
//	if root == nil {
//		return
//	}
//	if root.leftTag == 0 {
//		PostOrderThreaded(root.left)
//		if root.left != nil {
//			root.left.parent = root
//		}
//	}
//	if root.rightTag == 0 {
//		PostOrderThreaded(root.right)
//		if root.right != nil {
//			root.right.parent = root
//		}
//	}
//	if root.left == nil {
//		root.left = pre3
//		root.leftTag = 1
//	}
//	if pre3 != nil || pre3.right == nil {
//		pre3.right = root
//		pre3.rightTag = 1
//	}
//	pre3 = root
//}
//
//func PostOrder3(root *TreeNode3) {
//
//}
