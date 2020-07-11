package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s [url...url]\n", os.Args[0])

	}

	fetch(os.Args[1:])
}

func fetch(urls []string) {
	resp := make(chan string)
	done := make(chan struct{})

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp <- request(u, done)
		}(url)
	}

	go func(r chan<- string) {
		wg.Wait()
		close(r)
	}(resp)

	for {
		select {
		case <-done:
			// Drain resp to allow existing goroutines to finish
			for r := range resp {
				log.Println(r)
			}
			return
		case r := <-resp:
			log.Println(r)
			close(done)
		}
	}
}

func request(url string, done <-chan struct{}) string {
	start := time.Now()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("error creating request %s: %v", url, err)
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("%s: %v", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Sprintf("%s: %v", url, err)
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Sprintf("%s: %v", url, err)
	}

	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
