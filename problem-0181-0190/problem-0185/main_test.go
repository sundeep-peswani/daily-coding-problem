package main

import "testing"

func TestIntersectingAreas(t *testing.T) {
	tests := []struct{
		a, b rectangle
		expected int
	}{
		{newRect(1, 4, 3, 3), newRect(0, 5, 3, 4), 6},
	}

	for i, test := range tests {
		actual := intersectingAreas(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d:\n%v\n%v\n", i+1, test.expected, actual, test.a, test.b)
		}
	}
}