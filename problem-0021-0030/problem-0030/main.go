package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	var elevationMap []int
	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			usage()
		}

		elevationMap = append(elevationMap, n)
	}

	fmt.Println(calcuateWaterStorage(elevationMap))
}

func usage() {
	fmt.Printf("Usage: %s [elevation map]\n", os.Args[0])
	os.Exit(1)
}

func calcuateWaterStorage(elevationMap []int) int {
	total := 0

	if len(elevationMap) < 2 {
		return 0
	}

	interleavingUnits := 0
	startingIndex := 0
	startingElevation := elevationMap[startingIndex]

	for i := startingIndex + 1; i < len(elevationMap); i++ {
		if elevationMap[i] >= startingElevation {
			total += startingElevation*(i-startingIndex-1) - interleavingUnits
			startingIndex = i
			startingElevation = elevationMap[startingIndex]
			interleavingUnits = 0
		} else {
			interleavingUnits += elevationMap[i]
		}
	}

	interleavingUnits = 0
	startingIndex = len(elevationMap) - 1
	startingElevation = elevationMap[startingIndex]

	for i := startingIndex - 1; i >= 0; i-- {
		if elevationMap[i] > startingElevation {
			total += startingElevation*(startingIndex-i-1) - interleavingUnits
			startingIndex = i
			startingElevation = elevationMap[startingIndex]
			interleavingUnits = 0
		}
	}

	return total
}
