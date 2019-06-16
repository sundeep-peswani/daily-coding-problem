package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	n := 10000

	var err error

	if len(os.Args) >= 2 {
		if n, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatalf("Usage: %s <no. of approximations>", os.Args[0])
		}
	}

	fmt.Printf("Monte Carlo approximation of pi with %d approximations is: %.5f\n", n, monteCarloPi(n))
}

func monteCarloPi(n int) float64 {
	inside := 0

	for i := 0; i < n; i++ {
		x, y := rand.Float64(), rand.Float64()

		if x*x+y*y <= 1 {
			inside++
		}
	}

	return 4 * float64(inside) / float64(n)
}
