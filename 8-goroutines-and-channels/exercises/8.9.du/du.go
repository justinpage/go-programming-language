// Need to revisit. Not a complete working solution
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)

type root struct {
	nfiles, nbytes int64
}

func main() {
	// Determine the initial directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	for _, v := range roots {
		var n sync.WaitGroup

		fileSizes := make(chan int64)
		n.Add(1)

		// Traverse
		go walkDir(v, &n, fileSizes)

		// Wait until finish to close
		go func() {
			n.Wait()
			close(fileSizes)
		}()

		// Process files in tree
		var nfiles, nbytes int64
		go func(v string) {
		loop:
			for {
				select {
				case size, ok := <-fileSizes:
					if !ok {
						break loop // fileSizes was closed
					}
					nfiles++
					nbytes += size
				case <-tick:
					printDiskUsage(v, nfiles, nbytes)
				}
			}
			printDiskUsage(v, nfiles, nbytes) // final totals
		}(v)
	}

	select {}
}

func printDiskUsage(p string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", p, nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "du1: %v\n", err)
	}
	return entries
}
