package main

import (
	"flag"
	"fmt"

	"github.com/justinpage/go-programming-language/7-interfaces/exercises/7.6.flags/tempconv"
)

var temp = tempconv.KelvinFlag("temp", 294.261, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
