package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/justinpage/go-programming-language/5-functions/exercises/5.13.anonymous-functions/links"
)

func crawl(endpoint string) []string {
	fmt.Println(endpoint)

	link, _ := url.Parse(endpoint)
	basePath := "./downloads/"

	// Root
	if link.Path == "" {
		// Create parent directory
		err := os.MkdirAll(basePath+link.Hostname(), 0755)
		if err != nil {
			log.Print(err)
		}

		// Create index under parent directory
		file, err := os.Create(basePath + link.Hostname() + "/index.html")
		if err != nil {
			log.Print(err)
		}

		defer file.Close()

		// Download content and save to index
		err = links.Download(endpoint, file)
		if err != nil {
			log.Print(err)
		}
	}

	// Childen
	if link.Path != "" {
		// Create path under parent directory
		dir, node := path.Split(link.Path)

		err := os.MkdirAll(basePath+link.Hostname()+dir, 0755)
		if err != nil {
			log.Print(err)
		}

		// Most directories are pages with no file extension
		if node == "" {
			node = "index"
		}

		// Create file under path
		file, err := os.Create(basePath + link.Hostname() + dir + node + ".html")
		if err != nil {
			log.Print(err)
		}

		defer file.Close()

		// Download content and save to index
		err = links.Download(endpoint, file)
		if err != nil {
			log.Print(err)
		}
	}

	list, err := links.Extract(endpoint)

	list = filter(list, func(v string) bool {
		next, _ := url.Parse(v)
		return link.Hostname() == next.Hostname()
	})

	if err != nil {
		log.Print(err)
	}

	return list
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
