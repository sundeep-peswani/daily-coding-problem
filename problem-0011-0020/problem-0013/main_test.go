package main

import "testing"

func TestLongestSubstring(t *testing.T) {
	tables := []struct {
		str      string
		n        int
		expected int
	}{
		{"abcba", 2, 3},
		{"aaaaa", 3, 5},
		{"ababa", 1, 1},
		{"abaaa", 1, 3},
		{"aaabbbbaaa", 1, 4},
		{"ababababab", 2, 10},
		{"ababacabab", 2, 5},
		{"ababacccac", 2, 6},
		{"ababababc", 2, 8},
	}

	for i, table := range tables {
		actual := longestSubstring(table.str, table.n)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a, b, expected int
	}{
		{0, 1, 1},
		{-1, 2, 2},
		{1, 0, 1},
		{2, -1, 2},
	}

	for i, table := range tables {
		actual := max(table.a, table.b)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}

func TestIsValid(t *testing.T) {
	tables := []struct {
		count    []int
		k        int
		expected bool
	}{
		{[]int{0, 0, 0, 0, 1, 0, 1, 1, 0}, 2, false},
		{[]int{0, 0, 0, 0, 1, 0, 1, 1, 0}, 3, true},
		{[]int{0, 0, 0, 0, 1, 0, 1, 1, 0}, 4, true},
	}

	for i, table := range tables {
		actual := isValid(table.count, table.k)
		if actual != table.expected {
			t.Errorf("Test %d: expected %t, actual %t\n", i, table.expected, actual)
		}
	}
}

func TestAlphaIndex(t *testing.T) {
	tables := []struct {
		ch       byte
		expected int
	}{
		{'a', 0},
		{'z', 25},
		{'s', 18},
	}

	for i, table := range tables {
		actual := alphaIndex(table.ch)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d\n", i, table.expected, actual)
		}
	}
}
