package main

import (
	"bufio"
	"fmt"
	"os"
)

type file struct {
	name, line string
}

func main() {
	counts := make(map[file]int)
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
	for file, count := range counts {
		if count > 1 {
			fmt.Printf(
				"File: %s\tLine: %s\tCount: %d\n",
				file.name, file.line, count,
			)
		}
	}

}

func countLines(f *os.File, counts map[file]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[file{f.Name(), input.Text()}]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
