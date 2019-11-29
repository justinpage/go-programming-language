package main

import (
	"fmt"
	"io"
	"os"
)

type Counter struct {
	Writer io.Writer
	Count  int64
}

func (c *Counter) Write(p []byte) (int, error) {
	n, _ := c.Writer.Write(p)
	c.Count += int64(n)
	return len(p), nil
}

func main() {
	w, c := CountingWriter(os.Stdout)

	fmt.Fprintf(w, "spicy jalepeno pastrami ham short loin")
	fmt.Println()

	fmt.Println(w, *c)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &Counter{os.Stdout, 0}
	return cw, &cw.Count
}
