package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	s := os.Args[1]
	k, err := strconv.Atoi(os.Args[2])
	if err != nil {
		usage()
	}

	lines := splitIntoLines(s, k)
	for _, line := range lines {
		fmt.Println(line)
	}
}

func usage() {
	fmt.Printf("Usage: %s <some string> <num of characters per line>\n", os.Args[0])
	os.Exit(1)
}

func splitIntoLines(s string, k int) []string {
	var lines []string
	var err error

	if k == 0 {
		return nil
	}

	var line, word strings.Builder

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if lines, err = addWord(lines, k, &word, &line); err != nil {
				return nil
			}
		} else {
			word.WriteByte(s[i])
		}
	}

	if lines, err = addWord(lines, k, &word, &line); err != nil {
		return nil
	}

	if line.Len() > 0 {
		if line.Len() > k {
			return nil
		}

		lines = append(lines, line.String())
	}

	return lines
}

func addWord(lines []string, k int, word, line *strings.Builder) ([]string, error) {
	// start of new line
	if line.Len() == 0 {
		if word.Len() > k {
			return nil, fmt.Errorf("Word '%s' is longer than the limit", word.String())
		}

		line.WriteString(word.String())
		word.Reset()
		return lines, nil
	}

	// append word
	if word.Len() <= k-line.Len()-1 {
		line.WriteByte(' ')
		line.WriteString(word.String())
		word.Reset()
		return lines, nil
	}

	// add line to lines and re-try
	lines = append(lines, line.String())
	line.Reset()

	return addWord(lines, k, word, line)
}
