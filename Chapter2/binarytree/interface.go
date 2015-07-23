package binarytree

import (
	"fmt"
)

type Interface interface {
	AddLeftNode(interface{})
	AddRightNode(interface{})
	GetRightNode() Interface
	GetLeftNode() Interface
	String() string
}

func PreOrder(t Interface) {
	if t != nil {
		fmt.Println(t.String() + ", ")
		fmt.Printf("%#v\n", t)
		PreOrder(t.GetLeftNode())
		PreOrder(t.GetRightNode())
	}
}
func InOrder(t Interface) {
	if t == nil {
		return
	}
	InOrder(t.GetLeftNode())
	fmt.Println(t.String() + ", ")
	InOrder(t.GetRightNode())
}
func PostOrder(t Interface) {
	if t == nil {
		return
	}
	PostOrder(t.GetRightNode())
	PostOrder(t.GetLeftNode())
	fmt.Println(t.String() + ", ")
}
