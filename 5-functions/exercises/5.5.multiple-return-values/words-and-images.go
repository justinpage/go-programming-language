package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countWordsAndImages: %v\n", err)
			continue
		}

		fmt.Printf("%s ~ words=%v images=%v\n", url, words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return countWords(0, n), countImages(0, n)
}

func countImages(count int, n *html.Node) int {
	if n.Type == html.ElementNode && n.Data == "img" {
		fmt.Println(n)
		count++
	}

	if n.FirstChild != nil {
		count = countImages(count, n.FirstChild)
	}

	if n.NextSibling != nil {
		count = countImages(count, n.NextSibling)
	}

	return count
}

func countWords(count int, n *html.Node) int {
	if n.Type == html.TextNode && n.Parent.Data != "script" &&
		n.Parent.Data != "style" && n.Data != "" {
		lineOfText := strings.TrimSpace(n.Data)

		if lineOfText != "" {
			b := bytes.NewReader([]byte(lineOfText))
			input := bufio.NewScanner(b)
			input.Split(bufio.ScanWords)

			for input.Scan() {
				count++
			}
		}
	}

	if n.FirstChild != nil {
		count = countWords(count, n.FirstChild)
	}

	if n.NextSibling != nil {
		count = countWords(count, n.NextSibling)
	}

	return count
}
