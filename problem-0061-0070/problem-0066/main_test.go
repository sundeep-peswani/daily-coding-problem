package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestTossUnbiased(t *testing.T) {
	testRounds := 10000000
	acceptableError := testRounds / 100

	tests := []struct {
		heads float32
	}{
		{0.333333},
		{0.666677},
		{0.9},
		{0.25},
	}

	rand.Seed(time.Now().UnixNano())

	expectedHeads := testRounds / 2
	for _, test := range tests {
		tossBiased := generateTossBiased(test.heads)

		heads := 0
		for i := 0; i < testRounds; i++ {
			if tossUnbiased(tossBiased) == 0 {
				heads++
			}
		}

		if expectedHeads-heads > acceptableError {
			t.Errorf("Testing with head probability of %.3f, expected %d heads, got %d\n", test.heads, expectedHeads, heads)
		}
	}
}

func generateTossBiased(headProbability float32) func() int {
	return func() int {
		if rand.Float32() < headProbability {
			return 0
		}

		return 1
	}
}
