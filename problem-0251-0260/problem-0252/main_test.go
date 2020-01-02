package main

import "testing"

func Test_GCD(t *testing.T) {
	tests := []struct{
		a, b, expected int
	}{
		{4, 13, 1},
		{4, 16, 4},
		{8, 16, 8},
	}

	for i, test := range tests {
		actual := gcd(test.a, test.b)
		if test.expected != actual {
			t.Errorf("Test %d: expected %d, actual %d\n", i+1, test.expected, actual)
		}
	}
}

func Test_ToEgyptian(t *testing.T) {
	tests := []struct{
		f fraction
		e egyptian
	}{
		{newFraction(4, 13), newEgyptian([]fraction{newFraction(1, 4), newFraction(1, 18), newFraction(1, 468)})},
	}

	for i, test := range tests {
		actual := test.f.toEgyptian()

		if !test.e.equals(actual) {
			t.Errorf("Test %d: expected %s, actual %s\n", i+1, test.e, actual)
		}
	}
}