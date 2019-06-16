package main

import (
	"fmt"
	"testing"
)

func TestCountWaysToDecode(t *testing.T) {
	tables := []struct {
		msg      string
		expected int
	}{
		{"111", 3}, // aaa, ak, ka
		{"1", 1},   // a
		{"12", 2},  // ab, m
		{"29", 1},  // ai
		{"26", 2},  // af, z
		{"92", 1},  // ia
	}

	for i, table := range tables {
		fmt.Printf("Test %d: testing %s\n", i, table.msg)

		actual := countWaysToDecode(table.msg)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}
