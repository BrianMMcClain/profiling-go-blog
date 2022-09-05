package main

import (
	"testing"
)

func TestMonteCarlo(t *testing.T) {
	res := MonteCarloPow(1000000)
	if res < 3.1 || res > 3.2 {
		t.Fatalf("Inaccurate estimation of Pi: %v", res)
	}
}

func BenchmarkMonteCarlo(b *testing.B) {
	NUM_ITERATIONS := 1000000
	for n := 0; n < b.N; n++ {
		// Swap for MonteCarloSquare for comparison
		MonteCarloSquare(NUM_ITERATIONS)
	}
}
