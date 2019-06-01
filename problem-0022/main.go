package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	given := os.Args[1]
	dictionary := os.Args[2:]

	fmt.Println(reconstruct(dictionary, given))
}

func usage() {
	fmt.Printf("Usage: %s <given word> <dictionary...>\n", os.Args[0])
	os.Exit(1)
}

type match struct {
	words     []string
	remaining string
}

func reconstruct(dictionary []string, given string) []string {
	d := buildDictionary(dictionary)

	queue := []match{{[]string{}, given}}

	for len(queue) > 0 {
		currentNode := d
		current := queue[0]

		for i, ch := range current.remaining {
			if _, ok := currentNode.children[endOfWord]; ok {
				queue = append(queue, match{
					append(current.words, current.remaining[:i]),
					current.remaining[i:],
				})
			}

			if n, ok := currentNode.children[rune(ch)]; ok {
				currentNode = n
			} else {
				currentNode = d
				break
			}
		}

		if _, ok := currentNode.children[endOfWord]; ok && currentNode != d {
			return append(current.words, current.remaining)
		}

		queue = queue[1:]
	}

	return []string{}
}

type node struct {
	value    rune
	children map[rune]*node
}

const endOfWord = 'a' - 1

func buildDictionary(words []string) *node {
	d := newNode(0)

	for _, word := range words {
		d.add(word)
	}

	return d
}

func newNode(value byte) *node {
	return &node{rune(value), make(map[rune]*node)}
}

func (d *node) add(word string) {
	if len(word) == 0 {
		d.children[endOfWord] = newNode(endOfWord)
		return
	}

	ch := rune(word[0])
	if _, ok := d.children[ch]; !ok {
		d.children[ch] = newNode(word[0])
	}

	d.children[ch].add(word[1:])
}
