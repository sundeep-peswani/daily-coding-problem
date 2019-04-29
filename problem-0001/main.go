package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total int
	var arr []int

	if len(os.Args) < 2 || os.Args[1] == "help" {
		usage()
	}

	if t, err := strconv.Atoi(os.Args[1]); err != nil {
		usage()
	} else {
		total = t
	}

	for i := 2; i < len(os.Args); i++ {
		if t, err := strconv.Atoi(os.Args[i]); err != nil {
			usage()
		} else {
			arr = append(arr, t)
		}
	}

	fmt.Println(findSum(total, arr))
}

func findSum(total int, arr []int) bool {
	if len(arr) < 2 {
		return false
	}

	remaining := make(map[int]bool)

	for _, i := range arr {
		if _, ok := remaining[i]; ok {
			return true
		}

		remaining[total-i] = true
	}

	return false
}

func usage() {
	fmt.Printf("Usage: %s <total> <array of integers>\n", os.Args[0])
	os.Exit(1)
}
