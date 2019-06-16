package main

import "fmt"

func main() {
	fmt.Printf("car(cons(3, 4)): %d\n", car(cons(3, 4)))
	fmt.Printf("cdr(cons(3, 4)): %d\n", cdr(cons(3, 4)))
}

type lof func(int, int) int

type hof func(lof) int

func cons(a, b int) hof {
	pair := func(f lof) int {
		return f(a, b)
	}
	return pair
}

func car(f hof) int {
	left := func(a, b int) int {
		return a
	}

	return f(left)
}

func cdr(f hof) int {
	right := func(a, b int) int {
		return b
	}

	return f(right)
}
