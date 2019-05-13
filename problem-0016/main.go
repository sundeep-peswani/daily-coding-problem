package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	rand.Seed(time.Now().UnixNano())

	d, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		usage()
	}

	var l *log

	if n, err := strconv.Atoi(os.Args[2]); err != nil || n <= 0 {
		usage()
	} else {
		l = newLog(n)
	}

	for _, line := range strings.Split(string(d), "\n") {
		l.record(line)
	}

	i := rand.Intn(len(l.records))
	if n, err := strconv.Atoi(os.Args[3]); err != nil || n < 0 || n > len(l.records) {
		usage()
	} else {
		i = n
	}

	fmt.Printf("The %dth last record is: %s\n", i, l.getLast(i))
}

func usage() {
	fmt.Printf("Usage: %s <path to a large file with records> <size of log> <nth record to retrieve>\n", os.Args[0])
	os.Exit(1)
}

type log struct {
	records []string
	pos     int
}

func newLog(n int) *log {
	var l log

	l.records = make([]string, n)

	return &l
}

func (l *log) record(orderID string) {
	l.records[l.pos] = orderID
	l.pos = (l.pos + 1) % len(l.records)
}

func (l *log) getLast(i int) string {
	n := len(l.records)
	x := (l.pos + n - i) % n

	return l.records[x]
}
