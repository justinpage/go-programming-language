package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma([]float64{+1.2, -2.1, 3.4}))
}

// intsToString is alike fmt.Sprintf(values) but adds commas.
func comma(values []float64) string {
	var buf bytes.Buffer
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%.2f", v)
	}
	return buf.String()
}
