package main

import "fmt"

const R = 'R'
const B = 'B'
const G = 'G'

func main() {

}

func segregateRGB(input []byte) ([]byte, int) {
	changes := 0
	rIndex, bIndex := 0, len(input) - 1

	for i := 0; i < len(input); i++ {
		c := input[i]
		
		if c == R && i > rIndex {
			input[i], input[rIndex] = input[rIndex], input[i]
			changes++
			rIndex++

			if input[i] != G {
				i--
			}			
		}

		if c == B && i < bIndex {
			input[i], input[bIndex] = input[bIndex], input[i]
			changes++
			bIndex--

			if input[i] != G {
				i--
			}			
		}

		fmt.Println(i, input)
	}

	return input, changes
}