package main

import (
	"sort"
)

// 直接插入排序是一种简单的插入排序法，其基本思想是： *  把待排序的纪录
// 按其关键码值的大小逐个插入到一个已经排好序的有序序列中，直到所有的纪
// 录插入完为止，得到一个新的有序序列
func InsertionSort(data interface{}) {
	if v, ok := data.(sort.Interface); ok {
		for i := 1; i < v.Len(); i++ {
			j := i - 1
			for j >= 0 && v.Less(j+1, j) {
				v.Swap(j+1, j)
				j--
			}
		}
	}
}
