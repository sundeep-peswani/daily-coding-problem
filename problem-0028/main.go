package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	words := strings.Split(os.Args[1], " ")
	k, err := strconv.Atoi(os.Args[2])
	if err != nil {
		usage()
	}

	for _, line := range justifyText(words, k) {
		fmt.Println(line)
	}
}

func usage() {
	fmt.Printf("Usage: %s <string> <length of line>\n", os.Args[0])
	os.Exit(1)
}

func justifyText(words []string, k int) []string {
	var lines []string

	var line []string
	var wordslen = 0

	for _, word := range words {
		w := len(word)

		if wordslen+w+len(line) >= k {
			// add to lines
			lines = append(lines, justifyLine(line, wordslen, k))

			// reset
			line = []string{word}
			wordslen = w
		} else {
			line = append(line, word)
			wordslen += w
		}
	}
	if wordslen > 0 {
		lines = append(lines, justifyLine(line, wordslen, k))
	}

	return lines
}

func justifyLine(words []string, l, k int) string {
	var buff bytes.Buffer

	w := len(words)
	diff := k - l
	each := diff / w
	add := (diff + each) % w

	buff.WriteString(words[0])
	for i := 1; i < w; i++ {
		for j := 0; j < each; j++ {
			buff.WriteByte(' ')
		}
		if add > 0 {
			buff.WriteByte(' ')
			add--
		}

		buff.WriteString(words[i])
	}

	return buff.String()
}
