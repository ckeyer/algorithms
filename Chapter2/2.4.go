/**
 * 堆排序 错误
 */
package main

import (
	// "./binarytree"
	"fmt"
	"sort"
)

func HeapSort(data interface{}, args ...int) {
	if v, ok := data.(sort.Interface); ok {
		heapSort := func(start, end int) {
			fmt.Println("\nS", v)
			root := start
			for {
				child := root*2 + 1
				if child > end {
					break
				}
				if child+1 <= end && v.Less(child, child+1) {
					child++
				}
				if v.Less(root, child) {
					v.Swap(root, child)
				} else {
					break
				}
			}
			fmt.Println("E", v)
		}
		for start := 0; start < v.Len()-2; start++ {
			heapSort(start, v.Len()-1)
		}
		for end := v.Len() - 1; end > 0; end-- {
			v.Swap(0, end)
			heapSort(0, end-1)
		}
	}
}

// func main() {
// 	fmt.Println("Start Sort")
// 	root := binarytree.NewRootNode(1)
// 	root.AddLeftNode("adfsadf")
// 	root.AddRightNode(3)
// 	root.GetRightNode().AddRightNode(5)
// 	root.GetRightNode().AddLeftNode(4)

// 	binarytree.PreOrder(root)
// 	fmt.Println("Sort")
// 	binarytree.InOrder(root)
// 	fmt.Println("Sort")
// 	binarytree.PostOrder(root)
// 	// reflect.ValueOf(root).IsNil()
// 	fmt.Println("End Sort")
// }
