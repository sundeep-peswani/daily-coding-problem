package main

import (
	"fmt"
	"sort"
)

func main() {

}

type interval struct {
	start, end int
}

func interleavingIntervals(intervals []interval) int {
	if len(intervals) < 2 {
		return 0
	}

	sort.Slice(intervals, func (i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	end := intervals[0].end
	min := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < end {
			end = lower(end, intervals[i].end)
			min++
		} else {
			end = intervals[i].end
		}
	}

	return min
}

func lower(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (i interval) String() string {
	return fmt.Sprintf("(%d, %d)", i.start, i.end)
}