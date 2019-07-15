package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestFind(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	min, padding := 9, 10

	for i := 0; i < 10; i++ {
		arr := generateArray(min+rand.Intn(padding), rand.Intn(min))
		expected := rand.Intn(len(arr))

		actual := find(arr, arr[expected])
		if actual != expected {
			t.Errorf("Test %d: expected %d, actual %d\narray: %v\n", i+1, expected, actual, arr)
		}

		smallest := rand.Intn(min)
		if find(arr, smallest) != -1 {
			t.Errorf("Test %d: found %d unexpectedly\narray: %v\n", i+1, smallest, arr)
		}

		largest := int(math.Pow(float64(min+padding), 2.0))
		if find(arr, largest) != -1 {
			t.Errorf("Test %d: found %d unexpectedly\narray: %v\n", i+1, largest, arr)
		}
	}
}

func TestRotatedBinarySearch(t *testing.T) {
	tests := []struct {
		haystack []int
		needle   int
		expected int
	}{
		{[]int{13, 18, 25, 2, 8, 10}, 8, 4},
		{[]int{13, 18, 25, 2, 8, 10}, 25, 2},
		{[]int{13, 18, 25, 2, 8, 10}, 1, -1},
		{[]int{13, 18, 25, 2, 8, 10}, 26, -1},
	}

	for i, test := range tests {
		actual, count := rotatedBinarySearch(test.haystack, test.needle, 0, len(test.haystack)-1)

		if actual != test.expected {
			t.Errorf("Test %d: expected result of %d, got %d\n", i+1, test.expected, actual)
		}
		if count >= len(test.haystack) {
			t.Errorf("Test %d: expected better than O(n) performance, got %d operations\n", i+1, count)
		}
	}
}

func generateArray(len, rotation int) []int {
	arr := make([]int, len)
	last := len

	for i := 0; i < len; i++ {
		n := last + 1 + rand.Intn(len)

		arr[(i+rotation)%len] = n
		last = n
	}

	return arr
}
