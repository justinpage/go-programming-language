package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Reader struct {
	s string
}

func (r *Reader) Read(b []byte) (n int, err error) {
	copy(b, []byte(r.s))
	return len(r.s), io.EOF
}

func main() {
	body := NewReader("<html><body>hello</body></html>")
	doc, err := html.Parse(body)
	if err != nil {
		fmt.Errorf("parsing %s: as HTML: %v", err)
	}
	fmt.Println(doc.FirstChild.LastChild.FirstChild.Data)
}

func NewReader(s string) *Reader {
	return &Reader{s}
}
