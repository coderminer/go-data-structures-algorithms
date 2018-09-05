package binarysearchtree

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

func (tree *BinarySearchTree) Insert(val int) {
	if tree.root == nil {
		tree.root = &Node{val, nil, nil}
	} else {
		insertNode(val, tree.root)
	}
}

func insertNode(val int, root *Node) {
	if val < root.val {
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

func (tree *BinarySearchTree) Search(val int) *Node {
	if tree.root != nil {
		return findNode(val, tree.root)
	}
	return nil
}

func findNode(val int, root *Node) *Node {
	if val == root.val {
		return root
	} else if val < root.val {
		if root.left != nil {
			return findNode(val, root.left)
		}
	} else {
		if root.right != nil {
			return findNode(val, root.right)
		}
	}
	return nil
}

func (tree *BinarySearchTree) Parent(val int) *Node {
	if tree.root != nil {
		return findParent(val, tree.root)
	}
	return nil
}

func findParent(val int, root *Node) *Node {
	if val == root.val {
		return nil
	} else if val < root.val {
		if root.left != nil {
			if val == root.left.val {
				return root
			} else {
				return findParent(val, root.left)
			}
		}
	} else {
		if root.right != nil {
			if val == root.right.val {
				return root
			} else {
				return findParent(val, root.right)
			}
		}
	}
	return nil
}

func (tree *BinarySearchTree) Contains(val int) bool {
	if tree.root != nil {
		return containNode(val, tree.root)
	}
	return false
}

func containNode(val int, root *Node) bool {
	if val == root.val {
		return true
	} else if val < root.val {
		if root.left == nil {
			return false
		} else {
			return containNode(val, root.left)
		}
	} else {
		if root.right == nil {
			return false
		} else {
			return containNode(val, root.right)
		}
	}
}

func (tree *BinarySearchTree) Remove(val int) bool {
	nodeToRemoved := tree.Search(val)
	if nodeToRemoved == nil {
		return false
	}

	parent := tree.Parent(val)
	if tree.root.left == nil && tree.root.right == nil {
		tree.root = nil
		return true
	} else if nodeToRemoved.left == nil && nodeToRemoved.right == nil {
		if nodeToRemoved.val < parent.val {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if nodeToRemoved.left == nil && nodeToRemoved.right != nil {
		if nodeToRemoved.val < parent.val {
			parent.left = nodeToRemoved.right
		} else {
			parent.right = nodeToRemoved.right
		}
	} else if nodeToRemoved.left != nil && nodeToRemoved.right == nil {
		if nodeToRemoved.val < parent.val {
			parent.left = nodeToRemoved.left
		} else {
			parent.right = nodeToRemoved.left
		}
	} else {
		largestValue := nodeToRemoved.left
		for largestValue.right != nil {
			largestValue = largestValue.right
		}
		tree.Remove(largestValue.val)
		nodeToRemoved.val = largestValue.val
	}

	return true
}

func (tree *BinarySearchTree) Min() *Node {
	if tree.root != nil {
		return findMin(tree.root)
	}
	return nil
}

func findMin(root *Node) *Node {
	if root.left == nil {
		return root
	} else {
		return findMin(root.left)
	}
}

func (tree *BinarySearchTree) Max() *Node {
	if tree.root != nil {
		return findMax(tree.root)
	}
	return nil
}

func findMax(root *Node) *Node {
	if root.right == nil {
		return root
	} else {
		return findMax(root.right)
	}
}

func (tree *BinarySearchTree) PreOrder() {
	preOrder(tree.root)
}

func preOrder(root *Node) {
	if root != nil {
		fmt.Print(root.val, "->")
		preOrder(root.left)
		preOrder(root.right)
	}
}

func (tree *BinarySearchTree) InOrder() {
	inOrder(tree.root)
}

func inOrder(root *Node) {
	if root != nil {
		inOrder(root.left)
		fmt.Print(root.val, "->")
		inOrder(root.right)
	}
}

func (tree *BinarySearchTree) PostOrder() {
	postOrder(tree.root)
}

func postOrder(root *Node) {
	if root != nil {
		postOrder(root.right)
		postOrder(root.left)
		fmt.Print(root.val, "->")
	}
}
