package main

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func TestRLE(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		str := generateString(rand.Intn(100))

		encoded := encode(str)
		decoded := decode(encoded)

		if str == encoded {
			t.Errorf("Test %d: failed to encode '%s': '%s'\n", i+1, str, encoded)
		}

		if decoded != str {
			t.Errorf("Test %d: failed to decode '%s' into '%s': '%s'\n", i+1, encoded, str, decoded)
		}
	}
}

func generateString(len int) string {
	var buffer bytes.Buffer

	for i := 0; i < len; {
		count := max(1, rand.Intn(len-i))
		char := byte('A' + rand.Intn(26))

		for j := 0; j < count; j++ {
			buffer.WriteByte(char)
		}

		i += count
	}

	return buffer.String()
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
