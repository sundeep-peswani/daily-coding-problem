package main

func main() {

}

func calculateTransitiveClosure(graph [][]int) [][]int {
	result := make([][]int, len(graph))
	for i, row := range graph {
		result[i] = make([]int, len(graph))
		for _, vertex := range row {
			result[i][vertex] = 1
		}
	}

	for i, row := range result {
		for j := range row {
			result[i][j] = findPath(result, i, j)
		}
	}

	return result
}

func findPath(closure [][]int, i, j int) int {
	if closure[i][j] == 1 {
		return 1
	}

	for vertex, hasPath := range closure[i] {
		if hasPath == 0 || vertex == i {
			continue
		}

		if findPath(closure, vertex, j) == 1 {
			return 1
		}
	}

	return 0
}
