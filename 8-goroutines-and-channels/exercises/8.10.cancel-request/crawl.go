package main

import (
	"fmt"
	"log"
	"os"

	"github.com/justinpage/go-programming-language/8-goroutines-and-channels/exercises/8.10.cancel-request/links"
)

func crawl(url string, done <-chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, done)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	done := make(chan struct{})

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Support cancellation
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for {
				select {
				case <-done:
					// Drain unseenLinks to allow existing goroutines to finish.
					for range unseenLinks {
						// Do nothing
					}
					return
				case link := <-unseenLinks:
					foundLinks := crawl(link, done)
					go func() { worklist <- foundLinks }()
				}

			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for {
		select {
		case <-done:
			return
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}
	}
}
