package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) Insert(val int) {
	if t.root == nil {
		t.root = &Node{val, nil, nil}
	} else {
		insertNode(val, t.root)
	}
}

func insertNode(val int, root *Node) {
	if val <= root.val {
		if root.left == nil {
			root.left = &Node{val, nil, nil}
		} else {
			insertNode(val, root.left)
		}
	} else {
		if root.right == nil {
			root.right = &Node{val, nil, nil}
		} else {
			insertNode(val, root.right)
		}
	}
}

func (t *BinarySearchTree) Print() {
	printNode(t.root)
}

func printNode(root *Node) {
	if root != nil {
		fmt.Println(root.val)
		printNode(root.left)
		printNode(root.right)
	}
}

/*
	根据二叉树，获取该二叉树的镜像
*/
func Mirror(root *Node) {
	if root == nil  {
		return
	}
	root.left, root.right = root.right, root.left
	if root.left != nil {
		Mirror(root.left)
	}
	if root.right != nil {
		Mirror(root.right)
	}
}

func main() {
	tree1 := BinarySearchTree{}
	tree1.Insert(10)
	tree1.Insert(8)
	tree1.Insert(9)
	tree1.Insert(2)
	tree1.Insert(4)
	tree1.Insert(7)
	tree1.Insert(12)

	//tree1.Print()
	Mirror(tree1.root)
	tree1.Print()
}
