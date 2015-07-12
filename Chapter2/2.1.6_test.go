package main

import (
	"sort"
	"testing"
)

func Test_ShellSort(t *testing.T) {
	for i := 1000; i < 1006; i++ {
		ck := getCkints(i)
		ShellSort(ck)
		if !sort.IsSorted(ck) {
			t.Error("ShellSort Error")
		}
	}
}
