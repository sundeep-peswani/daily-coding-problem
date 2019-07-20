package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand5())
}

func rand5() int {
	n := rand7() * 5 // n = 5-35

	if n < 30 {
		return n / 5
	}

	return rand5()
}

func rand7() int {
	return 1 + rand.Intn(7)
}
