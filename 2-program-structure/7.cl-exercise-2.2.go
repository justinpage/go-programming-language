package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/justinpage/go-programming-language/2-program-structure/lenconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cl: %v\n", err)
			os.Exit(1)
		}
		f := lenconv.Feet(l)
		m := lenconv.Meter(l)
		fmt.Printf("%s = %s, %s = %s\n",
			f, lenconv.FToM(f), m, lenconv.MToF(m))
	}
}
