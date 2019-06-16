package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand7(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	count := make([]int, 7)

	for i := 0; i < 7000000; i++ {
		count[rand7()-1]++
	}

	diffCount := 0
	for i := 0; i < 7; i++ {
		diffCount += abs(1000000 - count[i])
	}

	fmt.Println(count)

	if diffCount >= 7000 { // 0.1%
		t.Errorf("rand7() does not seem balanced enough. diff = %d\n", diffCount)
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
