package main

import "fmt"

func main() {

}

func pickXYUsingMath(x, y, b uint32) uint32 {
	return (x * b) + (y * (1 - b))
}

func pickXYUsingBitOperations(x, y, b uint32) uint32 {
	b32 := b
	for i := 1; i <= 16; i *= 2 {
		b32 = ((b32 << uint(i)) | b32)
	}

	fmt.Printf("b = %d, b32 = %b\n", b, b32)

	return (x & b32) | (y &^ b32)
}
