package main

import (
	"fmt"
	"time"

	"github.com/justinpage/go-programming-language/12-reflection/examples/1.format/format"
)

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{x}))
	fmt.Println(format.Any([]time.Duration{d}))
}
