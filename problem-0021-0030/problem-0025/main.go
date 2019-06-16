package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	fmt.Println(regexMatch(os.Args[1], os.Args[2]))
}

func usage() {
	fmt.Printf("Usage: %s <regex> <string>\n", os.Args[0])
	os.Exit(1)
}

func regexMatch(regex, given string) bool {
	if len(given) == 0 {
		if len(regex) == 0 {
			return true
		}

		for i := len(regex) - 1; i >= 0; i-- {
			if regex[i] == '*' || regex[i] == '.' {
				continue
			}

			return false
		}

		return true
	}

	if len(regex) == 0 {
		return false
	}

	r := len(regex)
	l := len(given)

	switch regex[r-1] {
	case '*':
		prevMatch := regex[r-2] == given[l-1] || regex[r-2] == '.'
		return regexMatch(regex[:r-1], given) || (prevMatch && regexMatch(regex, given[:l-1]))
	case '.':
		return regexMatch(regex[:r-1], given[:l-1])
	case given[l-1]:
		return regexMatch(regex[:r-1], given[:l-1])
	default:
		return false
	}
}
