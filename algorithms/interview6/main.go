package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func construct(preOrder []int, inOrder []int) *Node {
	if preOrder == nil || inOrder == nil || len(preOrder) != len(inOrder) || len(inOrder) < 1 {
		return nil
	}
	return constructCore(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1)
}

/*
	根据二叉树的前序遍历和中序遍历，构建一个二叉树
*/

func constructCore(preOrder []int, preStart, preEnd int, inOrder []int, inStart, inEnd int) *Node {
	if preStart > preEnd {
		return nil
	}

	//前序遍历中的第一个数字，作为当前的根节点
	val := preOrder[preStart]
	index := inStart
	// 在中序遍历中查找根节点的位置
	for index <= inEnd && inOrder[index] != val {
		index += 1
	}
	if index > inEnd {
		return nil
	}

	//创建跟节点
	node := &Node{}
	node.val = val

	// 递归构建当前根节点的左子树，左子树的元素个数：index - inStart + 1
	// 左子树对应的前序遍历的位置在[preStart+1,preStart+index-inStart]
	// 左子树对应的中序遍历的位置在[isStart,index-1]
	node.left = constructCore(preOrder, preStart+1, preStart+index-inStart, inOrder, inStart, index-1)

	// 递归构建当前根节点的右子树，右子树的元素个数: inEnd - index
	// 右子树对应的前序遍历的位置在[preStart+index-inStart+1,preEnd]
	// 右子树对应的中序遍历的位置在[index+1,inEnd]
	node.right = constructCore(preOrder, preStart+index-inStart+1, preEnd, inOrder, index+1, inEnd)

	return node
}

func printTree(root *Node) {
	if root != nil {
		printTree(root.left)
		fmt.Print(root.val, " ")
		printTree(root.right)
	}
}

func test1() {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	in := []int{4, 7, 2, 1, 5, 3, 8, 6}
	root := construct(pre, in)
	printTree(root)
}

func test2() {
	pre := []int{1, 2, 3, 4, 5}
	in := []int{5, 4, 3, 2, 1}
	root := construct(pre, in)
	printTree(root)
}

func test3() {
	pre := []int{1, 2, 3, 4, 5}
	in := []int{1, 2, 3, 4, 5}
	root := construct(pre, in)
	printTree(root)
}

func test4() {
	pre := []int{1}
	in := []int{1}
	root := construct(pre, in)
	printTree(root)
}

func main() {
	test1()
	fmt.Println()
	test2()
	fmt.Println()
	test3()
	fmt.Println()
	test4()
}
