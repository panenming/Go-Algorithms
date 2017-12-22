package main

//treeNode 和 btree使用binary-tree的
//max newNode 前序，中序，后序遍历可以通用
func insert(root *treeNode, val int) *treeNode {
	if root == nil {
		return newTreeNode(val)
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

func inorderSuccessor(root *treeNode) *treeNode {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func bst_delete(root *treeNode, val int) *treeNode {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = bst_delete(root.left, val)
	} else if val > root.val {
		root.right = bst_delete(root.right, val)
	} else {
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			d := inorderSuccessor(root)
			root.val = d.val
			root.right = bst_delete(root.right, d.val)
		}
	}
	return root

}
