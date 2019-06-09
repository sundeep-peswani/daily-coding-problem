package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestStack(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	arr := generateArray(rand.Intn(10000))
	n := rand.Intn(10000)
	s := newStack()

	if s.max() != stackNil {
		t.Errorf("Expected max() to return %d, received %d\n", stackNil, s.max())
	}

	expectedMax := 0
	for i := 0; i < n; i++ {
		if arr[i] > expectedMax {
			expectedMax = arr[i]
		}
		s.push(arr[i])
	}

	if s.max() != expectedMax {
		t.Errorf("Expected max() to return %d, received %d\n", expectedMax, s.max())
	}

	for i := n - 1; i >= 0; i-- {
		k := s.pop()
		if k != arr[i] {
			t.Errorf("Expected pop() to return %dth item, %d, received %d\n", i, arr[i], k)
		}
		if i != 0 && s.max() == stackNil {
			t.Errorf("Expected max() not to return stacknil for %dth iteration\n", i)
		}
	}

	if s.max() != stackNil {
		t.Errorf("Expected max() to return %d, received %d\n", stackNil, s.max())
	}

	k := s.pop()
	if k != stackNil {
		t.Errorf("Expected pop() to return %d, received %d\n", stackNil, k)
	}
}

func generateArray(n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}

	return arr
}
