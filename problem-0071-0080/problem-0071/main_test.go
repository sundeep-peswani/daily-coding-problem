package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand5(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	count := make([]int, 5)

	for i := 0; i < 5000000; i++ {
		count[rand5()-1]++
	}

	diffCount := 0
	for i := 0; i < 5; i++ {
		diffCount += abs(1000000 - count[i])
	}

	fmt.Println(count)

	if diffCount >= 5000 { // 0.1%
		t.Errorf("rand5() does not seem balanced enough. diff = %d\n", diffCount)
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
