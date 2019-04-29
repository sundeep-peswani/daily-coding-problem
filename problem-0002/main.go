package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var arr []int

	if len(os.Args) < 3 {
		usage()
	}

	for i := 1; i < len(os.Args); i++ {
		if t, err := strconv.Atoi(os.Args[i]); err != nil {
			usage()
		} else {
			arr = append(arr, t)
		}
	}

	fmt.Println(arrayProduct(arr))
}

func arrayProduct(arr []int) []int {
	if len(arr) < 2 {
		usage()
	}

	total := 1
	for _, n := range arr {
		total *= n
	}

	result := make([]int, len(arr))
	for i, n := range arr {
		result[i] = total / n
	}

	return result
}

func usage() {
	fmt.Printf("Usage: %s <array of integers>\n", os.Args[0])
	os.Exit(1)
}
