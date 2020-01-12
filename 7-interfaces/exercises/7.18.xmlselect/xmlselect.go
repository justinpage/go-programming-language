// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	render(visit(dec))
}

func visit(dec *xml.Decoder) *Element {
	var el Element
	var travel func(dec *xml.Decoder, el *Element)

	travel = func(dec *xml.Decoder, el *Element) {
		for {
			tok, err := dec.Token()
			if err == io.EOF {
				return
			}
			switch tok := tok.(type) {
			case xml.StartElement:
				el.Type = tok.Name
				for _, v := range tok.Attr {
					el.Attr = append(el.Attr, v)
				}
				var ch Element
				travel(dec, &ch)
				el.Children = append(el.Children, ch)
			case xml.CharData:
				el.Children = append(el.Children, CharData(tok))
			}
		}
	}

	travel(dec, &el)

	return &el
}

func render(el *Element) {
	var depth int
	var outline func(el *Element)

	outline = func(el *Element) {
		for _, v := range el.Children {
			switch v := v.(type) {
			case Element:
				depth++
				fmt.Printf("%*s%s->\n", depth*2, "", el.Type.Local)
				depth++
				outline(&v)
			case CharData:
				if strings.TrimSpace(string(v)) != "" {
					fmt.Printf("%*s%s\n", depth*2, "", v)
				}
				depth--
			}
		}
		depth--
	}

	outline(el)
}
