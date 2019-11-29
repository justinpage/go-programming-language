package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type Counter struct {
	Words int
	Lines int
}

func (c *Counter) Write(p []byte) (int, error) {
	words := bufio.NewScanner(bytes.NewBuffer(p))
	lines := bufio.NewScanner(bytes.NewBuffer(p))

	words.Split(bufio.ScanWords)

	for words.Scan() {
		c.Words++
	}

	lines.Split(bufio.ScanLines)

	for lines.Scan() {
		c.Lines++
	}

	return len(p), nil
}

func main() {
	var w Counter

	fmt.Fprintf(&w, "spicy jalepeno pastrami ham short loin")
	fmt.Println(w)

	var l Counter

	fmt.Fprintf(&l, "spicy jalepeno\npastrami ham short\nloin")
	fmt.Println(l)
}
