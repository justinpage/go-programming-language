package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123"))       // 123
	fmt.Println(comma("1234"))      // 1,234
	fmt.Println(comma("12345"))     // 12,345
	fmt.Println(comma("123456"))    // 123,456
	fmt.Println(comma("1234567"))   // 1,234,567
	fmt.Println(comma("12345678"))  // 12,345,678
	fmt.Println(comma("123456789")) // 123,456,789
}

// comma inserts comma in a buffer every three places
func comma(s string) string {
	var buf bytes.Buffer

	l := len(s)
	r := l % 3

	if l <= 3 {
		return s
	}

	if r >= 1 {
		buf.WriteString(s[:r])
		buf.WriteString(",")
	}

	for i := r; i < len(s); i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < l {
			buf.WriteString(",")
		}
	}

	return buf.String()
}
