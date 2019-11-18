package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("getting %s: %s", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Printf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Errorf("parsing %s: as HTML: %v", url, err)
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	paragraphsOrDivs := ElementsByTagName(doc, "p", "div")

	fmt.Println("images:")
	for _, val := range images {
		fmt.Println(val)
	}

	fmt.Println("headings:")
	for _, val := range headings {
		fmt.Println(val)
	}

	fmt.Println("paragraphs or divs:")
	for _, val := range paragraphsOrDivs {
		fmt.Println(val)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var list []*html.Node

	hasElement := func(name string, list []string) bool {
		for _, item := range list {
			if name == item {
				return true
			}
		}
		return false
	}

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode && hasElement(n.Data, name) {
			list = append(list, n)
		}
	}

	forEachNode(doc, startElement, nil)

	return list
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
