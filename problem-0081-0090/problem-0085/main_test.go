package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestPickXY(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var actual uint32
	for i := 0; i < 100; i++ {
		x, y := rand.Uint32(), rand.Uint32()

		actual = pickXYUsingMath(x, y, 1)
		if x != actual {
			t.Errorf("Test %d: expected pickXYUsingMath(%d, %d, %d) to return x, got %d\n", i+1, x, y, 1, actual)
		}
		actual = pickXYUsingMath(x, y, 0)
		if y != actual {
			t.Errorf("Test %d: expected pickXYUsingMath(%d, %d, %d) to return y, got %d\n", i+1, x, y, 0, actual)
		}

		actual = pickXYUsingBitOperations(x, y, 1)
		if x != actual {
			t.Errorf("Test %d: expected pickXYUsingBitOperations(%d, %d, %d) to return x, got %d\n", i+1, x, y, 1, actual)
		}
		actual = pickXYUsingBitOperations(x, y, 0)
		if y != actual {
			t.Errorf("Test %d: expected pickXYUsingBitOperations(%d, %d, %d) to return y, got %d\n", i+1, x, y, 0, actual)
		}
	}
}
