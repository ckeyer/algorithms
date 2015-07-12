package main

import (
	"sort"
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	ck := getCkints(1000)
	SelectionSort(ck)
	if !sort.IsSorted(ck) {
		t.Error("SelectionSort Error")
	}
}
