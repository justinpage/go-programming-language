package main

import (
	"fmt"

	"github.com/justinpage/go-programming-language/9-concurrency-with-shared-variables/exercises/9.2.popcount/popcount"
)

func main() {
	go count(256)
	go count(42)
	count(42)
	count(256)
}

func count(n uint64) {
	fmt.Println(popcount.PopCount(n))
}
