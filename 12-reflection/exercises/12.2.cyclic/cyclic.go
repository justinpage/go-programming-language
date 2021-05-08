package main

import (
	"github.com/justinpage/go-programming-language/12-reflection/exercises/12.2.cyclic/display"
)

func main() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}

	display.Display("c", c)
}
