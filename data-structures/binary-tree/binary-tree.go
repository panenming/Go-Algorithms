package main

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

type btree struct {
	root *treeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newTreeNode(val int) *treeNode {
	n := &treeNode{val, nil, nil}
	return n
}

//中序遍历
func inorder(n *treeNode) {
	if n == nil {
		return
	}
	inorder(n.left)
	fmt.Print(n.val, " ")
	inorder(n.right)
}

//前序遍历
func preorder(n *treeNode) {
	if n == nil {
		return
	}
	fmt.Print(n.val, " ")
	inorder(n.left)
	inorder(n.right)
}

//后序遍历
func postorder(n *treeNode) {
	if n == nil {
		return
	}
	inorder(n.left)
	inorder(n.right)
	fmt.Print(n.val, " ")
}

// 按照树的结构从第一层开始，从左到右依次排序
func levelorder(root *treeNode) {
	var q []*treeNode
	var n *treeNode
	q = append(q, root)
	for len(q) != 0 {
		n, q = q[0], q[1:]
		fmt.Print(n.val, " ")
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

func _calculate_depth(n *treeNode, depth int) int {
	if n == nil {
		return depth
	}
	return max(_calculate_depth(n.left, depth+1), _calculate_depth(n.right, depth+1))
}

func (t *btree) depth() int {
	return _calculate_depth(t.root, 0)
}
