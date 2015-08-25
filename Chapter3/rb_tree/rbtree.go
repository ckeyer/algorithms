package main

import "fmt"

type Color bool

const (
	RED   Color = false
	BLACK Color = true
)

type RBTree struct {
	color  Color
	parent *RBTree
	right  *RBTree
	left   *RBTree
	Node   int
}

// NewRBNode ...
func NewRBNode(n int) (*RBTree, error) {

	return nil, nil
}

// (r *RBTree) 右旋
func (r *RBTree) RotRight() {

}

// (r *RBTree) 右旋
func (r *RBTree) RotLeft() {

}

// (r *RBTree)Insert ...
func (r *RBTree) Insert(newR RBTree) {

}

func main() {
	fmt.Println("")
}
