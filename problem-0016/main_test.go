package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const N = 100

func TestLog(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for n := 100; n < 100000; n *= 10 {
		l := newLog(N)
		data := fillLog(l, n)
		i := rand.Intn(N)

		fmt.Printf("Testing %d, with i = %d and n -i = %d\n", n, i, n-i)

		actual := l.getLast(i)
		if actual != data[n-i] {
			t.Errorf("Testing with %d logs, expected '%s', actual '%s'\n", n, data[n-i], actual)
		}
	}
}

func fillLog(l *log, n int) []string {
	data := make([]string, n)

	for i := 0; i < n; i++ {
		s := randStr10()
		l.record(s)
		data[i] = s
	}

	return data
}

func randStr10() string {
	b := make([]byte, 10)
	var sb strings.Builder

	for {
		if n, err := rand.Read(b); err == nil {
			for i := 0; i < n; i++ {
				sb.WriteByte((b[i]-0)%26 + 'a')
			}

			return sb.String()
		}
	}
}
