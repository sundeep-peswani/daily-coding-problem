package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestFindOddOneOut(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{6, 1, 3, 3, 3, 6, 6}, 1},
		{[]int{13, 19, 13, 13}, 19},
	}

	for i, test := range tests {
		actual := findOddOneOut(test.arr)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d: %v\n", i+1, test.expected, actual, test.arr)
		}
	}

	rand.Seed(time.Now().UnixNano())
	arr, magicNumber := generateTestArray(100000)
	actual := findOddOneOut(arr)
	if actual != magicNumber {
		t.Errorf("Randomized test: expected %d, actual %d: %v\n", magicNumber, actual, arr)
	}
}

func generateTestArray(cap int) ([]int, int) {
	var arr []int

	numOfTriplets := rand.Intn(cap)
	for i := 0; i < numOfTriplets; i++ {
		n := rand.Intn(cap)
		for j := 0; j < 3; j++ {
			arr = append(arr, n)
		}
	}

	magicNumber := rand.Intn(cap)
	arr = append(arr, magicNumber)

	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	return arr, magicNumber
}
