package main

import (
	"fmt"
)

func main() {
	t := btree{nil}
	t.root = newTreeNode(0)
	t.root.left = newTreeNode(1)
	t.root.right = newTreeNode(2)
	t.root.left.left = newTreeNode(3)
	t.root.left.right = newTreeNode(4)
	t.root.right.left = newTreeNode(5)
	t.root.right.right = newTreeNode(6)
	t.root.right.right.right = newTreeNode(10)
	inorder(t.root)
	fmt.Print("\n")
	preorder(t.root)
	fmt.Print("\n")
	postorder(t.root)
	fmt.Print("\n")
	levelorder(t.root)
	fmt.Print("\n")
	fmt.Print(t.depth(), "\n")

	fmt.Println("二叉查找树-----")
	bst := &btree{nil}
	inorder(bst.root)
	bst.root = insert(bst.root, 10)
	bst.root = insert(bst.root, 20)
	bst.root = insert(bst.root, 15)
	bst.root = insert(bst.root, 30)
	fmt.Print(bst.depth(), "\n")
	inorder(bst.root)
	fmt.Print("\n")
	bst.root = bst_delete(bst.root, 20)
	inorder(bst.root)
	fmt.Print("\n")
	bst.root = bst_delete(bst.root, 30)
	inorder(bst.root)
	fmt.Print("\n")
	bst.root = bst_delete(bst.root, 15)
	inorder(bst.root)
	fmt.Print("\n")
	bst.root = bst_delete(bst.root, 10)
	inorder(bst.root)
	fmt.Print("\n")
	fmt.Print(bst.depth(), "\n")
}
