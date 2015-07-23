package main

import (
	"./binarytree"
	"fmt"
)

func main() {
	fmt.Println("Start Sort")
	root := &binarytree.BinaryTree{Data: 11}
	root.Data = 11
	// root.AddLeftNode(22)
	// root.AddRightNode(33)
	// root.GetRightNode().AddRightNode(11)
	binarytree.InOrder(root)
	// binarytree.PreOrder(root)
	fmt.Println("End Sort")
}
