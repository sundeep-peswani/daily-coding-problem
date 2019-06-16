package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	switch os.Args[1] {
	case "encode":
		fmt.Println(encode(os.Args[2]))
	case "decode":
		fmt.Println(decode(os.Args[2]))
	default:
		usage()
	}
}

func usage() {
	fmt.Printf("Usage: %s <encode|decode> <string>\n", os.Args[0])
	os.Exit(1)
}

func encode(str string) string {
	var buffer bytes.Buffer

	if len(str) == 0 {
		return ""
	}

	current := str[0]
	count := 1

	for i := 1; i < len(str); i++ {
		if str[i] == current {
			count++
		} else {
			buffer.WriteString(fmt.Sprintf("%d%c", count, current))

			current = str[i]
			count = 1
		}
	}

	buffer.WriteString(fmt.Sprintf("%d%c", count, current))

	return buffer.String()
}

func decode(str string) string {
	var buffer bytes.Buffer

	if len(str) == 0 {
		return ""
	}

	i := 0
	for i < len(str) {
		j := i

		for str[j] >= '0' && str[j] <= '9' {
			j++
		}

		count, err := strconv.Atoi(str[i:j])
		if err != nil {
			return ""
		}

		char := str[j]

		for k := 0; k < count; k++ {
			buffer.WriteByte(char)
		}

		i = j + 1
	}

	return buffer.String()
}
