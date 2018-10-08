package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345"))
}

// intsToString is alike fmt.Sprintf(values) but adds commas.
func comma(values string) string {
	var buf bytes.Buffer
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%c", v)
	}
	return buf.String()
}
