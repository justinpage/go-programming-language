// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	categories := make(map[string]int) // counts of unicode categories
	invalid := 0                       // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error

		// check for any issues
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		switch {
		case unicode.IsLetter(r):
			categories["letter"]++
		case unicode.IsDigit(r):
			categories["digit"]++
		case unicode.IsPunct(r):
			categories["punct"]++
		case unicode.IsSpace(r):
			categories["space"]++
		}

	}

	fmt.Printf("rune\tcount\n")
	for c, n := range categories {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
