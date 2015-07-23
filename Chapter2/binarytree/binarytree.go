package binarytree

import (
	"fmt"
)

type BinaryTree struct {
	Right *BinaryTree
	Left  *BinaryTree
	Data  interface{}
}

func (h *BinaryTree) AddRightNode(v interface{}) {
	h.Right = &BinaryTree{Data: v}
}
func (h *BinaryTree) AddLeftNode(v interface{}) {
	h.Left = &BinaryTree{Data: v}
}
func (h *BinaryTree) GetLeftNode() Interface {
	return h.Left
}
func (h *BinaryTree) GetRightNode() Interface {
	return h.Right
}
func (h *BinaryTree) String() string {
	if h != nil {
		return fmt.Sprint(h.Data)
	} else {
		return ""
	}
}
