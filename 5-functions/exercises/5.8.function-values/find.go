package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var depth int

func main() {
	id := flag.String("id", "", "id attribute")

	flag.Parse()
	for _, url := range flag.Args() {
		outline(url, *id)
	}
}

func outline(url string, id string) {
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

	forEachNode(doc, id, startElement, endElement)
}

// forEachNode calls the functions pre(x) and post(x) for each node x in the
// tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder)
func forEachNode(
	n *html.Node, id string, pre, post func(n *html.Node, id string) bool,
) {
	if !pre(n, id) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, id, pre, post)

			if post(c, id) {
				return
			}
		}
	}
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode && ElementById(n, id) == nil {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++

		return false
	}

	if n.Type == html.ElementNode && ElementById(n, id) != nil {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s='%s'", a.Key, a.Val)
		}
		fmt.Printf(">")
		depth++

		return true
	}

	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode && ElementById(n, id) == nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)

		return false
	}

	if n.Type == html.ElementNode && ElementById(n, id) != nil {
		depth--
		fmt.Printf("</%s>\n", n.Data)

		return true
	}

	return false
}

func ElementById(doc *html.Node, id string) *html.Node {
	if doc.Type == html.ElementNode {
		for _, a := range doc.Attr {
			if a.Key == "id" && a.Val == id {
				return doc
			}
		}
	}
	return nil
}
