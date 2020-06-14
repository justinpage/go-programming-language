package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/justinpage/go-programming-language/5-functions/examples/8.find-links/links"
)

type link struct {
	url   string
	depth int
}

type linkList []link

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	depth := flag.Int("depth", 1, "depth limit")
	flag.Parse()

	worklist := make(chan linkList) // lists of URLs, may have duplicates
	unseenLinks := make(chan link)  // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		var l linkList
		for _, url := range flag.Args() {
			l = append(l, link{url, 1})
		}
		worklist <- l
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for k := range unseenLinks {
				foundLinks := crawl(k.url)
				go func(d int) {
					var l linkList
					for _, url := range foundLinks {
						fmt.Println(url)
						l = append(l, link{url, d})
					}
					worklist <- l
				}(k.depth + 1)
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] && link.depth <= *depth {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
