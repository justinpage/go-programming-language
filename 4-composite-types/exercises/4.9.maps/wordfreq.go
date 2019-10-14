package main

import (
	"bufio"
	"fmt"
	"os"
)

type file string

func main() {
	counts := make(map[file]int)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
			continue
		}
		countWords(f, counts)
		f.Close()
	}

	fmt.Printf("%-15s|%-15s\n", "word", "count")
	for w, c := range counts {
		fmt.Printf("%-15v|%-3d\n", w, c)
	}
}

func countWords(f *os.File, counts map[file]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[file(input.Text())]++
	}
}
