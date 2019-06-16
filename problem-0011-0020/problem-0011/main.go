package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// A Node in a Trie that represents the dictionary
type Node struct {
	ch       rune
	children map[rune]*Node
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	words, err := loadFromDictionary(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to load dictionary from %s: %v\n", os.Args[1], err)
	}

	if len(os.Args[2]) == 0 {
		usage()
	}

	dictionary := buildDictionary(words)
	fmt.Println(dictionary.find(os.Args[2]))
}

func usage() {
	fmt.Printf("Usage: %s <path to dictionary> <prefix>", os.Args[0])
	os.Exit(1)
}

func loadFromDictionary(path string) (words []string, err error) {
	var data []byte

	if data, err = ioutil.ReadFile(path); err != nil {
		return
	}

	return strings.Split(string(data), "\n"), nil
}

func buildDictionary(words []string) *Node {
	dictionary := newNode(0)

	for _, word := range words {
		dictionary.add(word)
	}

	return dictionary
}

func newNode(ch rune) *Node {
	var node Node
	node.ch = ch
	node.children = make(map[rune]*Node)

	return &node
}

func (node *Node) add(word string) {
	if len(word) == 0 {
		return
	}

	n := node
	for _, ch := range word {
		if _, ok := n.children[ch]; !ok {
			n.children[ch] = newNode(ch)
		}

		n = n.children[ch]
	}

	n.children[-1] = newNode(-1)
}

func (node *Node) find(prefix string) (words []string) {
	if len(prefix) == 0 {
		return
	}

	n := node
	for _, ch := range prefix {
		if next, ok := n.children[ch]; !ok {
			return
		} else {
			n = next
		}
	}

	// remove the last character since n contains it
	p := prefix[0 : len(prefix)-1]

	// build words
	for _, suffix := range n.getSuffixes() {
		words = append(words, p+suffix)
	}

	return words
}

func (node *Node) getSuffixes() (suffixes []string) {
	if node.ch == -1 {
		return []string{""}
	}

	for _, child := range node.children {
		for _, suffix := range child.getSuffixes() {
			suffixes = append(suffixes, fmt.Sprintf("%c%s", node.ch, suffix))
		}
	}

	return
}
