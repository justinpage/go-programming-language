package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/justinpage/go-programming-language/2-program-structure/exercises/2.2.unitconv/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cl: %v\n", err)
			os.Exit(1)
		}

		fh := unitconv.Fahrenheit(t)
		cl := unitconv.Celsius(t)

		ft := unitconv.Foot(t)
		mt := unitconv.Meter(t)

		lb := unitconv.Pound(t)
		kg := unitconv.Kilogram(t)

		fmt.Printf("%s = %s, %s = %s\n",
			fh, unitconv.FhToCl(fh), cl, unitconv.ClToFh(cl))

		fmt.Printf("%s = %s, %s = %s\n",
			ft, unitconv.FtToMt(ft), mt, unitconv.MtToFt(mt))

		fmt.Printf("%s = %s, %s = %s\n",
			lb, unitconv.LbToKg(lb), kg, unitconv.KgToLb(kg))
	}
}
