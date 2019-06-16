package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	fmt.Println(isMatching(os.Args[1]))
}

func usage() {
	fmt.Printf("Usage: %s <string to check>\n", os.Args[0])
	os.Exit(1)
}

func isMatching(given string) bool {
	s := stack{}

	for _, ch := range given {
		b := byte(ch)
		o, c := openingIndex(b), closingIndex(b)

		if o != -1 {
			s.push(b)
		} else if c != -1 {
			if s.isEmpty() {
				return false
			}

			o = openingIndex(s.peek())
			if o != c {
				return false
			}
			s.pop()
		}
	}

	return s.isEmpty()
}

func find(slice []byte, ch byte) int {
	for i, b := range slice {
		if b == ch {
			return i
		}
	}

	return -1
}

func openingIndex(ch byte) int {
	return find([]byte{'(', '[', '{'}, ch)
}

func closingIndex(ch byte) int {
	return find([]byte{')', ']', '}'}, ch)
}

type stack struct {
	data []byte
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() byte {
	return s.data[len(s.data)-1]
}

func (s *stack) push(ch byte) {
	s.data = append(s.data, ch)
}

func (s *stack) pop() byte {
	ch := s.peek()
	s.data = s.data[:len(s.data)-1]
	return ch
}
