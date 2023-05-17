package main

import "fmt"

type TreeNode struct {
	e     any
	left  *TreeNode
	right *TreeNode
}

func main() {
	a := new(TreeNode)
	b := new(TreeNode)
	c := new(TreeNode)
	d := new(TreeNode)
	e := new(TreeNode)
	f := new(TreeNode)

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
	//c.left = nil
	//d.left = nil
	//e.right = nil
	//e.left = nil
	//e.right = nil
	//f.left = nil
	//f.right = nil
	InOrder(a)
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.e)
	PreOrder(root.left)
	PreOrder(root.right)
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Println(root.e)
	InOrder(root.right)
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Println(root.e)
}
