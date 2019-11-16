package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int           // bad pattern: global
var buffer bytes.Buffer // bad pattern: global

func main() {
	for _, url := range os.Args[1:] {
		doc := Outline(url)
		fmt.Println(doc.String())
	}
}

func Outline(url string) bytes.Buffer {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(&buffer, "getting %s: %s", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(&buffer, "getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(&buffer, "parsing %s: as HTML: %v", url, err)
	}

	forEachNode(doc, startElement, endElement)

	return buffer
}

// forEachNode calls the functions pre(x) and post(x) for each node x in the
// tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder)
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

func startElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		fmt.Fprintf(&buffer, "%*s<!--%s-->\n", depth*2, "", n.Data)
		depth++
	case html.ElementNode:
		fmt.Fprintf(&buffer, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(&buffer, " %s='%s'", a.Key, a.Val)
		}
		fmt.Fprintf(&buffer, ">\n")
		depth++
	case html.TextNode:
		if n.Data != "" {
			fmt.Fprintf(&buffer, "%*s%s\n", depth*2, "", n.Data)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		depth--
	case html.ElementNode:
		depth--
		if n.FirstChild != nil && n.LastChild != nil {
			fmt.Fprintf(&buffer, "%*s</%s>\n", depth*2, "", n.Data)
		}
	case html.TextNode:
		depth--
	}
}
