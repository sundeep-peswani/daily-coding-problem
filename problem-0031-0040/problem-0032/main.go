package main

import (
	"math"
)

func main() {

}

func findArbitrage(rates [][]float64) bool {
	n := len(rates)
	graph := make([][]float64, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			graph[i][j] = -1 * math.Log(rates[i][j])
		}
	}

	minDistance := make([]float64, n)
	for i := 0; i < n; i++ {
		minDistance[i] = math.Inf(1)
	}

	minDistance[0] = 0

	for i := 0; i < n-1; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				minDistance[k] = math.Min(minDistance[k], minDistance[j]+graph[j][k])
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if minDistance[j] > minDistance[i]+graph[i][j] {
				return true
			}
		}
	}

	return false
}
