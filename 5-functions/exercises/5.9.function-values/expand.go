package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "$foo bar $foo buzz"

	replace := func(s string) string {
		return strings.ReplaceAll(input, "$foo", "foo")
	}

	fmt.Println(expand(input, replace))
}

func expand(s string, f func(string) string) string {
	return f(s)
}
