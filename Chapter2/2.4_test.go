package main

import (
	"sort"
	"testing"
)

func Test_HeapSort(t *testing.T) {
	ck := getCkints(1000)
	t.Log(ck)
	HeapSort(ck)
	t.Log(ck)
	if !sort.IsSorted(ck) {
		t.Error("HeapSort Error")
	}
}
