package main

import (
	"fmt"

	"github.com/justinpage/go-programming-language/2-program-structure/tempconv"
)

func main() {
	fmt.Printf("%.2f\n", tempconv.CToK(-30))
	fmt.Printf("%.2f", tempconv.KToF(160))
}
