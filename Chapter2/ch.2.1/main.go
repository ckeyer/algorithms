package main

import (
	"fmt"
	"math/rand"
	"time"
)

type sortFunc func(interface{})
type CkInt []int

func (c CkInt) Len() int {
	return len(c)
}
func (c CkInt) Less(i, j int) bool {
	if c[i] <= c[j] {
		return true
	}
	return false
}
func (c CkInt) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func getCkints(length int) CkInt {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Int()
	ck := make(CkInt, length)
	for i := 0; i < length; i++ {
	loop_get_int:
		n := r.Int()
		for j := 0; j < i; j++ {
			if ck[j] == n {
				goto loop_get_int
			}
		}
		ck[i] = n
	}
	return ck
}

func PerformanceTest(sortf sortFunc) {
	for i := 10; i < 100000; i *= 2 {
		func(N int) {
			start1 := time.Now()
			ck := getCkints(N)
			end1 := time.Now()
			sortf(ck)
			end2 := time.Now()
			fmt.Printf("长度: %d\t生产耗时: %fs\t 排序耗时: %f\n",
				N, end1.Sub(start1).Seconds(), end2.Sub(end1).Seconds())
		}(i)
	}
}

func main() {
	PerformanceTest(SelectionSort)
	PerformanceTest(InsertionSort)
}
