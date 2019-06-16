package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	var n int
	var err error

	if n, err = strconv.Atoi(os.Args[1]); err != nil || n <= 0 {
		usage()
	}

	var x int
	var steps []int
	for i := 2; i < len(os.Args); i++ {
		if x, err = strconv.Atoi(os.Args[i]); err != nil || x <= 0 {
			usage()
		}
		steps = append(steps, x)
	}

	memo := make(map[int]int)

	fmt.Println(countUniqueWaysToClimb(n, steps, memo))
}

func usage() {
	fmt.Printf("Usage: %s <size of staircase> <steps>...\n", os.Args[0])
	os.Exit(1)
}

func countUniqueWaysToClimb(n int, steps []int, memo map[int]int) int {
	if n <= 0 {
		return 0
	}

	if x, ok := memo[n]; ok {
		return x
	}

	total := 0
	for _, step := range steps {
		if n == step {
			total++
		} else if n > step {
			total += countUniqueWaysToClimb(n-step, steps, memo)
		}
	}

	memo[n] = total

	return total
}
