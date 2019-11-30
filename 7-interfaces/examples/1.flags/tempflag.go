package main

import (
	"flag"
	"fmt"

	"github.com/justinpage/go-programming-language/7-interfaces/examples/1.flags/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
