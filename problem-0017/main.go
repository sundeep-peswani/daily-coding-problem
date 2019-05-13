package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	fmt.Println(findLongestPathToFile(os.Args[1]))
}

func usage() {
	fmt.Printf("Usage: %s <string representing the filesystem>", os.Args[0])
	os.Exit(1)
}

func findLongestPathToFile(fs string) string {
	longest := ""
	subdirs := 0
	isFile := false
	isRoot := true

	var path strings.Builder

	for _, ch := range fs {
		switch ch {
		case '\n':
			if isFile {
				longest = max(path.String(), longest)
			}
			subdirs = 0
			isFile = false
			isRoot = true
		case '\t':
			isRoot = false
			subdirs++
		case '.':
			isFile = true
			path.WriteRune(ch)
		default:
			if subdirs > 0 {
				paths := strings.SplitN(path.String(), "/", subdirs+1)
				path.Reset()
				path.WriteString(strings.Join(paths[0:subdirs], "/"))
				path.WriteRune('/')
				subdirs = 0
			} else if subdirs == 0 && isRoot {
				path.Reset()
			}
			path.WriteRune(ch)
			isRoot = false
		}
	}

	if isFile {
		return max(path.String(), longest)
	}

	return longest
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}
