package main

import (
	"sort"
)

// 选择排序：从所有序列中先找到最小的，然后放到第一个位置。
// 之后再看剩余元素中最小的，放到第二个位置……
func SelectionSort(data interface{}, args ...int) {
	if v, ok := data.(sort.Interface); ok {
		for i := 0; i < v.Len()/2; i++ {
			min, max := i, i
			for j := i + 1; j < v.Len()-i; j++ {
				if v.Less(j, min) {
					min = j
				}
				if v.Less(max, j) {
					max = j
				}
			}
			v.Swap(min, i)
			if max == i {
				v.Swap(min, v.Len()-i-1)
			} else {
				v.Swap(max, v.Len()-i-1)
			}
		}
	}
}
