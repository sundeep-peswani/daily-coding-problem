package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		usage()
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(string(pickUniformFromStream(f)))
}

func usage() {
	fmt.Printf("Usage: %s <path to a large file>\n", os.Args[0])
	os.Exit(1)
}

func pickUniformFromStream(r io.Reader) []byte {
	b := make([]byte, 10)
	p := make([]byte, 10)

	n := 0
	for {
		_, err := r.Read(b)
		if err != nil {
			return p
		}

		n++
		if rand.Intn(n) == 0 {
			copy(p, b)
		}
	}
}
