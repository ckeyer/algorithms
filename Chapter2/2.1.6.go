package main

import (
	"fmt"
	"sort"
)

func ShellSort(data interface{}, args ...int) {
	if v, ok := data.(sort.Interface); ok {
		for gap := v.Len() / 2; gap > 0; gap /= 2 {
			for i := gap; i < v.Len(); i++ {
				for j := i - gap; j >= 0 && v.Less(j+gap, j); j -= gap {
					v.Swap(j, j+gap)
				}
			}
		}
	}
}

func BetterShellSort(data interface{}) {
	if v, ok := data.(sort.Interface); ok {
		for gap := v.Len() / 2; gap > 0; gap /= 2 {
			for i := gap; i < v.Len(); i++ {
				for j := i - gap; j >= 0 && v.Less(j+gap, j); j -= gap {
					// fmt.Printf("j: %d, i: %d, gap:%d %v\n", j, i, gap, v)
					v.Swap(j, j+gap)
					fmt.Println(v)
					// fmt.Println()
				}
			}
		}
	}
}
