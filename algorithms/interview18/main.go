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
	输入两课二叉树A和B，判断B是不是A的子结构
*/
func HasSubTree(root1 *Node, root2 *Node) bool {
	result := false
	if root1 != nil && root2 != nil {
		if root1.val == root2.val {
			result = DoesTree1HasTree2(root1, root2)
		}
		if !result {
			result = HasSubTree(root1.left, root2)
		}
		if !result {
			result = HasSubTree(root1.right, root2)
		}
	}
	return result
}

func DoesTree1HasTree2(root1 *Node, root2 *Node) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}

	if root1.val != root2.val {
		return false
	}

	return DoesTree1HasTree2(root1.left, root2.left) && DoesTree1HasTree2(root1.right, root2.right)
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

	tree1.Print()

	fmt.Println()

	tree2 := BinarySearchTree{}
	tree2.Insert(8)
	tree2.Insert(9)
	tree2.Insert(4)
	tree2.Print()

	fmt.Println()

	r := HasSubTree(tree1.root, tree2.root)
	fmt.Println(r)
}
