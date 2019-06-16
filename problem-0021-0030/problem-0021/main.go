package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) < 3 || len(os.Args)%2 == 0 {
		usage()
	}

	fmt.Println(countRooms(createSchedule(toInt(os.Args[1:]))))
}

func usage() {
	fmt.Printf("Usage: %s <pairs of start and end times>\n", os.Args[0])
	os.Exit(1)
}

type interval struct {
	start, end int
}

func toInt(strs []string) []int {
	nums := make([]int, len(strs))

	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return []int{}
		}

		nums = append(nums, num)
	}

	return nums
}

func createSchedule(times []int) []interval {
	if len(times)%2 != 0 || len(times) == 0 {
		return []interval{}
	}

	var schedule []interval
	for i := 0; i < len(times); i += 2 {
		start, end := times[i], times[i+1]
		if start > end {
			return []interval{}
		}

		schedule = append(schedule, interval{start, end})
	}

	return schedule
}

func countRooms(times []interval) int {
	n := len(times)
	if n <= 1 {
		return n
	}

	var starts, ends []int
	for _, time := range times {
		starts = append(starts, time.start)
		ends = append(ends, time.end)
	}

	sort.Ints(starts)
	sort.Ints(ends)

	i, j, inUse, total := 1, 0, 1, 1
	for i < n && j < n {
		if starts[i] < ends[j] {
			inUse++
			i++
			if inUse > total {
				total = inUse
			}
		} else {
			inUse--
			j++
		}
	}

	return total
}
