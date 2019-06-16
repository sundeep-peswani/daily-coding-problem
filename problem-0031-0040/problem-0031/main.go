package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	fmt.Println(editDistance(os.Args[1], os.Args[2]))
}

func usage() {
	fmt.Printf("Usage: %s <string> <string>\n", os.Args[0])
	os.Exit(1)
}

func editDistance(a, b string) int {
	al, bl := len(a), len(b)

	matrix := make([][]int, al+1)

	matrix[0] = make([]int, bl+1)
	for j := 0; j < bl; j++ {
		matrix[0][j] = j
	}

	for i := 0; i < al; i++ {
		matrix[i+1] = make([]int, bl+1)
		matrix[i+1][0] = i

		for j := 0; j < bl; j++ {
			m := min3(matrix[i][j], matrix[i][j+1], matrix[i+1][j])

			if a[i] != b[j] {
				m++
			}

			matrix[i+1][j+1] = m
		}
	}

	return matrix[al][len(b)]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}
