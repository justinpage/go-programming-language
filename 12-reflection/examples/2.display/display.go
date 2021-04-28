package main

import (
	"github.com/justinpage/go-programming-language/12-reflection/examples/1.display/display"
	"github.com/justinpage/go-programming-language/12-reflection/examples/1.display/eval"
)

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	display.Display("e", e)
}
