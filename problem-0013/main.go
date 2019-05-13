package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	str := os.Args[1]
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		usage()
	}

	fmt.Println(longestSubstring(str, n))
}

func usage() {
	fmt.Printf("Usage: %s <string> <num of unique characters>\n", os.Args[0])
	os.Exit(1)
}

func longestSubstring(str string, k int) int {
	if k == 0 {
		return 0
	}

	if len(str) <= k {
		return len(str)
	}

	count := make([]int, 26)
	start, maxLen := 0, 1

	count[alphaIndex(str[0])]++

	for i := 1; i < len(str); i++ {
		count[alphaIndex(str[i])]++

		for !isValid(count, k) {
			count[alphaIndex(str[start])]--
			start++
		}

		maxLen = max(maxLen, i-start+1)
	}

	fmt.Printf("Longest is %d to %d for %s\n", start, start+maxLen, str)

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func isValid(count []int, k int) bool {
	uniq := 0
	for _, c := range count {
		if c == 0 {
			continue
		}

		uniq++
		if uniq > k {
			return false
		}
	}

	return true
}

func alphaIndex(ch byte) int {
	return int(ch - 'a')
}
