package main

import (
	"sort"
)

func MergeSort(data interface{}) {
	if v, ok := data.(sort.Interface); ok {
		_ = v
	}
}
