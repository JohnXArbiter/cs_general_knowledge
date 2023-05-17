package main

type TreeNodeS struct {
	e     int
	left  *TreeNodeS
	right *TreeNodeS
}

func main() {

}

func InsertS(root *TreeNodeS, e int) *TreeNodeS {
	if root != nil {
		if e < root.e {
			root.left = InsertS(root.left, e)
		} else if e > root.e {
			root.right = InsertS(root.right, e)
		}
	} else {
		root = &TreeNodeS{
			e: e,
		}
	}
	return root
}

func FindS(root *TreeNodeS, target int) *TreeNodeS {
	for root != nil {
		if target < root.e {
			root = root.left
		} else if target > root.e {
			root = root.right
		} else if target == root.e {
			return root
		}
	}
	return nil
}

func FindMax(root *TreeNodeS) *TreeNodeS {
	for root != nil && root.right != nil {
		root = root.right
	}
	return root
}

func DeleteS(root *TreeNodeS, target int) *TreeNodeS {
	if root == nil {
		return nil
	}
	if root.e > target {
		root.left = DeleteS(root.left, target)
	} else if root.e < target {
		root.right = DeleteS(root.right, target)
	} else {
		// 找到了
		// 先处理当前节点左右孩子都有的情况
		if root.left != nil && root.right != nil {
			max := FindMax(root.left) // 找到左子树中最大的节点
			root.e = max.e            // 直接替换
			root.left = DeleteS(root.left, root.e)
		} else {
			if root.right != nil {
				root = root.right
			} else {
				root = root.left
			}
		}
	}
	return root
}
