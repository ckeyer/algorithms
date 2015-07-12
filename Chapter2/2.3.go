package main

import (
	"sort"
)

func QuickSort(data interface{}, args ...int) {
	if v, ok := data.(sort.Interface); ok {
		left, right := 0, v.Len()-1
		if len(args) == 2 {
			left, right = args[0], args[1]
		}
		p, _ := left, left
		i, j := left, right
		if left >= right {
			return
		}
		for i < j {
			for j > p && v.Less(p, j) {
				j--
			}
			if j > p {
				v.Swap(p, j)
				p = j
			}
			for v.Less(i, p) && i < p {
				i++
			}
			if p > i {
				v.Swap(i, p)
				p = i
			}
		}
		QuickSort(v, left, i-1)
		QuickSort(v, j+1, right)
	}
}
