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

	var set []int
	for i := 1; i < len(os.Args); i++ {
		if n, err := strconv.Atoi(os.Args[i]); err != nil {
			usage()
		} else {
			set = append(set, n)
		}
	}

	fmt.Println(powerSet(set))
}

func usage() {
	fmt.Printf("Usage: %s <set of ints...>\n", os.Args[0])
	os.Exit(1)
}

func powerSet(set []int) [][]int {
	var result [][]int
	result = append(result, []int{})

	if len(set) == 0 {
		return result
	}

	for i, s := range set {
		rest := powerSet(set[i+1:])
		for _, r := range rest {
			result = append(result, append([]int{s}, r...))
		}
	}

	return result
}
