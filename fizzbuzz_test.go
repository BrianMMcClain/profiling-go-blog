package main

import (
	"math/rand"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		input  int
		output string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
		{18, "Fizz"},
	}

	for _, test := range cases {
		output := fizzbuzz(test.input)
		if output != test.output {
			t.Fatal(output, "does not match expected value:", test.output)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	// Precalculate the random input
	inputs := make([]int, b.N)
	for i, _ := range inputs {
		inputs[i] = rand.Int()
	}

	b.ResetTimer()

	for _, v := range inputs {
		fizzbuzz(v)
	}
}
