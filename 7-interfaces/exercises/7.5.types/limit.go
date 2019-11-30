package main

import (
	"io"
	"os"
)

type LimitedReader struct {
	R io.Reader // underlying reader we will use
	N int64     // max number of bytes to read
}

type Reader struct {
	s string
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s[0:len(p)])
	return n, io.EOF
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}

	n, err = l.R.Read(p)

	l.N -= int64(n)

	return
}

func main() {
	r := NewReader("Hello, World")
	lr := LimitReader(r, 5)
	io.Copy(os.Stdout, lr)
}

func NewReader(s string) *Reader {
	return &Reader{s}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
