package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	var arr []int
	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			log.Fatalf("Unable to parse %s: %v\n", os.Args[i], err)
		}

		arr = append(arr, n)
	}

	fmt.Println(largestSum(arr))
}

func usage() {
	fmt.Printf("Usage: %s <a list of integers>", os.Args[0])
	os.Exit(1)
}

func largestSum(arr []int) int {
	arrLen := len(arr)

	if arrLen == 0 {
		return 0
	}

	if arrLen == 1 {
		return arr[0]
	}

	arr[1] = max(arr[0], arr[1])
	for i := 2; i < arrLen; i++ {
		arr[i] = max(arr[i-2]+arr[i], arr[i-1])
	}

	return arr[arrLen-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
