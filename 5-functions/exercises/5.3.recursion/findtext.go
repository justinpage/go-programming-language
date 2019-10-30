package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findtext", err)
		os.Exit(1)
	}

	for _, text := range visit(nil, doc) {
		fmt.Printf("%#v\n", text)
	}
}

func visit(text []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Parent.Data != "script" &&
		n.Parent.Data != "style" && n.Data != "" {
		lineOfText := strings.TrimSpace(n.Data)

		if lineOfText != "" {
			text = append(text, lineOfText)
		}
	}

	if n.FirstChild != nil {
		text = visit(text, n.FirstChild)
	}

	if n.NextSibling != nil {
		text = visit(text, n.NextSibling)
	}

	return text
}
