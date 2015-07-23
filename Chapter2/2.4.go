package main

import (
	"./binarytree"
	"fmt"
)

func main() {
	fmt.Println("Start Sort")
	root := binarytree.NewRootNode(1)
	root.AddLeftNode(2)
	root.AddRightNode(3)
	root.GetRightNode().AddRightNode(5)
	root.GetRightNode().AddLeftNode(4)

	binarytree.PreOrder(root)
	fmt.Println("Sort")
	binarytree.InOrder(root)
	fmt.Println("Sort")
	binarytree.PostOrder(root)
	// reflect.ValueOf(root).IsNil()
	fmt.Println("End Sort")
}
