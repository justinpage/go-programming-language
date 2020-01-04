// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type element struct {
	Name  string
	Value string
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []element // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			var e element
			e.Name = tok.Name.Local
			for _, v := range tok.Attr {
				if v.Name.Local == "class" {
					e.Value = "." + v.Value
				}
				if v.Name.Local == "id" {
					e.Value = "#" + v.Value
				}
			}
			stack = append(stack, e)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				var buf bytes.Buffer
				for i, e := range stack {
					fmt.Fprintf(&buf, "%s", e.Name) // tag name
					if e.Value != "" {
						fmt.Fprintf(&buf, "%s", e.Value) // attribute
					}

					if i < len(stack)-1 {
						buf.WriteByte(' ')
					}
				}
				fmt.Fprintf(&buf, ": %s\n", tok)
				fmt.Print(buf.String())
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []element, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if compare(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// compare reports whether x is identical to y
func compare(x element, y string) bool {
	switch y {
	case x.Name:
		return true
	case x.Value:
		return true
	}
	return false
}
