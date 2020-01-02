package main

import (
	"fmt"
	"strings"
)

func main() {

}

type fraction struct {
	numerator, denominator int
}

func newFraction(n, d int) fraction {
	f := fraction{n, d}
	return f.simplify()
}

func (f fraction) toEgyptian() egyptian {
	var result []fraction
	remainder := f

	n := fraction{1, 1}
	for remainder.numerator != 0 {
		for n.greaterThan(remainder) {
			n.denominator++
		}

		result = append(result, n)
		remainder = remainder.subtract(n)
	}

	return newEgyptian(result)
}

func (f fraction) greaterThan(n fraction) bool {
	a, b := f.numerator * n.denominator, n.numerator * f.denominator
	return a > b
}

func (f fraction) simplify() fraction {
	h := gcd(f.numerator, f.denominator)
	return fraction{f.numerator / h, f.denominator / h}
}

func (f fraction) subtract(a fraction) fraction {
	d := f.denominator * a.denominator
	f.numerator = f.numerator * a.denominator - a.numerator * f.denominator
	f.denominator = d
	return f.simplify()
}

func (f fraction) equals(a fraction) bool {
	fs, as := f.simplify(), a.simplify()
	return fs.numerator == as.numerator && fs.denominator == as.denominator
}

func (f fraction) String() string {
	return fmt.Sprintf("%d / %d", f.numerator, f.denominator)
}

type egyptian struct {
	sum []fraction
}

func newEgyptian(input []fraction) egyptian {
	return egyptian{input}
}

func (e egyptian) String() string {
	var strs []string
	for _, f := range e.sum {
		strs = append(strs, f.String())
	}

	return strings.Join(strs, " + ")
}

func (e egyptian) equals(a egyptian) bool {
	if len(e.sum) != len(a.sum) {
		return false
	}

	for i, f := range e.sum {
		if !f.equals(a.sum[i]) {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a % b)
}