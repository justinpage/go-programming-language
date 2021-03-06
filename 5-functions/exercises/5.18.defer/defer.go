package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		filename, n, err := fetch(url)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("filename:", filename)
		fmt.Println("written:", n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create("/tmp/" + local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)

	// Close file, but prefer error from Copy, if any
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}
