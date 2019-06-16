package main

import "testing"

func TestCons(t *testing.T) {
	tables := []struct {
		a, b     int
		f        lof
		expected int
	}{
		{1, 2, sum, 3},
		{1, 2, product, 2},
		{1, -1, sum, 0},
		{1, -1, product, -1},
	}

	for i, table := range tables {
		actual := cons(table.a, table.b)(table.f)
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, table.expected, actual)
		}
	}
}

func TestCar(t *testing.T) {
	tables := []struct {
		a, b, expected int
	}{
		{1, 2, 1},
		{-1, -2, -1},
		{2, 1, 2},
		{-2, -1, -2},
	}

	for i, table := range tables {
		actual := car(cons(table.a, table.b))
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, table.expected, actual)
		}
	}
}

func TestCdr(t *testing.T) {
	tables := []struct {
		a, b, expected int
	}{
		{1, 2, 2},
		{-1, -2, -2},
		{2, 1, 1},
		{-2, -1, -1},
	}

	for i, table := range tables {
		actual := cdr(cons(table.a, table.b))
		if actual != table.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, table.expected, actual)
		}
	}
}

func sum(a, b int) int {
	return a + b
}

func product(a, b int) int {
	return a * b
}
