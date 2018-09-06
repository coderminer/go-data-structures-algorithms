package main

import(
	"go-data-structures-algorithms/avltree/avltree"
)

func main(){
	tree := avltree.AVLTree{}
	tree.Insert(14)
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(9)

	tree.PreOrder()
}