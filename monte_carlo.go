package main

import (
	"math"
	"math/rand"
	"time"
)

func MonteCarloPow(iterations int) float64 {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	hitCount := 0

	for i := 0; i < iterations; i++ {
		x := randGen.Float64()
		y := randGen.Float64()
		hit := (math.Pow(x, 2) + math.Pow(y, 2)) <= 1.0
		if hit {
			hitCount++
		}
	}

	return (float64(hitCount) / float64(iterations)) * 4.0
}

func MonteCarloSquare(iterations int) float64 {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	hitCount := 0

	for i := 0; i < iterations; i++ {
		x := randGen.Float64()
		y := randGen.Float64()
		hit := ((x * x) + (y * y)) <= 1
		if hit {
			hitCount++
		}
	}

	return (float64(hitCount) / float64(iterations)) * 4.0
}
