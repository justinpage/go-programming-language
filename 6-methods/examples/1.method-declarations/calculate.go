package main

import (
	"fmt"

	"github.com/justinpage/go-programming-language/6-methods/examples/1.method-declarations/geometry"
)

func main() {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
