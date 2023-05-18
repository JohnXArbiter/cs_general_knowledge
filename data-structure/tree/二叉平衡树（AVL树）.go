package main

func main() {
	//var root *AVLNode

}

type AVLNode struct {
	e      int
	left   *AVLNode
	right  *AVLNode
	height int
}

func LeftRotate(root *AVLNode) *AVLNode {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	newRoot.height = max(getHeight(newRoot.left), getHeight(newRoot.right)) + 1
	return newRoot
}

func RightRotate(root *AVLNode) *AVLNode {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1
	newRoot.height = max(getHeight(newRoot.left), getHeight(newRoot.right)) + 1
	return newRoot
}

func LeftRightRotate(root *AVLNode) *AVLNode {
	root.left = LeftRotate(root.left)
	return RightRotate(root)
}

func RightLeftRotate(root *AVLNode) *AVLNode {
	root.right = RightRotate(root.right)
	return LeftRotate(root)
}

func Insert2AVL(root *AVLNode, e int) *AVLNode {
	if root == nil {
		root = new(AVLNode)
	} else if e < root.e {
		root.left = Insert2AVL(root.left, e)
		if getHeight(root.left)-getHeight(root.right) > 1 {
			if e < root.left.e {
				root = RightRotate(root)
			} else {
				root = LeftRightRotate(root)
			}
		}
	} else if e > root.e {
		root.right = Insert2AVL(root.right, e)
		if getHeight(root.left)-getHeight(root.right) < -1 {
			if e > root.left.e {
				root = LeftRotate(root)
			} else {
				root = RightLeftRotate(root)
			}
		}
	}
	root.height = max(getHeight(root.right), getHeight(root.left)) + 1
	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getHeight(root *AVLNode) int {
	if root == nil {
		return 0
	}
	return root.height
}
