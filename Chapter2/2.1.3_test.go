package main

import (
	"sort"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	ck := getCkints(1000)
	InsertionSort(ck)
	if !sort.IsSorted(ck) {
		t.Error("InsertionSort Error")
	}
}
