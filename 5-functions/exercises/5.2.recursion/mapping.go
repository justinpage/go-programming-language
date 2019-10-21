package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for node, count := range visit(nil, doc) {
		fmt.Printf("%s:%v\n", node, count)
	}
}

func visit(count map[string]int, n *html.Node) map[string]int {
	if count == nil {
		count = make(map[string]int)
	}

	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	if n.FirstChild != nil {
		count = visit(count, n.FirstChild)
	}

	if n.NextSibling != nil {
		count = visit(count, n.NextSibling)
	}

	return count
}
