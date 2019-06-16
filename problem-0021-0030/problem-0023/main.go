package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		usage()
	}

	start := makeCoordinates(os.Args[1])
	end := makeCoordinates(os.Args[2])
	maze := newMaze(os.Args[3])

	fmt.Println(maze.solve(start, end))
}

func usage() {
	fmt.Printf("Usage: %s <start> <end> <comma-separated maze>\n", os.Args[0])
	os.Exit(1)
}

func makeCoordinates(str string) []int {
	coords := []int{0, 0}

	vals := strings.Split(str, ",")
	if len(vals) != 2 {
		return coords
	}

	for i, v := range vals {
		if c, err := strconv.Atoi(v); err != nil {
			coords[i] = c
		}
	}

	return coords
}

type maze struct {
	board [][]bool
}

func newMaze(str string) *maze {
	var m maze

	for _, row := range strings.Split(str, ",") {
		var line []bool

		for _, ch := range row {
			line = append(line, ch == '#')
		}

		m.board = append(m.board, line)

	}

	return &m
}

type step struct {
	count  int
	coords []int
}

func (m *maze) solve(start, end []int) (int, error) {
	q := []step{{0, start}}

	for len(q) > 0 {
		next := q[0]
		q = q[1:]

		if m.isMarked(next.coords) {
			continue
		}

		fmt.Printf("Checking %#v\n", next)

		if next.coords[0] == end[0] && next.coords[1] == end[1] {
			return next.count, nil
		}

		m.mark(next.coords)

		for _, c := range m.next(next.coords) {
			q = append(q, step{next.count + 1, c})
		}
	}

	return 0, fmt.Errorf("Unable to solve")
}

func (m *maze) mark(coords []int) {
	m.board[coords[0]][coords[1]] = true
}

func (m *maze) isMarked(coords []int) bool {
	return m.board[coords[0]][coords[1]]
}

func (m *maze) next(coords []int) [][]int {
	var possibles [][]int

	a, b := coords[0], coords[1]

	// vertical
	if a > 0 && !m.isMarked([]int{a - 1, b}) {
		possibles = append(possibles, []int{a - 1, b})
	}
	if a < len(m.board)-1 && !m.isMarked([]int{a + 1, b}) {
		possibles = append(possibles, []int{a + 1, b})
	}

	// horizontal
	if b > 0 && !m.isMarked([]int{a, b - 1}) {
		possibles = append(possibles, []int{a, b - 1})
	}
	if b < len(m.board[a])-1 && !m.isMarked([]int{a, b + 1}) {
		possibles = append(possibles, []int{a, b + 1})
	}

	return possibles
}
