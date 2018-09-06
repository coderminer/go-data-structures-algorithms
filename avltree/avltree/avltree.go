package avltree

import (
	"fmt"
)

type Node struct {
	val    int
	height int
	left   *Node
	right  *Node
}

type AVLTree struct {
	root *Node
}

func (tree *AVLTree) Insert(val int) {
	if tree.root == nil{
		tree.root = &Node{val,0,nil,nil}
	}else{
		insert(val,tree.root)
	}
}

func insert(val int, node *Node) *Node {
	if node == nil {
		node = &Node{val, 0, nil, nil}
	} else if val > node.val {
		node.right = insert(val, node.right)
		if height(node.right)-height(node.left) == 2 {
			if val > node.right.val {
				node = leftRotate(node)
			} else {
				node = rightLeftRotate(node)
			}
		}
	} else {
		node.left = insert(val, node.left)
		if height(node.left)-height(node.right) == 2 {
			if val < node.left.val {
				node = rightRotate(node)
			} else {
				node = leftRightRotate(node)
			}
		}
	}

	node.height = max(height(node.left), height(node.right)) + 1
	return node
}

func (tree *AVLTree) Remove(val int) bool {
	return false
}

func (tree *AVLTree) PreOrder() {
	preOrder(tree.root)
}

func preOrder(node *Node) {
	for node != nil {
		fmt.Print(node.val, "->")
		preOrder(node.left)
		preOrder(node.right)
	}
}

func leftRotate(node *Node) *Node {
	rchild := node.right
	node.right = rchild.left
	rchild.left = node

	node.height = max(height(node.left), height(node.right)) + 1
	rchild.height = max(height(rchild.left), height(rchild.right)) + 1
	return rchild
}

func rightRotate(node *Node) *Node {
	lchild := node.left
	node.left = lchild.right
	lchild.right = node

	node.height = max(height(node.left), height(node.right)) + 1
	lchild.height = max(height(lchild.left), height(lchild.right)) + 1
	return lchild
}

func leftRightRotate(node *Node) *Node {
	node.left = leftRotate(node.left)
	return rightRotate(node)
}

func rightLeftRotate(node *Node) *Node {
	node.right = rightRotate(node.right)
	return leftRotate(node)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(node *Node) int {
	if node != nil {
		return node.height
	}
	return 0
}
