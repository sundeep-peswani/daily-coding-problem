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

	needle, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	var haystack []int
	for i := 2; i < len(os.Args); i++ {
		if n, err := strconv.Atoi(os.Args[i]); err != nil {
			usage()
		} else {
			haystack = append(haystack, n)
		}
	}

	fmt.Println(find(haystack, needle))
}

func usage() {
	fmt.Printf("Usage: %s [number to look for] [array of numbers...]\n", os.Args[0])
	os.Exit(1)
}

func find(haystack []int, needle int) int {
	i, _ := rotatedBinarySearch(haystack, needle, 0, len(haystack)-1)
	return i
}

func rotatedBinarySearch(haystack []int, needle, start, end int) (int, int) {
	if end < start || end >= len(haystack) || start < 0 {
		return -1, 1
	}

	mid := (start + end) / 2

	if haystack[mid] == needle {
		return mid, 1
	}

	if start == end {
		return -1, 1
	}

	if haystack[start] < haystack[mid] {
		// left is sorted
		if needle >= haystack[start] && needle <= haystack[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	} else {
		// right is sorted
		if needle >= haystack[mid+1] && needle <= haystack[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	i, count := rotatedBinarySearch(haystack, needle, start, end)
	return i, count + 1
}
