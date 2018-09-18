package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

func main() {
	var output []byte

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		if has := strings.HasPrefix(url, "http://"); has != true {
			url = "http://" + url
		}

		go fetch(url, ch)
	}
	for range os.Args[1:] {
		result := []byte(<-ch)
		output = append(output, result...)
	}

	elapsed := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	output = append(output, []byte(elapsed)...)

	ioutil.WriteFile("/tmp/fetchall-results", output, 0644)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("While reading %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
