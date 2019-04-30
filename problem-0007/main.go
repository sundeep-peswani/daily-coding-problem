package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	fmt.Println(countWaysToDecode(os.Args[1]))
}

func countWaysToDecode(msg string) int {
	if len(msg) == 0 {
		return 0
	}

	c := msg[0]
	if len(msg) == 1 {
		return 1
	}

	total := countWaysToDecode(msg[1:])

	if c == '1' || (c == '2' && msg[1] <= '6') {
		if len(msg) > 2 {
			total += countWaysToDecode(msg[2:])
		} else {
			total++
		}
	}

	return total
}

func usage() {
	fmt.Printf("Usage: %s <encoded-message>", os.Args[0])
	os.Exit(1)
}
