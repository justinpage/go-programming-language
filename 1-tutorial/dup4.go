package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line struct {
    File, Word string
}

func main() {
	counts := make(map[Line]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line.Word, line.File)
		}
	}
}

func countLines(f *os.File, counts map[Line]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[Line{f.Name(), input.Text()}]++
	}
}
