package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 5 {
		usage()
	}

	var arr []int
	for i := 1; i < len(os.Args); i++ {
		if n, err := strconv.Atoi(os.Args[1]); err != nil {
			usage()
		} else {
			arr = append(arr, n)
		}
	}

	fmt.Println(findOddOneOut(arr))
}

func usage() {
	fmt.Printf("Usage: %s [at least 4 ints...]\n", os.Args[0])
	os.Exit(1)
}

func findOddOneOut(arr []int) int {
	once, twice, thrice := 0, 0, 0

	for _, n := range arr {
		// if n re-appears in once, we add it to twice
		twice = twice | (once & n)

		// we flip once with n
		// if it's the 1st time we've seen n, n is tracked
		// if it's the 2nd time we've seen n, n is cleared
		// if it's the 3rd time we've seen n, n is tracked
		once = once ^ n

		// thrice is the numbers we've seen both once *and* twice
		thrice = ^(once & twice)

		// clear the numbers we've seen thrice from both once and twice
		once = once & thrice
		twice = twice & thrice
	}

	return once
}
