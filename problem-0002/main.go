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
	arrLen := len(arr)

	if arrLen < 2 {
		usage()
	}

	result := make([]int, arrLen)
	left := make([]int, arrLen)
	right := make([]int, arrLen)

	// construct left[] where each element at i is a product of all elements to the left of i
	left[0] = 1
	for i := 1; i < arrLen; i++ {
		left[i] = left[i-1] * arr[i-1]
	}

	// construct right[]  where each element at i is a product of all elements to the right of i
	right[arrLen-1] = 1
	for i := arrLen - 2; i >= 0; i-- {
		right[i] = right[i+1] * arr[i+1]
	}

	// generate result by combining left[] and right[]
	for i := 0; i < arrLen; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

func usage() {
	fmt.Printf("Usage: %s <array of integers>\n", os.Args[0])
	os.Exit(1)
}
