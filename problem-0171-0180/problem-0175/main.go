package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	transitions := []transition{
		transition{"a", "a", 0.9},
		transition{"a", "b", 0.075},
		transition{"a", "c", 0.025},
		transition{"b", "a", 0.15},
		transition{"b", "b", 0.8},
		transition{"b", "c", 0.05},
		transition{"c", "a", 0.25},
		transition{"c", "b", 0.25},
		transition{"c", "c", 0.5},
	}

	result := markovChainSimulation(transitions, 5000, "a")
	fmt.Println(result)
}

type transition struct {
	src, dest string
	probability float64
}

func markovChainSimulation(transitions []transition, steps int, start string) map[string]int {
	current := start
	graph := make(map[string]map[string]float64)
	result := make(map[string]int)

	rand.Seed(time.Now().UnixNano())

	for _, transition := range(transitions) {
		src, dest := transition.src, transition.dest

		if _, ok := graph[src]; !ok {
			graph[src] = make(map[string]float64)
		}

		graph[src][dest] = transition.probability
	}

	for i := 0; i < steps; i++ {
		result[current]++
		current = markovNextState(graph, current)
	}

	return result
}

func markovNextState(graph map[string]map[string]float64, current string) string {
	r := rand.Float64()
	c := 0.0

	for k, v := range graph[current] {
		c += v
		if c > r {
			return k
		}
	}

	return current
}
