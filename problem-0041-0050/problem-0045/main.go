package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand7())
}

func rand5() int {
	return 1 + rand.Intn(5)
}

func rand7() int {
	n := 5*rand5() + rand5() // (5 - 25) + (1 - 5) = (6 - 30)
	if n < 27 {              // 6-26 = 21 numbers
		return n%7 + 1
	}
	return rand7()
}
