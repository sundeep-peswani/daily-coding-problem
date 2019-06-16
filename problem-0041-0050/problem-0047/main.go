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

	var prices []int
	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			usage()
		}

		prices = append(prices, n)
	}

	fmt.Println(maxProfit(prices))
}

func usage() {
	fmt.Printf("Usage: %s <stock prices...>\n", os.Args[0])
	os.Exit(1)
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	profit := 0

	for _, price := range prices {
		if price > min {
			profit = max(profit, price-min)
		} else if price < min {
			min = price
		}
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
