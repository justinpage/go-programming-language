package main

import (
	"fmt"

	"github.com/justinpage/go-programming-language/2-program-structure/exercises/2.1.tempconv/tempconv"
)

func main() {
	fmt.Printf("%.2f\n", tempconv.CToK(-30))
	fmt.Printf("%.2f\n", tempconv.KToF(160))
}
