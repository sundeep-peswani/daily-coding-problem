package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	var arr []int
	for i := 1; i < len(os.Args); i++ {
		if t, err := strconv.Atoi(os.Args[i]); err != nil {
			usage()
		} else {
			arr = append(arr, t)
		}
	}

	fmt.Println(findFirstMissingPositiveInteger(arr))
}

// use negative as an indicator of having seen the value
// e.g. if 3 is seen, turn arr[3 - 1] into a negative
// arr[0] represents 1, arr[1] represents 2, etc.
func findFirstMissingPositiveInteger(arr []int) int {
	// clear all negatives
	for i, n := range arr {
		if n < 0 {
			arr[i] = 0
		}
	}

	var j int
	for _, n := range arr {
		if n == 0 {
			continue
		}

		if n < 0 {
			j = (n * -1) - 1
		} else {
			j = n - 1
		}

		if arr[j] > 0 {
			arr[j] *= -1
		} else if arr[j] == 0 {
			arr[j] = -1
		}
	}

	// find first index for which the value is positive
	for i, n := range arr {
		if n >= 0 {
			return i + 1
		}
	}

	return len(arr) + 1
}

func usage() {
	fmt.Printf("Usage: %s <array of integers>\n", os.Args[0])
	os.Exit(1)
}
