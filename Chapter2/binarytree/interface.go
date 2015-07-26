package binarytree

import (
	"fmt"
)

type Interface interface {
	AddLeftNode(interface{})
	AddRightNode(interface{})
	GetRightNode() Interface
	GetLeftNode() Interface
	IsNil() bool
	String() string
}

func PreOrder(t Interface) {
	if t.IsNil() {
		return
	}
	fmt.Println(t.String())
	PreOrder(t.GetLeftNode())
	PreOrder(t.GetRightNode())
}
func InOrder(t Interface) {
	if t.IsNil() {
		return
	}
	InOrder(t.GetLeftNode())
	fmt.Println(t.String())
	InOrder(t.GetRightNode())
}
func PostOrder(t Interface) {
	if t.IsNil() {
		return
	}
	PostOrder(t.GetRightNode())
	PostOrder(t.GetLeftNode())
	fmt.Println(t.String())
}
