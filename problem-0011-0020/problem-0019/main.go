package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	K, err := strconv.Atoi(os.Args[2])
	if err != nil {
		usage()
	}

	costMatrix := generateRandomCostMatrix(N, K)
	fmt.Println(calculateMinimumCost(costMatrix))
}

func usage() {
	fmt.Printf("Usage: %s <num of houses> <num of colors>\n", os.Args[0])
	os.Exit(0)
}

func calculateMinimumCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	previousMinimums := []int{0, 0}
	previousColor := -1

	for n, costsForHouse := range costs {
		currentMinimums := []int{math.MaxInt32, math.MaxInt32}
		currentColor := -1

		for k := range costsForHouse {
			if previousColor == k {
				costs[n][k] += previousMinimums[1]
			} else {
				costs[n][k] += previousMinimums[0]
			}

			if costs[n][k] < currentMinimums[0] {
				currentMinimums = []int{costs[n][k], currentMinimums[0]}
				currentColor = k
			} else if costs[n][k] < currentMinimums[1] {
				currentMinimums[1] = costs[n][k]
			}
		}

		copy(previousMinimums, currentMinimums)
		previousColor = currentColor
	}

	return previousMinimums[0]
}

func generateRandomCostMatrix(N, K int) [][]int {
	var costMatrix [][]int

	rand.Seed(time.Now().UnixNano())

	for n := 0; n < N; n++ {
		costMatrix[n] = make([]int, K)
		for k := 0; k < K; k++ {
			costMatrix[n][k] = rand.Intn(100)
		}
	}

	return costMatrix
}
