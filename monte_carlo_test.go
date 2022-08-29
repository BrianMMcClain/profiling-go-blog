package main

import (
	"testing"
)

func TestMonteCarlo(t *testing.T) {
	res := MonteCarlo(1000000)
	if res < 3.1 || res > 3.2 {
		t.Fatalf("Inaccurate estimation of Pi: %v", res)
	}
}

func BenchmarkMonteCarlo(b *testing.B) {
	NUM_ITERATIONS := 1
	for n := 0; n < b.N; n++ {
		MonteCarlo(NUM_ITERATIONS)
	}
}
