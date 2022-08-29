package main

import (
	"sort"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []int{5, 4, 3, 2, 1}
		sort.Ints(s)
	}
}
