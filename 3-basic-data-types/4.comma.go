package main

import (
	"fmt"
)

func main() {
	fmt.Println(comma("12345594949"))
}

// comma inserts comma in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
