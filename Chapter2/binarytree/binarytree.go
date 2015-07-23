package binarytree

import (
	"fmt"
	"reflect"
)

type BinaryTree struct {
	Right *BinaryTree
	Left  *BinaryTree
	Data  interface{}
}

func NewRootNode(v interface{}) (b *BinaryTree) {
	b = &BinaryTree{Data: v}
	return
}
func (b *BinaryTree) AddRightNode(v interface{}) {
	b.Right = &BinaryTree{Data: v}
}
func (b *BinaryTree) AddLeftNode(v interface{}) {
	b.Left = &BinaryTree{Data: v}
}
func (b *BinaryTree) GetLeftNode() Interface {
	return b.Left
}
func (b *BinaryTree) GetRightNode() Interface {
	return b.Right
}
func (b *BinaryTree) IsNil() bool {
	return reflect.ValueOf(b).IsNil()
}
func (b *BinaryTree) String() string {
	if b != nil {
		return fmt.Sprint(b.Data)
	} else {
		return ""
	}
}
