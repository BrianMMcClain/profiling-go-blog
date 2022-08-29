package main

import (
	"math"
	"math/rand"
	"time"
)

func MonteCarlo(iterations int) float64 {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	hitCount := 0

	for i := 0; i < iterations; i++ {
		hit := generatePoint(randGen)
		if hit {
			hitCount++
		}
	}

	return (float64(hitCount) / float64(iterations)) * 4.0

}

func generatePoint(randGen *rand.Rand) bool {
	x := randGen.Float64()
	y := randGen.Float64()
	hit := (math.Pow(x, 2) + math.Pow(y, 2)) < 1.0
	return hit
}
