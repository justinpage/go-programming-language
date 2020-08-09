package main

import (
	"fmt"
	"time"
)

// Run this service with the time:
// e.g. `$ time go run main.go`
//
// Using this program, we measured a total time of 2m11s
// After that, the program is killed by SIGKILL which cannot be affected by the
// signals package: https://golang.org/pkg/os/signal/#hdr-Types_of_signals
//
// Some other interesting statistics include:
// Real memory size: 9.01 GB
// Virtual memory size: 34.44 GB
func main() {
	start := time.Now()
	var run func(ch chan struct{})
	var count int
	run = func(ch chan struct{}) {
		count++
		for {
			// Number of goroutines we counted 1104686
			// panic: too many concurrent operations on a single file or socket (max 1048575)
			// fmt.Printf("%d %d\n", time.Since(start), count)
			ch := make(chan struct{})
			go run(ch)
		}
	}
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("%d %d\n", time.Since(start), count)
		default:
			ch := make(chan struct{})
			go run(ch)
		}
	}
}
