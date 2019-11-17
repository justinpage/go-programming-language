package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/justinpage/go-programming-language/5-functions/exercises/5.13.anonymous-functions/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

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
			node = "index.html"
		}

		// Create file under path
		file, err := os.Create(basePath + link.Hostname() + dir + node)
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
	// Crawl the web breadth-first,
	// starting from the command-line arguments
	breadthFirst(crawl, os.Args[1:])
}
