package main

import (
	"fmt"
	"go-data-structures-algorithms/binarysearchtree/binarysearchtree"
)

func main() {
	tree := binarysearchtree.BinarySearchTree{}

	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(18)
	tree.Insert(9)
	tree.Insert(12)

	tree.PreOrder()
	fmt.Println()
	tree.InOrder()
	fmt.Println()
	tree.PostOrder()
	fmt.Println()
	//tree.LevelOrder()
	fmt.Println()
	fmt.Println(tree.Min())
	fmt.Println(tree.Max())
	fmt.Println(tree.Contains(28))
	fmt.Println(tree.Search(25))
	fmt.Println(tree.Parent(12))

	tree.Remove(14)
	tree.PreOrder()
	fmt.Println()
	tree.LevelOrder()
}
